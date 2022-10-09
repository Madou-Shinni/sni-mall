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

// RoleGetAuth 角色查询授权
func (e *RbacRole) RoleGetAuth(ctx context.Context, req *pb.RoleGetAuthRequest, rsp *pb.RoleGetAuthResponse) error {
	// 1.获取角色id req.RoleId

	// 2.获取所有权限
	var accessList []models.Access
	mysql.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)

	// 3.获取当前角色拥有的权限，并把权限id放在一个map对象里
	var roleAccessList []models.RoleAccess
	mysql.DB.Where("role_id = ?", req.RoleId).Find(&roleAccessList)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccessList {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	// 4.遍历所有的权限，判断当前权限的id是否在角色权限的map对象中，如果是就给当前对象加一个checked属性
	// 注意：for range 无法直接修改，使用for i
	for i := 0; i < len(accessList); i++ { // 遍历顶级模块
		if _, ok := roleAccessMap[accessList[i].Id]; ok { // 如果存在checked=true
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ { // 遍历二级模块
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok { // 如果存在checked=true
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}

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
				Checked:     k.Checked,
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
			Checked:     v.Checked,
			AccessItem:  tempItemList,
		})
	}

	rsp.AccessList = tempList

	return nil
}

// RoleAuth 角色授权
func (e *RbacRole) RoleAuth(ctx context.Context, req *pb.RoleAuthRequest, rsp *pb.RoleAuthResponse) error {
	//1、删除当前角色对应的权限
	roleAccess := models.RoleAccess{}
	mysql.DB.Where("role_id=?", req.RoleId).Delete(&roleAccess)

	//2、增加当前角色对应的权限
	for _, v := range req.AccessIds {
		roleAccess.RoleId = int(req.RoleId)
		accessId, _ := strconv.Atoi(v)
		roleAccess.AccessId = accessId
		mysql.DB.Create(&roleAccess)
	}

	rsp.Success = true

	return nil
}

// MiddlewaresAuth 中间件权限判断
func (e *RbacRole) MiddlewaresAuth(ctx context.Context, req *pb.MiddlewaresAuthRequest, rsp *pb.MiddlewaresAuthResponse) error {
	// 1.获取用户信息（角色id） req.RoleId

	// 2.获取当前用户的权限id列表
	var roleAccessList []models.RoleAccess
	var access models.Access
	// 把权限id放在一个map类型的对象里面
	roleAccessMap := make(map[int]int)
	mysql.DB.Select("access_id").Where("role_id = ?", req.RoleId).Find(&roleAccessList)
	for _, v := range roleAccessList {
		roleAccessMap[v.AccessId] = v.AccessId
	}
	// 查询url对应的权限id
	mysql.DB.Select("id").Where("url = ?", req.Url).Find(&access)
	// 3.匹配当前用户是否有访问当前路由的权限
	if _, ok := roleAccessMap[access.Id]; !ok {
		rsp.HasPermission = false
	} else {
		rsp.HasPermission = true
	}
	return nil
}
