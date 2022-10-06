package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

type GoodsController struct {
	BaseController
}

const (
	FailedAdd = "添加失败！"
)

var wg sync.WaitGroup

// List 商品列表
func (con GoodsController) List(c *gin.Context) {
	con.Success(c)
}

// GoodsInfoList 商品信息列表（商品分类、颜色、类型列表）
func (con GoodsController) GoodsInfoList(c *gin.Context) {
	//获取商品分类
	var goodsCateList []models.GoodsCate
	mysql.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取所有颜色信息
	var goodsColorList []models.GoodsColor
	mysql.DB.Find(&goodsColorList)

	//获取商品规格包装
	var goodsTypeList []models.GoodsType
	mysql.DB.Find(&goodsTypeList)

	var goodsInfoMap map[string]interface{}
	goodsInfoMap["goodsCateList"] = goodsCateList
	goodsInfoMap["goodsColorList"] = goodsCateList
	goodsInfoMap["goodsTypeList"] = goodsCateList

	con.SuccessAndData(c, goodsInfoMap)
}

// Add 添加商品
func (con GoodsController) Add(c *gin.Context) {
	//1、获取表单提交过来的数据 进行判断
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := utils.StringToInt(c.PostForm("cate_id"))
	goodsNumber, _ := utils.StringToInt(c.PostForm("goods_number"))
	//注意小数点
	marketPrice, _ := utils.Float64(c.PostForm("market_price"))
	price, _ := utils.Float64(c.PostForm("price"))

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")

	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := utils.StringToInt(c.PostForm("is_delete"))
	isHot, _ := utils.StringToInt(c.PostForm("is_hot"))
	isBest, _ := utils.StringToInt(c.PostForm("is_best"))
	isNew, _ := utils.StringToInt(c.PostForm("is_new"))
	goodsTypeId, _ := utils.StringToInt(c.PostForm("goods_type_id"))
	sort, _ := utils.StringToInt(c.PostForm("sort"))
	status, _ := utils.StringToInt(c.PostForm("status"))
	addTime := int(utils.GetUnix())

	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")

	//3、上传图片   生成缩略图
	goodsImg, _ := utils.UploadImg(c, "goods_img")

	//4、增加商品数据

	goods := models.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}
	err := mysql.DB.Create(&goods).Error
	if err != nil {
		con.Error(c, FailedAdd)
		return
	}

	//5、增加图库 信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(utils.GetUnix())
			mysql.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	//6、增加规格包装
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := utils.StringToInt(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				mysql.DB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := models.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i]
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(utils.GetUnix())
				mysql.DB.Create(&goodsAttrObj)
			}

		}
		wg.Done()
	}()
	wg.Wait() // 等待两个协程执行完成
	con.Success(c)
}

// Update 修改商品
func (con GoodsController) Update(c *gin.Context) {
	con.Success(c)
}

// Delete 删除商品
func (con GoodsController) Delete(c *gin.Context) {
	con.Success(c)
}

// ImageUpload 上传图片
func (con GoodsController) ImageUpload(c *gin.Context) {
	//上传图片
	imgDir, err := utils.UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"link": "/" + imgDir,
		})
	}
}
