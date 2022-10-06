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

		// 首页菜单
		adminRouter.GET("/", admin.IndexController{}.Index)

		// 管理员列表
		adminRouter.GET("/manager", admin.ManagerController{}.List)
		// 添加管理员
		adminRouter.POST("/manager", admin.ManagerController{}.Add)
		// 修改管理员
		adminRouter.PUT("/manager", admin.ManagerController{}.Update)
		// 删除管理员
		adminRouter.DELETE("/manager", admin.ManagerController{}.Delete)

		// 角色列表
		adminRouter.GET("/role", admin.RoleController{}.List)
		// 添加角色
		adminRouter.POST("/role", admin.RoleController{}.Add)
		// 修改角色
		adminRouter.PUT("/role", admin.RoleController{}.Update)
		// 删除角色
		adminRouter.DELETE("/role", admin.RoleController{}.Delete)
		// 角色授权
		adminRouter.POST("/role/auth", admin.RoleController{}.Auth)

		// 权限列表
		adminRouter.GET("/access", admin.AccessController{}.List)
		// 添加权限
		adminRouter.POST("/access", admin.AccessController{}.Add)
		// 修改权限
		adminRouter.PUT("/access", admin.AccessController{}.Update)
		// 删除权限
		adminRouter.DELETE("/access", admin.AccessController{}.Delete)

		// 轮播图列表
		adminRouter.GET("/focus", admin.FocusController{}.List)
		// 添加轮播图
		adminRouter.POST("/focus", admin.FocusController{}.Add)
		// 修改轮播图
		adminRouter.PUT("/focus", admin.FocusController{}.Update)
		// 删除轮播图
		adminRouter.DELETE("/focus", admin.FocusController{}.Delete)
	}
}
