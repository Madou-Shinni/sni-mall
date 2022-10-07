package models

type User struct {
	Id       int    `json:"id,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	AddTime  int    `json:"addTime,omitempty"`
	LastIp   string `json:"lastIp,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   int    `json:"status,omitempty"`
}

func (User) TableName() string {
	return "user"
}
