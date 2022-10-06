package models

type GoodsColor struct {
	Id         int    `json:"id,omitempty"`
	ColorName  string `json:"colorName,omitempty"`
	ColorValue string `json:"colorValue,omitempty"`
	Status     int    `json:"status,omitempty"`
	Checked    bool   `gorm:"_"`
}

func (GoodsColor) TableName() string {
	return "goods_color"
}
