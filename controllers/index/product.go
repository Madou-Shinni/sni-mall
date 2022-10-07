package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaomi-mall/models/utils"
)

type ProductController struct {
	BaseController
}

// Category 获取分类数据
func (con ProductController) Category(c *gin.Context) {
	id, _ := utils.StringToInt(c.Param("id"))
	fmt.Println(id)

	var dataMap map[string]interface{}

	con.Render(c, dataMap)
}
