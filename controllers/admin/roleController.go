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
	TopModuleId       = 0 // 顶层模块
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

// Auth 角色授权
func (con RoleController) Auth(c *gin.Context) {
	// 获取角色id
	roleId, err := utils.StringToInt(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	// 先删除现在有的权限（避免表中有多个相同的权限）
	var roleAccessList []models.RoleAccess
	var roleAccess models.RoleAccess
	mysql.DB.Select("id").Where("role_id = ?", roleId).Find(&roleAccessList) // 获取角色对应的权限
	var ids []int                                                            // 用来存放所有角色对应权限的主键
	for _, v := range roleAccessList {
		ids = append(ids, v.Id)
	}
	mysql.DB.Where("id in (?)", ids).Delete(&roleAccess) // 批量删除权限

	// 获取权限id（多个）
	accessIds := c.PostFormArray("access_node[]")
	for _, value := range accessIds {
		intValue, _ := utils.StringToInt(value)
		roleAccess = models.RoleAccess{RoleId: roleId, AccessId: intValue}
		roleAccessList = append(roleAccessList, roleAccess)
	}
	mysql.DB.Create(&roleAccessList) // 批量插入权限
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

	// 2.获取所有权限
	var accessList []models.Access
	mysql.DB.Where("module_id = ?", TopModuleId).Preload("AccessItem").Find(&accessList)

	// 3.获取当前角色拥有的权限，并把权限id放在一个map对象里
	var roleAccessList []models.RoleAccess
	mysql.DB.Where("role_id = ?", roleId).Find(&roleAccessList)
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
}
