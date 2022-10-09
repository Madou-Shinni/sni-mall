package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	"xiaomi-mall/models/utils"
	pbManager "xiaomi-mall/proto/rbacManager"
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

	roleClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	rsp, _ := roleClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{})

	con.SuccessAndData(c, rsp.ManagerList)
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
	roleClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	rsp, _ := roleClient.ManagerAdd(context.Background(), &pbManager.ManagerAddRequest{
		Username: username,
		Password: utils.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   int64(roleId),
		Status:   1,
		AddTime:  int64(utils.GetUnix()),
	})
	if !rsp.Success {
		con.Error(c, rsp.Message)
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

	//注意：判断密码是否为空 为空表示不修改密码 不为空表示修改密码
	if password != "" {
		password = utils.Md5(password)
	}

	roleClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	rsp, _ := roleClient.ManagerUpdate(context.Background(), &pbManager.ManagerUpdateRequest{
		Id:       id,
		Username: username,
		Password: password,
		Email:    email,
		Mobile:   mobile,
		RoleId:   int64(roleId),
		Status:   1,
		AddTime:  int64(utils.GetUnix()),
	})
	if !rsp.Success {
		con.Error(c, rsp.Message)
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

	roleClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	rsp, _ := roleClient.ManagerDelete(context.Background(), &pbManager.ManagerDeleteRequest{
		Id: id,
	})
	if !rsp.Success {
		con.Error(c, rsp.Message)
		return
	}

	con.Success(c)
}
