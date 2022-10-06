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
		// 通用修改status接口
		adminRouter.GET("/changeStatus", admin.CommonController{}.ChangeStatus)

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

		// 商品顶级分类列表
		adminRouter.GET("/goodsCate-topList", admin.GoodsCateController{}.TopList)
		// 商品分类列表
		adminRouter.GET("/goodsCate", admin.GoodsCateController{}.List)
		// 添加商品分类
		adminRouter.POST("/goodsCate/add", admin.GoodsCateController{}.Add)
		// 修改商品分类
		adminRouter.PUT("/goodsCate/update", admin.GoodsCateController{}.Update)
		// 删除商品分类
		adminRouter.DELETE("/goodsCate/delete", admin.GoodsCateController{}.Delete)

		// 商品类型列表
		adminRouter.GET("/goodsType", admin.GoodsTypeController{}.List)
		// 添加商品类型
		adminRouter.POST("/goodsType/add", admin.GoodsTypeController{}.Add)
		// 修改商品类型
		adminRouter.PUT("/goodsType/update", admin.GoodsTypeController{}.Update)
		// 删除商品类型
		adminRouter.DELETE("/goodsType/delete", admin.GoodsTypeController{}.Delete)

		// 商品类型属性列表
		adminRouter.GET("/goodsTypeAttribute", admin.GoodsTypeAttributeController{}.List)
		adminRouter.GET("/goodsTypeAttribute-ListById", admin.GoodsTypeAttributeController{}.ListById)
		// 添加商品类型属性
		adminRouter.POST("/goodsTypeAttribute/add", admin.GoodsTypeAttributeController{}.Add)
		// 修改商品类型属性
		adminRouter.PUT("/goodsTypeAttribute/update", admin.GoodsTypeAttributeController{}.Update)
		// 删除商品类型属性
		adminRouter.DELETE("/goodsTypeAttribute/delete", admin.GoodsTypeAttributeController{}.Delete)

		// 商品类型属性列表
		adminRouter.GET("/goods", admin.GoodsController{}.List)
		// 商品信息列表（商品分类、颜色、类型列表）
		adminRouter.GET("/goodsInfoList", admin.GoodsController{}.GoodsInfoList)
		// 添加商品类型属性
		adminRouter.POST("/goods/add", admin.GoodsController{}.Add)
		// 修改商品类型属性
		adminRouter.PUT("/goods/update", admin.GoodsController{}.Update)
		// 删除商品类型属性
		adminRouter.DELETE("/goods/delete", admin.GoodsController{}.Delete)
	}
}
