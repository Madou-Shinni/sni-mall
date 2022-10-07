package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/redis"
	"xiaomi-mall/models/utils"
)

const (
	MaxSendCount = 10 // 短信发送上限
)

type UserController struct {
	BaseController
}

// Captcha 获取验证码
func (con UserController) Captcha(c *gin.Context) {
	id, b64s, err := models.CaptMake()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

// SendCode 发送短信验证码
func (con UserController) SendCode(c *gin.Context) {
	phone := c.Query("phone")
	verifyCode := c.Query("verifyCode")
	captchaId := c.Query("captchaId")
	fmt.Println(captchaId, verifyCode)
	// 1、验证图形验证码是否正确
	if flag := models.CaptVerify(captchaId, verifyCode); !flag {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "验证码输入错误，请重试",
		})
		return
	}

	/*
		2、判断手机格式是否合法
				pattern := `^[\d]{11}$`
				reg := regexp.MustCompile(pattern)
				reg.MatchString(phone)
	*/
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "手机号格式不合法",
		})
		return
	}

	//3、验证手机号是否注册过
	var userList []models.User
	mysql.DB.Where("phone = ?", phone).Find(&userList)
	if len(userList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "手机号已经注册，请直接登录",
		})
		return
	}
	//4、判断当前ip地址今天发送短信的次数

	ip := c.ClientIP()
	currentDay := utils.GetDay() //20211211
	var sendCount int64
	mysql.DB.Table("user_temp").Where("ip=? AND add_day=?", ip, currentDay).Count(&sendCount)
	if sendCount > MaxSendCount {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "此ip今天发送短信的次数已经达到上限，请明天再试",
		})
		return
	}
	//5、验证当前手机号今天发送的次数是否合法
	var userTemp []models.UserTemp
	smsCode := utils.RandCode()
	sign := utils.Md5(phone + currentDay) //签名：主要用于页面跳转传值
	mysql.DB.Where("phone = ? AND add_day=?", phone, currentDay).Find(&userTemp)
	if len(userTemp) > 0 { // 今天已经给这个号码发送过验证码
		if userTemp[0].SendCount > MaxSendCount {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "此手机号今天发送短信的次数已经达到上限，请明天再试",
			})
			return
		} else {
			//1、生成短信验证码  发送验证码  调用前面课程的接口
			//fmt.Println("----------自己集成发送短信的接口--------")
			utils.SMS(userTemp[0].Phone, smsCode)
			//2、服务器保持验证码
			redis.CacheDB.Set(userTemp[0].Phone, smsCode, 5*60)

			//3、更新发送短信的次数
			oneUserTemp := models.UserTemp{}
			mysql.DB.Where("id=?", userTemp[0].Id).Find(&oneUserTemp)
			oneUserTemp.SendCount = oneUserTemp.SendCount + 1
			mysql.DB.Save(&oneUserTemp)

			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "发送短信成功",
				"sign":    sign,
			})
			return
		}

	} else { // 今天没有 给这个号码发送过验证码
		//1、生成短信验证码  发送验证码  调用前面课程的接口
		//fmt.Println("----------自己集成发送短信的接口--------")
		utils.SMS(phone, smsCode)
		//2、服务器保持验证码
		redis.CacheDB.Set(phone, smsCode, 5*60)

		//3、记录发送短信的次数

		oneUserTemp := models.UserTemp{
			Ip:        ip,
			Phone:     phone,
			SendCount: 1,
			AddDay:    currentDay,
			AddTime:   int(utils.GetUnix()),
			Sign:      sign,
		}
		mysql.DB.Create(&oneUserTemp)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "发送短信成功",
			"sign":    sign,
		})
		return

	}
}

// Register 注册
func (con UserController) Register(c *gin.Context) {
	phone := c.PostForm("phone")
	smsCode := c.PostForm("smsCode") // 短信验证码
	// 验证短信验证码
	if b := utils.VerifyCode(phone, smsCode); !b {
		c.String(http.StatusOK, "验证码错误或者失效！")
		return
	}
	password := c.PostForm("password")
	confirmPwd := c.PostForm("confirm_password") // 确认密码
	if password != confirmPwd {
		c.String(http.StatusOK, "密码和确认密码不相同，请重新输入！")
		return
	}

	// 完成注册
	user := models.User{
		Phone:    phone,
		Password: utils.Md5(password),
		AddTime:  int(utils.GetUnix()),
		LastIp:   c.ClientIP(),
		Email:    "",
		Status:   1,
	}
	if sqlErr := mysql.DB.Create(&user).Error; sqlErr != nil {
		c.String(http.StatusOK, "注册失败，请稍后再试！")
		return
	}
	// 删除redis中的短信验证码
	redis.CacheDB.Del(phone)
	// 执行登录
	token, err := utils.GenToken(int64(user.Id), phone)
	if err != nil {
		c.String(http.StatusOK, "系统繁忙，登陆失败！")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Login 登录
func (con UserController) Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	verifyCode := c.PostForm("verifyCode")
	// 验证验证码
	if flag := models.CaptVerify(captchaId, verifyCode); !flag {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "验证码输入错误，请重试",
		})
		return
	}
	// 查询数据库
	user := models.User{}
	mysql.DB.Where("phone = ?", phone).Find(&user)

	// 比较加密密码
	md5Pwd := utils.Md5(password)
	if user.Password != md5Pwd {
		c.String(http.StatusOK, "密码错误！")
		return
	}

	// 登录
	token, err := utils.GenToken(int64(user.Id), phone)
	if err != nil {
		c.String(http.StatusOK, "登陆失败，请稍后再试！")
		return
	}
	// 返回用户信息前先清空密码
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"userInfo": user,
	})
}
