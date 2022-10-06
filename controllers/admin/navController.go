package admin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

type NavController struct {
	BaseController
}

// List 导航列表
func (con NavController) List(c *gin.Context) {
	var navList []models.Nav
	mysql.DB.Find(&navList)
	con.SuccessAndData(c, navList)
}

// Page 导航分页列表
func (con NavController) Page(c *gin.Context) {
	pageNum, _ := utils.StringToInt(c.DefaultQuery("pageNum", DefaultPageNum))
	pageSize, _ := utils.StringToInt(c.DefaultQuery("pageNum", "8"))
	if pageSize > 1000 {
		pageSize = MaxPageSize
	}
	var navList []models.Nav
	mysql.DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&navList)
	con.SuccessAndData(c, navList)
}

// Add 添加导航
func (con NavController) Add(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	link := strings.Trim(c.PostForm("link"), " ")
	position, err2 := utils.StringToInt(c.PostForm("position"))
	isOpennew, err3 := utils.StringToInt(c.PostForm("isOpennew"))
	relation := strings.Trim(c.PostForm("relation"), " ")
	sort, err4 := utils.StringToInt(c.PostForm("sort"))
	status, err5 := utils.StringToInt(c.PostForm("status"))
	if err2 != nil || err3 != nil || err5 != nil {
		con.Error(c, ParameterError)
		return
	}
	if err4 != nil {
		con.Error(c, FailedSort)
		return
	}
	nav := models.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(utils.GetUnix()),
	}
	sqlErr := mysql.DB.Create(nav).Error
	if sqlErr != nil {
		con.Error(c, FailedAdd)
		return
	}
	con.Success(c)
}

// Update 修改导航
func (con NavController) Update(c *gin.Context) {
	id, err1 := utils.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	link := strings.Trim(c.PostForm("link"), " ")
	position, err2 := utils.StringToInt(c.PostForm("position"))
	isOpennew, err3 := utils.StringToInt(c.PostForm("isOpennew"))
	relation := strings.Trim(c.PostForm("relation"), " ")
	sort, err4 := utils.StringToInt(c.PostForm("sort"))
	status, err5 := utils.StringToInt(c.PostForm("status"))
	if err1 != nil || err2 != nil || err3 != nil || err5 != nil {
		con.Error(c, ParameterError)
		return
	}
	if err4 != nil {
		con.Error(c, FailedSort)
		return
	}
	nav := models.Nav{Id: id}
	mysql.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status
	sqlErr := mysql.DB.Save(nav).Error
	if sqlErr != nil {
		con.Error(c, FiledUpdate)
		return
	}
	con.Success(c)
}

// Delete 删除导航
func (con NavController) Delete(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	nav := models.Nav{Id: id}
	mysql.DB.Delete(&nav)
	con.Success(c)
}
