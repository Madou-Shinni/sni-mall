package models

type RoleAccess struct {
	Id       int
	AccessId int
	RoleId   int
}

func (RoleAccess) TableName() string {
	return "role_access"
}
