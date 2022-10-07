package index

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/models/utils"
)

type ProductController struct {
}

// Category 获取分类数据
func (con ProductController) Category(c *gin.Context) {
	id, _ := utils.StringToInt(c.Param("id"))
}
