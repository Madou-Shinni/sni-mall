package admin

import (
	"github.com/gin-gonic/gin"
	request "xiaomi-mall/controllers"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
)

type IndexController struct {
	BaseController
}

// Index 根据用户权限动态显示菜单
func (con IndexController) Index(c *gin.Context) {
	// 1.获取用户信息（角色id）
	userId := request.GetCurrentUserId(c)
	userInfo := models.Manager{}
	mysql.DB.Select("role_id").Where("id = ?", userId).Find(&userInfo)
	// 2.获取当前角色的权限
	var roleAccessList []models.RoleAccess
	var accessIds []int
	var accessList []models.Access
	mysql.DB.Select("access_id").Where("role_id = ?", userInfo.RoleId).Find(&roleAccessList)
	for _, v := range roleAccessList {
		accessIds = append(accessIds, v.AccessId)
	}
	mysql.DB.Where("id in (?)", accessIds).Preload("AccessItem").Find(&accessList)
	con.SuccessAndData(c, accessList)
}
