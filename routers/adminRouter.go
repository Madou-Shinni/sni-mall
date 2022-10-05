package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/controllers/admin"
	"xiaomi-mall/middlewares"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		// 获取管理员登录验证码
		adminRouter.GET("/captcha", admin.LoginController{}.Captcha)
		// 管理员登录
		adminRouter.POST("/login", admin.LoginController{}.Login)
		// 管理员退出登录
		adminRouter.POST("/logout", admin.LoginController{}.Logout)

		adminRouter.Use(middlewares.InitAdminAuthMiddleware)

		// 管理员列表
		adminRouter.GET("/manager/list", admin.ManagerController{}.List)
		// 添加管理员
		adminRouter.POST("/manager", admin.ManagerController{}.Add)
		// 修改管理员
		adminRouter.PUT("/manager", admin.ManagerController{}.Update)
		// 删除管理员
		adminRouter.DELETE("/manager", admin.ManagerController{}.Delete)

		// 角色列表
		adminRouter.GET("/role/list", admin.RoleController{}.List)
		// 添加角色
		adminRouter.POST("/role/add", admin.RoleController{}.Add)
		// 修改角色
		adminRouter.PUT("/role/update", admin.RoleController{}.Update)
		// 删除角色
		adminRouter.DELETE("/role/delete", admin.RoleController{}.Delete)

		// 权限列表
		adminRouter.GET("/access/list", admin.AccessController{}.List)
		// 添加权限
		adminRouter.POST("/access", admin.AccessController{}.Add)
		// 修改权限
		adminRouter.PUT("/access", admin.AccessController{}.Update)
		// 删除权限
		adminRouter.DELETE("/access", admin.AccessController{}.Delete)
	}
}
