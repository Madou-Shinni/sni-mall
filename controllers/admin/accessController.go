package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"xiaomi-mall/models"
	"xiaomi-mall/models/utils"
	pbAccess "xiaomi-mall/proto/rbacAccess"
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
	roleClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	rsp, _ := roleClient.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})
	con.SuccessAndData(c, rsp.AccessList)
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

	roleClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	rsp, _ := roleClient.AccessAdd(context.Background(), &pbAccess.AccessAddRequest{
		ModuleName:  moduleName,
		Type:        int64(accessType),
		ActionName:  actionName,
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
		AddTime:     int64(addTime),
	})

	if !rsp.Success {
		con.Error(c, rsp.Message)
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

	roleClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	rsp, _ := roleClient.AccessUpdate(context.Background(), &pbAccess.AccessUpdateRequest{
		Id:          int64(id),
		ModuleName:  moduleName,
		Type:        int64(accessType),
		ActionName:  actionName,
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
	})

	if !rsp.Success {
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

	roleClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	rsp, _ := roleClient.AccessDelete(context.Background(), &pbAccess.AccessDeleteRequest{
		Id: int64(id),
	})

	if !rsp.Success {
		con.Error(c, rsp.Message)
		return
	}

	con.Success(c)
}
