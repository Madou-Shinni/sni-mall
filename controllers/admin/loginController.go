package admin

import (
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

// Login 登录
func (con LoginController) Login(c *gin.Context) {
	con.Success(c)
}
