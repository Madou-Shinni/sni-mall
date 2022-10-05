package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"xiaomi-mall/models/utils"
)

const (
	//USERINFO          = "userInfo" // 用户信息
	NOPERMISSIONSCODE = 302
	NOPERMISSIONSMSG  = "没有权限！"
	ContextUserIdKey  = "userId"
)

// InitAdminAuthMiddleware 权限判断中间件
func InitAdminAuthMiddleware(c *gin.Context) {
	// 1.获取访问的url
	//url := strings.Split(c.Request.URL.String(), "?")[0]
	// 2.获取session里面保存的信息
	//session := sessions.Default(c)
	//userInfo := session.Get(USERINFO)
	//userInfoStr, ok := userInfo.(string) // 类型断言
	//if ok {
	//	// userInfoStr存在
	//	var userInfoStruct models.Manager
	//	err := json.Unmarshal([]byte(userInfoStr), &userInfoStruct)
	//	if err != nil || userInfoStruct == (models.Manager{}) { // 反序列化失败 或者 结构体为空（登陆失败）
	//		if url != "/admin/login" && url != "/admin/captcha" {
	//			c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
	//		}
	//	}
	//} else {
	//	// 用户没有登录
	//	if url != "/admin/login" && url != "/admin/captcha" {
	//		c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
	//	}
	//}

	// ---------------------------- 以下使用token代替session ----------------------------
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
	c.Next()
}
