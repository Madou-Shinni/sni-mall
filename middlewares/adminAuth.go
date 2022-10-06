package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"strings"
	"time"
	request "xiaomi-mall/controllers"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	//USERINFO          = "userInfo" // 用户信息
	NOPERMISSIONSCODE = 302
	NOPERMISSIONSMSG  = "没有权限！"
	ContextUserIdKey  = "userId" // 当前用户id
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

	// 用户登陆成功的权限处理
	// 获取路由，当前："/admin/xx"，表中："/xx"，需要做处理去掉"/admin"之后再匹配
	path := strings.Split(c.Request.URL.String(), "?")[0]
	url := strings.Replace(path, "/admin", "", 1) // 去除一个 -1全部去除
	// 判断是否是需要做权限处理的url
	if !excludeAuthPath(url) {
		// 1.获取用户信息（角色id）
		userId := request.GetCurrentUserId(c)
		userInfo := models.Manager{}
		mysql.DB.Select("role_id").Where("id = ?", userId).Find(&userInfo)
		// 2.获取当前用户的权限id列表
		var roleAccessList []models.RoleAccess
		var access models.Access
		// 把权限id放在一个map类型的对象里面
		roleAccessMap := make(map[int]int)
		mysql.DB.Select("access_id").Where("role_id = ?", userInfo.RoleId).Find(&roleAccessList)
		for _, v := range roleAccessList {
			roleAccessMap[v.AccessId] = v.AccessId
		}
		// 查询url对应的权限id
		mysql.DB.Select("id").Where("url = ?", url).Find(&access)
		// 3.匹配当前用户是否有访问当前路由的权限
		if _, ok := roleAccessMap[access.Id]; !ok {
			c.String(NOPERMISSIONSCODE, NOPERMISSIONSMSG)
			c.Abort()
			return
		}
	}

	c.Next()
}

// excludeAuthPath 对指定的路由url排除权限判断
func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",") // 转化成切片

	for _, v := range excludeAuthPathSlice {
		if v != urlPath { // 访问的url不等于排除的url
			return false
		}
	}

	return true
}
