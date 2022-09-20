package main

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/routers"
)

func main() {
	// 创建默认路由引擎
	r := gin.Default()

	// 注册路由
	routers.AdminRouterInit(r)

	// 运行
	r.Run(":8080")
}
