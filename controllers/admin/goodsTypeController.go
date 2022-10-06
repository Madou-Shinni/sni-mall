package admin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	FailedAddGoodsType    = "添加商品类型失败！"
	FailedUpdateGoodsType = "修改商品类型失败！"
)

type GoodsTypeController struct {
	BaseController
}

// List 商品类型列表
func (con GoodsTypeController) List(c *gin.Context) {
	var goodsTypeList []models.GoodsType
	mysql.DB.Find(&goodsTypeList)
	con.SuccessAndData(c, goodsTypeList)
}

// Add 添加商品类型
func (con GoodsTypeController) Add(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err1 := utils.StringToInt(c.PostForm("status"))

	if err1 != nil {
		con.Error(c, ParameterError)
		return
	}

	if title == "" {
		con.Error(c, TitleIsEmpty)
		return
	}
	goodsType := models.GoodsType{
		Title:       title,
		Description: description,
		Status:      status,
		AddTime:     int(utils.GetUnix()),
	}

	err := mysql.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(c, FailedAddGoodsType)
		return
	}
	con.Success(c)
}

// Update 修改商品类型
func (con GoodsTypeController) Update(c *gin.Context) {
	id, err1 := utils.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err2 := utils.StringToInt(c.PostForm("status"))
	if err1 != nil || err2 != nil {
		con.Error(c, ParameterError)
		return
	}

	if title == "" {
		con.Error(c, TitleIsEmpty)
		return
	}
	goodsType := models.GoodsType{Id: id}
	mysql.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err3 := mysql.DB.Save(&goodsType).Error
	if err3 != nil {
		con.Error(c, FailedUpdateGoodsType)
		return
	}
	con.Success(c)
}

// Delete 删除商品类型
func (con GoodsTypeController) Delete(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	goodsType := models.GoodsType{Id: id}
	mysql.DB.Delete(&goodsType)
	con.Success(c)
}
