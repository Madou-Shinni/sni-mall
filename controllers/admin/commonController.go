package admin

import (
	"github.com/gin-gonic/gin"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	FailedUpdateStatus = "修改状态失败！"
)

type CommonController struct {
	BaseController
}

// ChangeStatus 通用接收ajax请求修改状态status
func (con CommonController) ChangeStatus(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id")) // 主键id
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	table := c.Query("table") // 表
	field := c.Query("field") // 字段（status）

	// status = ABS(0-1)   =   1
	// status = ABS(1-1)   =   0

	sqlErr := mysql.DB.Exec("update ? set ? = ABS(?-1) where id = ?", table, field, field, id).Error
	if sqlErr != nil {
		con.Error(c, FailedUpdateStatus)
		return
	}

	con.Success(c)
}
