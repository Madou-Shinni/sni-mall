package models

type Goods struct {
	Id            int     `json:"id,omitempty"`
	Title         string  `json:"title,omitempty"`
	SubTitle      string  `json:"subTitle,omitempty"`
	GoodsSn       string  `json:"goodsSn,omitempty"`
	CateId        int     `json:"cateId,omitempty"`
	ClickCount    int     `json:"clickCount,omitempty"`
	GoodsNumber   int     `json:"goodsNumber,omitempty"`
	Price         float64 `json:"price,omitempty"`
	MarketPrice   float64 `json:"marketPrice,omitempty"`
	RelationGoods string  `json:"relationGoods,omitempty"`
	GoodsAttr     string  `json:"goodsAttr,omitempty"`
	GoodsVersion  string  `json:"goodsVersion,omitempty"`
	GoodsImg      string  `json:"goodsImg,omitempty"`
	GoodsGift     string  `json:"goodsGift,omitempty"`
	GoodsFitting  string  `json:"goodsFitting,omitempty"`
	GoodsColor    string  `json:"goodsColor,omitempty"`
	GoodsKeywords string  `json:"goodsKeywords,omitempty"`
	GoodsDesc     string  `json:"goodsDesc,omitempty"`
	GoodsContent  string  `json:"goodsContent,omitempty"`
	IsDelete      int     `json:"isDelete,omitempty"`
	IsHot         int     `json:"isHot,omitempty"`
	IsBest        int     `json:"isBest,omitempty"`
	IsNew         int     `json:"isNew,omitempty"`
	GoodsTypeId   int     `json:"goodsTypeId,omitempty"`
	Sort          int     `json:"sort,omitempty"`
	Status        int     `json:"status,omitempty"`
	AddTime       int     `json:"addTime,omitempty"`
}

func (Goods) TableName() string {
	return "goods"
}
