package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"xiaomi-mall/models/utils"
)

func InitUserAuthMiddleware(c *gin.Context) {
	// 判断用户是否登录
	// 从请求头中获取token
	tokenStr := c.Request.Header.Get("Authorization")
	// 用户不存在
	if tokenStr == "" {
		c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
		c.Abort() //阻止执行
		return
	}
	// token格式错误
	tokenSlice := strings.SplitN(tokenStr, " ", 2)
	if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
		c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
		c.Abort() //阻止执行
		return
	}
	// 解析token
	tokenStruck, err := utils.ParseToken(tokenSlice[1])
	if err != nil {
		c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
		c.Abort() //阻止执行
		return
	}
	// token超时
	if time.Now().Unix() > tokenStruck.ExpiresAt {
		c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
		c.Abort() //阻止执行
		return
	}
	// 将请求的userId信息保存到上下文，后续可以通过c.Get("userId")获取
	c.Set(ContextUserIdKey, tokenStruck.UserId)
}
