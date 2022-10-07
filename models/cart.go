package models

type Cart struct {
	Id           int     `json:"id,omitempty"`
	Title        string  `json:"title,omitempty"`
	Price        float64 `json:"price,omitempty"`
	GoodsVersion string  `json:"goodsVersion,omitempty"`
	Uid          int     `json:"uid,omitempty"`
	Num          int     `json:"num,omitempty"`
	GoodsGift    string  `json:"goodsGift,omitempty"`
	GoodsFitting string  `json:"goodsFitting,omitempty"`
	GoodsColor   string  `json:"goodsColor,omitempty"`
	GoodsImg     string  `json:"goodsImg,omitempty"`
	GoodsAttr    string  `json:"goodsAttr,omitempty"`
	Checked      bool    `json:"checked,omitempty"`
}

// HasCartData 判断购物车里面有没有当前数据
func HasCartData(cartList []Cart, currentData Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
			return true
		}
	}
	return false
}
