package index

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	DefaultPageSize = 5
)

type ProductController struct {
	BaseController
}

// Category 获取分类数据
func (con ProductController) Category(c *gin.Context) {
	cateId, _ := utils.StringToInt(c.Param("id"))
	//当前页
	page, _ := utils.StringToInt(c.Query("page"))
	// 获取当前分类
	currentCate := models.GoodsCate{}
	mysql.DB.Where("id = ?", cateId).Find(&currentCate)
	var subCate []models.GoodsCate
	var tempSlice []int
	if currentCate.Pid == 0 {
		// 获取二级分类
		mysql.DB.Where("pid = ?", currentCate.Id).Find(&subCate)
		for i := 0; i < len(subCate); i++ {
			tempSlice = append(tempSlice, subCate[i].Id)
		}
	} else {
		// 兄弟分类
		mysql.DB.Where("pid=?", currentCate.Pid).Find(&subCate)
	}

	tempSlice = append(tempSlice, cateId)
	where := "cate_id in ?"
	var goodsList []models.Goods
	mysql.DB.Where(where, tempSlice).Offset((page - 1) * DefaultPageSize).Limit(DefaultPageSize).Find(&goodsList)

	//获取总数量
	var count int64
	mysql.DB.Where(where, tempSlice).Table("goods").Count(&count)

	var dataMap map[string]interface{}

	con.Render(c, dataMap)
	c.JSON(http.StatusOK, gin.H{
		"page":        page,
		"goodsList":   goodsList,
		"subCate":     subCate,
		"currentCate": currentCate,
		"totalPages":  math.Ceil(float64(count) / float64(DefaultPageSize)),
	})
}

// Detail 商品详情
func (con ProductController) Detail(c *gin.Context) {

	id, err := utils.StringToInt(c.Query("id"))

	if err != nil {
		c.String(302, "业务繁忙！")
		return
	}

	//1、获取商品信息
	goods := models.Goods{Id: id}
	mysql.DB.Find(&goods)

	//2、获取关联商品  RelationGoods
	var relationGoods []models.Goods
	goods.RelationGoods = strings.ReplaceAll(goods.RelationGoods, "，", ",")
	relationIds := strings.Split(goods.RelationGoods, ",")

	mysql.DB.Where("id in ?", relationIds).Select("id,title,price,goods_version").Find(&relationGoods)

	//3、获取关联赠品 GoodsGift

	var goodsGift []models.Goods
	goods.GoodsGift = strings.ReplaceAll(goods.GoodsGift, "，", ",")
	giftIds := strings.Split(goods.GoodsGift, ",")
	mysql.DB.Where("id in ?", giftIds).Select("id,title,price,goods_version").Find(&goodsGift)

	//4、获取关联颜色 GoodsColor
	var goodsColor []models.GoodsColor
	goods.GoodsColor = strings.ReplaceAll(goods.GoodsColor, "，", ",")
	colorIds := strings.Split(goods.GoodsColor, ",")
	mysql.DB.Where("id in ?", colorIds).Find(&goodsColor)

	//5、获取关联配件 GoodsFitting
	var goodsFitting []models.Goods
	goods.GoodsFitting = strings.ReplaceAll(goods.GoodsFitting, "，", ",")
	fittingIds := strings.Split(goods.GoodsFitting, ",")
	mysql.DB.Where("id in ?", fittingIds).Select("id,title,price,goods_version").Find(&goodsFitting)

	//6、获取商品关联的图片 GoodsImage
	var goodsImage []models.GoodsImage
	mysql.DB.Where("goods_id=?", goods.Id).Limit(6).Find(&goodsImage)

	//7、获取规格参数信息 GoodsAttr
	var goodsAttr []models.GoodsAttr
	mysql.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)

	con.Render(c, gin.H{
		"goods":         goods,
		"relationGoods": relationGoods,
		"goodsGift":     goodsGift,
		"goodsColor":    goodsColor,
		"goodsFitting":  goodsFitting,
		"goodsImage":    goodsImage,
		"goodsAttr":     goodsAttr,
	})
}
