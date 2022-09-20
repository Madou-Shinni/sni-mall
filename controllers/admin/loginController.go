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
	con.Success(c)
}

// Captcha 验证码
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
