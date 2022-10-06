package models

type GoodsImage struct {
	Id      int    `json:"id,omitempty"`
	GoodsId int    `json:"goodsId,omitempty"`
	ImgUrl  string `json:"imgUrl,omitempty"`
	ColorId int    `json:"colorId,omitempty"`
	Sort    int    `json:"sort,omitempty"`
	AddTime int    `json:"addTime,omitempty"`
	Status  int    `json:"status,omitempty"`
}

func (GoodsImage) TableName() string {
	return "goods_image"
}
