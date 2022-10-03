package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaomi-mall/models"
)

type LoginController struct {
	BaseController
}

// Login 登录
func (con LoginController) Login(c *gin.Context) {
	// 1.获取请求参数
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	// 2.验证验证码
	if flag := models.CaptVerify(captchaId, verifyValue); flag == false {
		c.String(http.StatusOK, "验证失败！")
	} else {
		c.String(http.StatusOK, "验证成功!")
	}
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
