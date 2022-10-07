package index

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
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
