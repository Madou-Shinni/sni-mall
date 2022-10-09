package handler

import (
	"context"
	"rbac/models"
	mysql "rbac/models/mysql"
	pb "rbac/proto/rbacManager"
	"strconv"
)

type RbacManager struct{}

// ManagerGet 获取管理员(有id查询id管理员，没有获取全部管理员)
func (e *RbacManager) ManagerGet(ctx context.Context, req *pb.ManagerGetRequest, rsp *pb.ManagerGetResponse) error {
	where := "1=1"
	var ManagerList []models.Manager
	if req.Id > 0 { // 有id查询id管理员，没有获取全部管理员
		where += "AND id =" + strconv.Itoa(int(req.Id))
	}
	if len(req.Username) > 0 { // 有id查询username管理员，没有获取全部管理员
		where += "AND id =" + req.Username
	}
	mysql.DB.Where(where).Preload("Role").Find(&ManagerList)

	// 处理数据 不能直接rsp.ManagerList = ManagerList
	var tempList []*pb.ManagerModel
	for _, v := range ManagerList {
		tempList = append(tempList, &pb.ManagerModel{
			Id:       int64(v.Id),
			Username: v.Username,
			Mobile:   v.Mobile,
			Email:    v.Email,
			Status:   int64(v.Status),
			RoleId:   int64(v.RoleId),
			AddTime:  int64(v.AddTime),
			IsSuper:  int64(v.IsSuper),
			Role: &pb.RoleModel{
				Title:       v.Role.Title,
				Description: v.Role.Description,
			},
		})
	}
	rsp.ManagerList = tempList

	return nil
}

// ManagerAdd 增加管理员
func (e *RbacManager) ManagerAdd(ctx context.Context, req *pb.ManagerAddRequest, rsp *pb.ManagerAddResponse) error {
	//执行增加管理员
	manager := models.Manager{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
		RoleId:   int(req.RoleId),
		Status:   int(req.Status),
		AddTime:  int(req.AddTime),
	}
	err := mysql.DB.Create(&manager).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = FailedAdd
		return err
	}

	rsp.Success = true
	return nil
}

// ManagerUpdate 修改管理员
func (e *RbacManager) ManagerUpdate(ctx context.Context, req *pb.ManagerUpdateRequest, rsp *pb.ManagerUpdateResponse) error {
	//执行修改
	manager := models.Manager{Id: int64(req.Id)}
	mysql.DB.Find(&manager)
	manager.Username = req.Username
	manager.Email = req.Email
	manager.Mobile = req.Mobile
	manager.RoleId = int(req.RoleId)

	//注意：判断密码是否为空 为空表示不修改密码 不为空表示修改密码
	if req.Password != "" {
		manager.Password = req.Password
	}
	err := mysql.DB.Save(&manager).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = FailedUpdate
		return err
	}

	rsp.Success = true
	return nil
}

// ManagerDelete 删除管理员
func (e *RbacManager) ManagerDelete(ctx context.Context, req *pb.ManagerDeleteRequest, rsp *pb.ManagerDeleteResponse) error {
	manager := models.Manager{Id: int64(req.Id)}
	err := mysql.DB.Delete(&manager).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = FailedDelete
		return err
	}

	rsp.Success = false
	return nil
}
