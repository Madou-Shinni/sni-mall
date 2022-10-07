package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "xiaomi-mall/controllers"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	MaxAddress = 10 // 最大收货地址上限
)

type AddressController struct {
	BaseController
}

// List 收货地址列表
func (con AddressController) List(c *gin.Context) {
	uid := GetCurrentUserId(c) // 从请求上下文中获取用户id
	var address []models.Address
	mysql.DB.Where("uid = ?", uid).Find(&address)
	c.JSON(http.StatusOK, address)
}

// Add 添加收货地址
func (con AddressController) Add(c *gin.Context) {
	uid := GetCurrentUserId(c) // 从请求上下文中获取用户id
	address := models.Address{}
	if err2 := c.ShouldBind(&address); err2 != nil {
		c.JSON(http.StatusOK, "参数错误，请你稍后重试！")
		return
	}

	var addressList []models.Address
	if affected := mysql.DB.Where("uid = ?", uid).Find(&addressList).RowsAffected; affected > MaxAddress {
		c.JSON(http.StatusOK, "收货地址达到上限，请你稍后重试！")
		return
	}

	address.Uid = int(uid)
	if sqlErr := mysql.DB.Create(address).Error; sqlErr != nil {
		c.JSON(http.StatusOK, "添加收货地址失败，请你稍后重试！")
		return
	}
	c.String(http.StatusOK, "success")
}

// Update 修改收货地址
func (con AddressController) Update(c *gin.Context) {
	address := models.Address{}
	if err2 := c.ShouldBind(&address); err2 != nil {
		c.JSON(http.StatusOK, "参数错误，请你稍后重试！")
		return
	}

	if err := mysql.DB.Save(&address).Error; err != nil {
		c.JSON(http.StatusOK, "修改失败，请你稍后重试！")
		return
	}

	c.String(http.StatusOK, "success")
}

// Delete 删除收货地址
func (con AddressController) Delete(c *gin.Context) {
	id, _ := utils.StringToInt(c.Query("id"))
	address := models.Address{Id: id}
	if err := mysql.DB.Delete(&address).Error; err != nil {
		c.JSON(http.StatusOK, "删除失败，请你稍后重试！")
		return
	}
	c.String(http.StatusOK, "success")
}
