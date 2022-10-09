package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
	pbRole "xiaomi-mall/proto/rbacRole"
)

const (
	TitleIsEmpty      = "标题不能为空！"
	FailedAddRole     = "添加角色失败！"
	ParameterError    = "参数错误！"
	FailedUpdateRole  = "修改角色失败！"
	RoleStatusDefault = 1
	TopModuleId       = 0 // 顶层模块
)

type RoleController struct {
	BaseController
}

// List 角色列表
func (con RoleController) List(c *gin.Context) {
	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	rsp, _ := roleClient.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	con.SuccessAndData(c, rsp.RoleList)
}

// Add 添加角色
func (con RoleController) Add(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ") // 获取标题去除空格
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, TitleIsEmpty)
	}

	roleModel := &pbRole.RoleModel{
		Title:       title,
		Description: description,
		Status:      RoleStatusDefault,
		AddTime:     utils.GetUnix(),
	}
	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	rsp, err := roleClient.RoleAdd(context.Background(), &pbRole.RoleAddRequest{RoleModel: roleModel})
	if err != nil {
		con.Error(c, rsp.Msg)
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

	roleModel := &pbRole.RoleModel{
		Id:          int64(idInt),
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     utils.GetUnix(),
	}
	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	rsp, err := roleClient.RoleUpdate(context.Background(), &pbRole.RoleUpdateRequest{RoleModel: roleModel})
	if err != nil {
		con.Error(c, rsp.Msg)
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

	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	rsp, err := roleClient.RoleDelete(context.Background(), &pbRole.RoleDeleteRequest{Id: int64(id)})
	if err != nil {
		con.Error(c, rsp.Msg)
		return
	}

	con.Success(c)
}

// Auth 角色授权
func (con RoleController) Auth(c *gin.Context) {
	// 获取角色id
	roleId, err := utils.StringToInt(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	// 获取权限id（多个）
	accessIds := c.PostFormArray("access_node[]")

	// 调用微服务
	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	rsp, _ := roleClient.RoleAuth(context.Background(), &pbRole.RoleAuthRequest{
		RoleId:    int64(roleId),
		AccessIds: accessIds,
	})
	if !rsp.Success {
		con.Error(c, "授权失败！")
		return
	}

	con.Success(c)
}

// GetAuth 获取权限列表并且获取角色权限（当前角色的权限在权限列表中显示被选中状态）
func (con RoleController) GetAuth(c *gin.Context) {
	// 1.获取角色id
	roleId, err := utils.StringToInt(c.Query("role_id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}

	// 调用微服务
	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	rsp, _ := roleClient.RoleGetAuth(context.Background(), &pbRole.RoleGetAuthRequest{RoleId: int64(roleId)})

	con.SuccessAndData(c, rsp.AccessList)
}
