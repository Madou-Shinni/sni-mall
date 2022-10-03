package middlewares

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
)

const (
	USERINFO          = "userInfo" // 用户信息
	NOPERMISSIONSCODE = 302
	NOPERMISSIONSMSG  = "没有权限！"
)

// InitAdminAuthMiddleware 权限判断中间件
func InitAdminAuthMiddleware(c *gin.Context) {
	// 1.获取访问的url
	url := strings.Split(c.Request.URL.String(), "?")[0]
	// 2.获取session里面保存的信息
	session := sessions.Default(c)
	userInfo := session.Get(USERINFO)
	userInfoStr, ok := userInfo.(string) // 类型断言
	if ok {
		// userInfoStr存在
		var userInfoStruct models.Manager
		err := json.Unmarshal([]byte(userInfoStr), &userInfoStruct)
		if err != nil || userInfoStruct == (models.Manager{}) { // 反序列化失败 或者 结构体为空（登陆失败）
			if url != "/admin/login" && url != "/admin/captcha" {
				c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
			}
		}
	} else {
		// 用户没有登录
		if url != "/admin/login" && url != "/admin/captcha" {
			c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
		}
	}
}
