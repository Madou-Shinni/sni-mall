package models

type GoodsType struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      int    `json:"status,omitempty"`
	AddTime     int    `json:"addTime,omitempty"`
}

func (GoodsType) TableName() string {
	return "goods_type"
}
