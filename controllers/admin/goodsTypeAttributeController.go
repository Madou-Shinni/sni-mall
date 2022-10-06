package admin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	FailedAddGoodsTypeAttribute    = "添加商品类型属性失败！"
	FailedUpdateGoodsTypeAttribute = "修改商品类型属性失败！"
)

type GoodsTypeAttributeController struct {
	BaseController
}

// List 商品类型属性列表
func (con GoodsTypeAttributeController) List(c *gin.Context) {
	con.Success(c)
}

// ListById 商品类型属性
func (con GoodsTypeAttributeController) ListById(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	var attributeList models.GoodsTypeAttribute
	mysql.DB.Where("id = ?", id).Find(&attributeList)
	con.SuccessAndData(c, attributeList)
}

// Add 添加商品类型属性
func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err1 := utils.StringToInt(c.PostForm("cate_id"))
	attrType, err2 := utils.StringToInt(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := utils.StringToInt(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(c, ParameterError)
		return
	}
	if title == "" {
		con.Error(c, TitleIsEmpty)
		return
	}

	if err3 != nil {
		con.Error(c, FailedSort)
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(utils.GetUnix()),
	}
	err := mysql.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, FailedAddGoodsTypeAttribute)
		return
	}
	con.Success(c)
}

// Update 修改商品类型属性
func (con GoodsTypeAttributeController) Update(c *gin.Context) {
	id, err1 := utils.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err2 := utils.StringToInt(c.PostForm("cate_id"))
	attrType, err3 := utils.StringToInt(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err4 := utils.StringToInt(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, ParameterError)
		return
	}
	if title == "" {
		con.Error(c, TitleIsEmpty)
		return
	}
	if err4 != nil {
		con.Error(c, FailedSort)
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	mysql.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := mysql.DB.Save(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, FailedUpdateGoodsTypeAttribute)
		return
	}
	con.Success(c)
}

// Delete 删除商品类型属性
func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	id, err1 := utils.StringToInt(c.Query("id"))
	if err1 != nil {
		con.Error(c, ParameterError)
		return
	}
	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	mysql.DB.Delete(&goodsTypeAttr)
	con.Success(c)
}
