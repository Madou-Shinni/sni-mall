package models

type GoodsAttr struct {
	Id              int    `json:"id,omitempty"`
	GoodsId         int    `json:"goodsId,omitempty"`
	AttributeCateId int    `json:"attributeCateId,omitempty"`
	AttributeId     int    `json:"attributeId,omitempty"`
	AttributeTitle  string `json:"attributeTitle,omitempty"`
	AttributeType   int    `json:"attributeType,omitempty"`
	AttributeValue  string `json:"attributeValue,omitempty"`
	Sort            int    `json:"sort,omitempty"`
	AddTime         int    `json:"addTime,omitempty"`
	Status          int    `json:"status,omitempty"`
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
