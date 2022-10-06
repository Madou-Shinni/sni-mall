package models

// Focus 轮播图
type Focus struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	FocusType int    `json:"focusType,omitempty"`
	FocusImg  string `json:"focusImg,omitempty"`
	Link      string `json:"link,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	Status    int    `json:"status,omitempty"`
	AddTime   int    `json:"addTime,omitempty"`
}

func (Focus) TableName() string {
	return "focus"
}
