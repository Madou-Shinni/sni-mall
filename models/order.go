package models

type Order struct {
	Id          int     `json:"id,omitempty"`
	OrderId     string  `json:"orderId,omitempty"`
	Uid         int     `json:"uid,omitempty"`
	AllPrice    float64 `json:"allPrice,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Name        string  `json:"name,omitempty"`
	Address     string  `json:"address,omitempty"`
	PayStatus   int     `json:"payStatus,omitempty"`   // 支付状态： 0 表示未支付     1 已经支付
	PayType     int     `json:"payType,omitempty"`     // 支付类型： 0 alipay    1 wechat
	OrderStatus int     `json:"orderStatus,omitempty"` // 订单状态： 0 已下单  1 已付款  2 已配货  3、发货   4、交易成功   5、退货   6、取消
	AddTime     int     `json:"addTime,omitempty"`
}

func (Order) TableName() string {
	return "order"
}
