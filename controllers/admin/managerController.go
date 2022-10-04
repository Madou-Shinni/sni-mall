package admin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	LengthErrorUsernameAndPassword = "用户名或密码长度不合法！"
	LengthErrorPassword            = "密码长度不能小于6"
	UsernameExists                 = "用户名已存在！"
	FailedAddManager               = "添加管理员失败！"
	FailedUpdateManager            = "修改管理员失败!"
)

type ManagerController struct {
	BaseController
}

// List 管理员列表
func (con ManagerController) List(c *gin.Context) {
	var managerList []models.Manager
	mysql.DB.Preload("Role").Find(&managerList)
	con.SuccessAndData(c, managerList)
}

// Add 添加管理员
func (con ManagerController) Add(c *gin.Context) {
	roleId, err := utils.StringToInt(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, LengthErrorUsernameAndPassword)
		return
	}
	// 判断管理员是否存在
	manager := models.Manager{Username: username}
	if affected := mysql.DB.Where("username = ?", username).Find(&manager).RowsAffected; affected > 0 {
		con.Error(c, UsernameExists)
		return
	}
	manager = models.Manager{
		Username: username,
		Password: utils.Md5(password),
		Mobile:   mobile,
		Email:    email,
		RoleId:   roleId,
		AddTime:  int(utils.GetUnix()),
	}
	if sqlErr := mysql.DB.Create(&manager).Error; sqlErr != nil {
		con.Error(c, FailedAddManager)
		return
	}
	con.Success(c)
}

// Update 修改管理员
func (con ManagerController) Update(c *gin.Context) {
	id, err := utils.StringToInt64(c.PostForm("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	roleId, err2 := utils.StringToInt(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, ParameterError)
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	manager := models.Manager{Id: id}
	mysql.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId
	// 如果密码为空则不修改，不为空则修改
	if password != "" {
		if len(password) > 6 {
			con.Error(c, LengthErrorPassword)
			return
		}
		manager.Password = utils.Md5(password)
	}
	sqlErr := mysql.DB.Save(&manager).Error
	if sqlErr != nil {
		con.Error(c, FailedUpdateManager)
		return
	}
	con.Success(c)
}

// Delete 删除管理员
func (con ManagerController) Delete(c *gin.Context) {
	id, err := utils.StringToInt64(c.PostForm("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	manager := models.Manager{Id: id}
	mysql.DB.Delete(&manager)
	con.Success(c)
}
