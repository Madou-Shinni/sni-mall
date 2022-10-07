package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
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
	FailedAdd       = "添加失败！"
	FailedUpdate    = "修改失败！"
	DefaultPageNum  = "1"  // 默认当前页为第1页
	DefaultPageSize = "10" // 默认每页查询的数量10条
	MaxPageSize     = 1000 // 最大限制数量1000条
)

var wg sync.WaitGroup

// List 商品列表（分页）
func (con GoodsController) List(c *gin.Context) {
	// 获取当前页
	pageNum, _ := utils.StringToInt(c.DefaultQuery("pageNum", DefaultPageNum))
	// 每页查询的数量
	pageSize, _ := utils.StringToInt(c.DefaultQuery("pageSize", DefaultPageSize))
	if pageSize > 1000 {
		pageSize = MaxPageSize
	}
	var goodsList []models.Goods
	// 分页查询商品列表
	mysql.DB.Where("is_delete = 0").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&goodsList)
	// 获取总数量
	var count int64
	mysql.DB.Table("goods").Count(&count)

	goodsPage := make(map[string]interface{})
	goodsPage["goodsList"] = goodsList
	// 注意：必须使用float64类型
	goodsPage["totalPages"] = math.Ceil(float64(count) / float64(pageSize)) // 总数÷每页数量 向上取整得到总页数

	con.SuccessAndData(c, goodsPage)
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
	if len(goodsImg) > 0 {
		// 不是oss上传时才生成缩略图
		if utils.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				utils.ResizeGoodsImage(goodsImg)
				wg.Done()
			}()
		}
	}

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

// Goods 获取商品
func (con GoodsController) Goods(c *gin.Context) {
	// 1、获取要修改的商品数据
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	goods := models.Goods{Id: id}
	mysql.DB.Find(&goods)

	// 2、获取商品分类
	var goodsCateList []models.GoodsCate
	mysql.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	// 3、获取所有颜色 以及选中的颜色
	goodsColorSlice := strings.Split(goods.GoodsColor, ",")
	goodsColorMap := make(map[string]string)
	for _, v := range goodsColorSlice {
		goodsColorMap[v] = v
	}

	var goodsColorList []models.GoodsColor
	mysql.DB.Find(&goodsColorList)
	for i := 0; i < len(goodsColorList); i++ {
		if _, ok := goodsColorMap[utils.IntToString(goodsColorList[i].Id)]; ok {
			goodsColorList[i].Checked = true
		}
	}

	// 4、商品的图库信息
	var goodsImageList []models.GoodsImage
	mysql.DB.Where("goods_id=?", goods.Id).Find(&goodsImageList)

	// 5、获取商品类型
	var goodsTypeList []models.GoodsType
	mysql.DB.Find(&goodsTypeList)

	// 6、获取规格信息
	var goodsAttr []models.GoodsAttr
	mysql.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
	goodsAttrStr := ""

	for _, v := range goodsAttr {
		if v.AttributeType == 1 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: </span> <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else {
			//获取当前类型对应的值
			goodsTypeAttribute := models.GoodsTypeAttribute{Id: v.AttributeId}
			mysql.DB.Find(&goodsTypeAttribute)
			attrValueSlice := strings.Split(goodsTypeAttribute.AttrValue, "\n")

			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += fmt.Sprintf(`<select name="attr_value_list">`)
			for i := 0; i < len(attrValueSlice); i++ {
				if attrValueSlice[i] == v.AttributeValue {
					goodsAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[i], attrValueSlice[i])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[i], attrValueSlice[i])
				}
			}
			goodsAttrStr += fmt.Sprintf(`</select>`)
			goodsAttrStr += fmt.Sprintf(`</li>`)

		}
	}

	goodsInfoMap := make(map[string]interface{})
	goodsInfoMap["goods"] = goods
	goodsInfoMap["goodsCateList"] = goodsCateList
	goodsInfoMap["goodsColorList"] = goodsColorList
	goodsInfoMap["goodsTypeList"] = goodsTypeList
	goodsInfoMap["goodsAttrStr"] = goodsAttrStr
	goodsInfoMap["goodsImageList"] = goodsImageList
	con.SuccessAndData(c, goodsInfoMap)
}

// Update 修改商品
func (con GoodsController) Update(c *gin.Context) {
	//1、获取表单提交过来的数据
	id, err1 := utils.StringToInt(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, ParameterError)
		return
	}
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

	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")
	//3、修改数据
	goods := models.Goods{Id: id}
	mysql.DB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	//4、上传图片   生成缩略图
	goodsImg, err2 := utils.UploadImg(c, "goods_img")
	if err2 == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
		// 不是oss上传时才生成缩略图
		if utils.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				utils.ResizeGoodsImage(goodsImg)
				wg.Done()
			}()
		}
	}

	err3 := mysql.DB.Save(&goods).Error
	if err3 != nil {
		con.Error(c, FailedUpdate)
		return
	}

	//5、修改图库 增加图库信息
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
	//6、修改规格包装  1、删除当前商品下面的规格包装   2、重新执行增加

	// 6.1删除当前商品下面的规格包装
	goodsAttrObj := models.GoodsAttr{}
	mysql.DB.Where("goods_id=?", goods.Id).Delete(&goodsAttrObj)
	//6.2、重新执行增加
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
	wg.Wait()
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
