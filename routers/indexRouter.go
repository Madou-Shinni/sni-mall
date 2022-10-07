package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/controllers/index"
)

// IndexRouterInit 前台路由
func IndexRouterInit(r *gin.Engine) {
	indexRouter := r.Group("/")
	{
		// 获取首页数据
		indexRouter.GET("/", index.DefaultController{}.Index)
		// 获取分类数据（动态路由）
		indexRouter.PUT("/category/:id", index.ProductController{}.Category)
	}
}
