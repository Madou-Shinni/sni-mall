package models

type OrderItem struct {
	Id           int     `json:"id,omitempty"`
	OrderId      int     `json:"orderId,omitempty"`
	Uid          int     `json:"uid,omitempty"`
	ProductTitle string  `json:"productTitle,omitempty"`
	ProductId    int     `json:"productId,omitempty"`
	ProductImg   string  `json:"productImg,omitempty"`
	ProductPrice float64 `json:"productPrice,omitempty"`
	ProductNum   int     `json:"productNum,omitempty"`
	GoodsVersion string  `json:"goodsVersion,omitempty"`
	GoodsColor   string  `json:"goodsColor,omitempty"`
	AddTime      int     `json:"addTime,omitempty"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
