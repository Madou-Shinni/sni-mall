package handler

import (
	"context"
	mysql "rbac/models/mysql"
	"strconv"

	"rbac/models"
	pb "rbac/proto/rbacAccess"
)

type RbacAccess struct{}

// AccessGet 获取权限
func (e *RbacAccess) AccessGet(ctx context.Context, req *pb.AccessGetRequest, res *pb.AccessGetResponse) error {
	var accessList []models.Access
	where := "1=1"
	if req.Id > 0 {
		where += " AND id=" + strconv.Itoa(int(req.Id))
	} else {
		where += " AND module_id = 0"
	}
	mysql.DB.Where(where).Preload("AccessItem").Find(&accessList)

	//处理数据
	var tempList []*pb.AccessModel
	for _, v := range accessList {
		var tempItemList []*pb.AccessModel
		for _, k := range v.AccessItem {
			tempItemList = append(tempItemList, &pb.AccessModel{
				Id:          int64(k.Id),
				ModuleName:  k.ModuleName,
				ActionName:  k.ActionName,
				Type:        int64(k.Type),
				Url:         k.Url,
				ModuleId:    int64(k.ModuleId),
				Sort:        int64(k.Sort),
				Description: k.Description,
				Status:      int64(k.Status),
				AddTime:     int64(k.AddTime),
			})
		}
		tempList = append(tempList, &pb.AccessModel{
			Id:          int64(v.Id),
			ModuleName:  v.ModuleName,
			ActionName:  v.ActionName,
			Type:        int64(v.Type),
			Url:         v.Url,
			ModuleId:    int64(v.ModuleId),
			Sort:        int64(v.Sort),
			Description: v.Description,
			Status:      int64(v.Status),
			AddTime:     int64(v.AddTime),
			AccessItem:  tempItemList,
		})
	}

	res.AccessList = tempList
	return nil
}

// AccessAdd 增加权限
func (e *RbacAccess) AccessAdd(ctx context.Context, req *pb.AccessAddRequest, res *pb.AccessAddResponse) error {
	access := models.Access{
		ModuleName:  req.ModuleName,
		Type:        int(req.Type),
		ActionName:  req.ActionName,
		Url:         req.Url,
		ModuleId:    int(req.ModuleId),
		Sort:        int(req.Sort),
		Description: req.Description,
		Status:      int(req.Status),
		AddTime:     int(req.AddTime),
	}
	err := mysql.DB.Create(&access).Error
	if err != nil {
		res.Success = false
		res.Message = FailedAdd
	} else {
		res.Success = true
		res.Message = "增加数据成功"
	}
	return err
}

// AccessUpdate 修改权限
func (e *RbacAccess) AccessUpdate(ctx context.Context, req *pb.AccessUpdateRequest, res *pb.AccessUpdateResponse) error {
	access := models.Access{Id: int(req.Id)}
	mysql.DB.Find(&access)

	access.ModuleName = req.ModuleName
	access.Type = int(req.Type)
	access.ActionName = req.ActionName
	access.Url = req.Url
	access.ModuleId = int(req.ModuleId)
	access.Sort = int(req.Sort)
	access.Description = req.Description
	access.Status = int(req.Status)

	err := mysql.DB.Save(&access).Error
	if err != nil {
		res.Success = false
		res.Message = FailedUpdate
		return err
	}

	res.Success = true
	return nil
}

// AccessDelete 删除权限
func (e *RbacAccess) AccessDelete(ctx context.Context, req *pb.AccessDeleteRequest, res *pb.AccessDeleteResponse) error {

	access := models.Access{Id: int(req.Id)}
	mysql.DB.Find(&access)
	if access.ModuleId == 0 { //顶级模块
		var accessList []models.Access
		mysql.DB.Where("module_id = ?", access.Id).Find(&accessList)

		if len(accessList) > 0 {
			res.Success = false
			res.Message = "当前模块下面有菜单或者操作，请删除菜单或者操作以后再来删除这个数据"
		} else {
			mysql.DB.Delete(&access)
			res.Success = true
			res.Message = "删除数据成功"
		}
	}
	//操作 或者菜单
	mysql.DB.Delete(&access)
	res.Success = true
	res.Message = "删除数据成功"
	return nil
}
