package admin

import "github.com/gin-gonic/gin"

type AdminController struct {
	BaseController
}

// List 管理员列表
func (con AdminController) List(c *gin.Context) {
	con.Success(c)
}

// Add 添加管理员
func (con AdminController) Add(c *gin.Context) {
	con.Success(c)
}

// Update 修改管理员
func (con AdminController) Update(c *gin.Context) {
	con.Success(c)
}

// Delete 删除管理员
func (con AdminController) Delete(c *gin.Context) {
	con.Success(c)
}
