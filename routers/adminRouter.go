package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/controllers/admin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		// 获取管理员登录验证码
		adminRouter.GET("/captcha", admin.LoginController{}.Captcha)
		// 管理员登录
		adminRouter.POST("/login", admin.LoginController{}.Login)
		// 管理员列表
		adminRouter.GET("/list", admin.AdminController{}.List)
		// 添加管理员
		adminRouter.POST("/add", admin.AdminController{}.Add)
		// 修改管理员
		adminRouter.PUT("/update", admin.AdminController{}.Update)
		// 删除管理员
		adminRouter.DELETE("/delete", admin.AdminController{}.Delete)
	}
}
