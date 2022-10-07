package models

import mysql "xiaomi-mall/models/mysql"

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

/*
根据商品分类获取推荐商品
	@param {Number} cateId - 分类id
	@param {String} goodsType -  hot  best  new all
	@param {Number} limitNum -  数量


	1  表示顶级分类
		21
		23
		24


*/

func GetGoodsByCategory(cateId int, goodsType string, limitNum int) []Goods {

	//判断cateId 是否是顶级分类
	goodsCate := GoodsCate{Id: cateId}
	mysql.DB.Find(&goodsCate)
	var tempSlice []int
	if goodsCate.Pid == 0 { //顶级分类
		//获取顶级分类下面的二级分类
		var goodsCateList []GoodsCate
		mysql.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)

		for i := 0; i < len(goodsCateList); i++ {
			tempSlice = append(tempSlice, goodsCateList[i].Id)
		}

	}
	tempSlice = append(tempSlice, cateId)

	var goodsList []Goods
	where := "cate_id in ?"
	switch goodsType {
	case "hot":
		where += " AND is_hot=1"
	case "best":
		where += " AND is_best=1"
	case "new":
		where += " AND is_new=1"
	default:
		break
	}

	mysql.DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Limit(limitNum).Order("sort desc").Find(&goodsList)
	return goodsList
}
