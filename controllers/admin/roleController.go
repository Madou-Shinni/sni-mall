package admin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	TitleIsEmpty      = "标题不能为空！"
	FailedAddRole     = "添加角色失败！"
	ParameterError    = "参数错误！"
	FailedUpdateRole  = "修改角色失败！"
	RoleStatusDefault = 1
)

type RoleController struct {
	BaseController
}

// List 角色列表
func (con RoleController) List(c *gin.Context) {
	var roleList []models.Role
	mysql.DB.Find(&roleList)
	con.SuccessAndData(c, roleList)
}

// Add 添加角色
func (con RoleController) Add(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ") // 获取标题去除空格
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, TitleIsEmpty)
	}
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = RoleStatusDefault
	role.AddTime = int(time.Now().Unix())

	err := mysql.DB.Create(&role).Error
	if err != nil {
		con.Error(c, FailedAddRole)
		return
	}
	con.Success(c)
}

// Update 修改角色
func (con RoleController) Update(c *gin.Context) {
	id := c.PostForm("id")
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	idInt, err := utils.StringToInt(id)
	if err != nil {
		con.Error(c, ParameterError)
	}
	role := models.Role{Id: idInt}
	if affected := mysql.DB.Find(&role).RowsAffected; affected < 1 { // 数据库里面查询不到
		return
	}
	role.Title = title
	role.Description = description
	sqlErr := mysql.DB.Save(&role).Error
	if sqlErr != nil {
		con.Error(c, FailedUpdateRole)
		return
	}
	con.Success(c)
}

// Delete 删除角色
func (con RoleController) Delete(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	role := models.Role{Id: id}
	mysql.DB.Delete(&role)
	con.Success(c)
}
