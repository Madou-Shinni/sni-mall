package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (b BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "success!")
}

func (b BaseController) SuccessAndData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": data,
	})
}

func (b BaseController) Error(c *gin.Context, msg string) {
	c.String(http.StatusOK, msg)
}
