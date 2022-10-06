package admin

import (
	"fmt"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	BaseController
}

// Get 获取系统设置
func (con SettingController) Get(c *gin.Context) {
	setting := models.Setting{}
	mysql.DB.First(&setting)
	con.SuccessAndData(c, setting)
}

// Update 修改系统设置
func (con SettingController) Update(c *gin.Context) {
	setting := models.Setting{Id: 1}
	mysql.DB.Find(&setting)
	if err := c.ShouldBind(&setting); err != nil {
		fmt.Println(err)
		con.Error(c, FailedUpdate)
		return
	}

	// 上传图片 logo
	siteLogo, err1 := utils.UploadImg(c, "site_logo")
	if len(siteLogo) > 0 && err1 == nil {
		setting.SiteLogo = siteLogo
	}
	//上传图片 no_picture
	noPicture, err2 := utils.UploadImg(c, "no_picture")
	if len(noPicture) > 0 && err2 == nil {
		setting.NoPicture = noPicture
	}

	sqlErr := mysql.DB.Save(&setting).Error
	if sqlErr != nil {
		con.Error(c, FailedUpdate)
		return
	}

	con.Success(c)

}
