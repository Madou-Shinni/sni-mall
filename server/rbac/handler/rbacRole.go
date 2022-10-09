package handler

import (
	"context"
	"rbac/models"
	mysql "rbac/models/mysql"
	pb "rbac/proto/rbacRole"
	"strconv"
)

const (
	FailedAdd    = "添加失败！"
	FailedUpdate = "修改失败！"
	FailedDelete = "删除失败！"
)

type RbacRole struct{}

// RoleGet 获取角色(有id查询id角色，没有获取全部角色)
func (e *RbacRole) RoleGet(ctx context.Context, req *pb.RoleGetRequest, rsp *pb.RoleGetResponse) error {
	where := "1=1"
	var roleList []models.Role
	if req.Id > 0 { // 有id查询id角色，没有获取全部角色
		where += "AND id =" + strconv.Itoa(int(req.Id))
	}
	mysql.DB.Where(where).Find(&roleList)

	// 处理数据 不能直接rsp.RoleList = roleList
	var tempList []*pb.RoleModel
	for _, v := range roleList {
		tempList = append(tempList, &pb.RoleModel{
			Id:          int64(v.Id),
			Title:       v.Title,
			Description: v.Description,
			Status:      int64(v.Status),
			AddTime:     int64(v.AddTime),
		})
	}
	rsp.RoleList = tempList

	return nil
}

// RoleAdd 增加角色
func (e *RbacRole) RoleAdd(ctx context.Context, req *pb.RoleAddRequest, rsp *pb.RoleAddResponse) error {
	role := models.Role{}
	role.Title = req.RoleModel.Title
	role.Description = req.RoleModel.Description
	role.Status = int(req.RoleModel.Status)
	role.AddTime = int(req.RoleModel.AddTime)

	err := mysql.DB.Create(&role).Error
	if err != nil {
		rsp.Success = false
		rsp.Msg = FailedAdd
		return err
	}

	rsp.Success = true
	return nil
}

// RoleUpdate 修改角色
func (e *RbacRole) RoleUpdate(ctx context.Context, req *pb.RoleUpdateRequest, rsp *pb.RoleUpdateResponse) error {
	role := models.Role{Id: int(req.RoleModel.Id)}
	role.Title = req.RoleModel.Title
	role.Description = req.RoleModel.Description
	sqlErr := mysql.DB.Save(&role).Error
	if sqlErr != nil {
		rsp.Success = false
		rsp.Msg = FailedUpdate
		return sqlErr
	}

	rsp.Success = true
	return nil
}

// RoleDelete 删除角色
func (e *RbacRole) RoleDelete(ctx context.Context, req *pb.RoleDeleteRequest, rsp *pb.RoleDeleteResponse) error {
	role := models.Role{Id: int(req.Id)}
	err := mysql.DB.Delete(&role).Error
	if err != nil {
		rsp.Success = false
		rsp.Msg = FailedDelete
		return err
	}

	rsp.Success = false
	return nil
}
