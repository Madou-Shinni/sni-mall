package admin

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	FailedAddAccess             = "添加权限失败！"
	FailedUpdateAccess          = "修改权限失败！"
	FailedDeleteChildrenAtFirst = "删除失败，请先删除子模块！"
)

type AccessController struct {
	BaseController
}

// List 权限列表
func (con AccessController) List(c *gin.Context) {
	var accessList []models.Access
	mysql.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
	con.SuccessAndData(c, accessList)
}

// Add 添加权限
func (con AccessController) Add(c *gin.Context) {
	moduleName := c.PostForm("module_name")
	accessType, err := utils.StringToInt(c.PostForm("access_type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err2 := utils.StringToInt(c.PostForm("module_id"))
	sort, err3 := utils.StringToInt(c.PostForm("sort"))
	status, err4 := utils.StringToInt(c.PostForm("status"))
	description := c.PostForm("description")
	addTime := int(utils.GetUnix())
	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, ParameterError)
		return
	}
	access := models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
		AddTime:     addTime,
	}
	sqlErr := mysql.DB.Create(&access).Error
	if sqlErr != nil {
		con.Error(c, FailedAddAccess)
		return
	}
	con.Success(c)
}

// Update 修改权限
func (con AccessController) Update(c *gin.Context) {
	id, err := utils.StringToInt(c.PostForm("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	moduleName := c.PostForm("module_name")
	accessType, err := utils.StringToInt(c.PostForm("access_type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err2 := utils.StringToInt(c.PostForm("module_id"))
	sort, err3 := utils.StringToInt(c.PostForm("sort"))
	status, err4 := utils.StringToInt(c.PostForm("status"))
	description := c.PostForm("description")
	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, ParameterError)
		return
	}
	access := models.Access{
		Id:          id,
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	sqlErr := mysql.DB.Save(&access).Error
	if sqlErr != nil {
		con.Error(c, FailedUpdateAccess)
		return
	}
	con.Success(c)
}

// Delete 删除权限
func (con AccessController) Delete(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	access := models.Access{Id: id}
	mysql.DB.Find(&access)
	if access.ModuleId == 0 { // 顶级模块
		var accessList []models.Access
		mysql.DB.Where("module_id = ?", access.Id).Find(&accessList)
		if len(accessList) == 0 {
			con.Error(c, FailedDeleteChildrenAtFirst)
			return
		}
	}
	mysql.DB.Delete(&access)
	con.Success(c)
}
