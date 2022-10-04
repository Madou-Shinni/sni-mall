package models

type Manager struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   int    `json:"status,omitempty"`
	RoleId   int    `json:"roleId,omitempty"`
	AddTime  int    `json:"addTime,omitempty"`
	IsSuper  int    `json:"isSuper,omitempty"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleId;references:Id"`
}

func (Manager) TableName() string {
	return "manager"
}
