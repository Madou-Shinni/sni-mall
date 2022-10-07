package models

type UserTemp struct {
	Id        int    `json:"id,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Phone     string `json:"phone,omitempty"`
	SendCount int    `json:"sendCount,omitempty"`
	AddDay    string `json:"addDay,omitempty"`
	AddTime   int    `json:"addTime,omitempty"`
	Sign      string `json:"sign,omitempty"`
}

func (UserTemp) TableName() string {
	return "user_temp"
}
