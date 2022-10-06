package models

type GoodsTypeAttribute struct {
	Id        int    `json:"id,omitempty"`
	CateId    int    `json:"cateId,omitempty"`
	Title     string `json:"title,omitempty"`
	AttrType  int    `json:"attrType,omitempty"`
	AttrValue string `json:"attrValue,omitempty"`
	Status    int    `json:"status,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	AddTime   int    `json:"addTime,omitempty"`
}

func (GoodsTypeAttribute) TableName() string {
	return "goods_type_attribute"
}
