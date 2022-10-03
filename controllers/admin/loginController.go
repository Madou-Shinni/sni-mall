package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	utils "xiaomi-mall/models/utils"
	"xiaomi-mall/models/utils/jwt"
)

const (
	FailedCaptVerify = "验证码错误！"
	FailedLogin      = "用户名或密码错误！"
	FailedSystem     = "系统故障请联系管理员！"
)

type LoginController struct {
	BaseController
}

// Login 登录
func (con LoginController) Login(c *gin.Context) {
	// 1.获取请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	// 2.验证验证码
	if flag := models.CaptVerify(captchaId, verifyValue); flag == false {
		con.Error(c, FailedCaptVerify)
	} else { // 验证成功
		// 3.查询数据库，判断用户名密码是否存在
		userInfo := models.Manager{}
		password = utils.Md5(password)
		result := mysql.DB.Where("username = ? AND password = ?", username, password).First(&userInfo).RowsAffected // 返回找到的记录数
		if result > 0 {
			// 4.登录成功，保存用户信息
			//session := sessions.Default(c)
			// 注意：session无法直接保存结构体！ 把结构体转换成json字符串
			//userInfoJsonStr, _ := json.Marshal(userInfo)
			//session.Set("userInfo", userInfoJsonStr)
			//session.Save()
			token, err := jwt.GenToken(int64(userInfo.Id), username)
			if err != nil {
				con.Error(c, FailedSystem)
				return
			}
			con.SuccessAndData(c, token)
		} else {
			con.Error(c, FailedLogin)
		}
	}
}

// Logout 退出登录
func (con LoginController) Logout(c *gin.Context) {
	//session := sessions.Default(c)
	//session.Delete("userInfo")
	//session.Save()
	con.Success(c)
}

// Captcha 获取验证码
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.CaptMake()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
