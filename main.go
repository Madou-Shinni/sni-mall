package main

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/routers"
)

func main() {
	// 创建默认路由引擎
	r := gin.Default()

	// 基于cookie的存储引擎，security111 参数是密钥
	//store := cookie.NewStore([]byte("security111"))
	// 配置session中间件,store 是前面配置的存储引擎,我们可以替换成其他存储引擎
	//r.Use(sessions.Sessions("mysession", store))

	// 注册路由
	routers.AdminRouterInit(r)

	// 运行
	r.Run(":8080")
}
