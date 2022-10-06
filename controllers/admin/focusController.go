package admin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	FailedSort        = "请输入正确的排序"
	FailedUpload      = "上传失败！"
	FailedAddFocus    = "添加轮播图失败！"
	FailedUpdateFocus = "修改轮播图失败！"
)

type FocusController struct {
	BaseController
}

// List 轮播图列表
func (con FocusController) List(c *gin.Context) {
	var focus []models.Focus
	mysql.DB.Find(&focus)
	con.SuccessAndData(c, focus)
}

// Add 添加轮播图
func (con FocusController) Add(c *gin.Context) {
	title := c.PostForm("title")
	focusType, err := utils.StringToInt(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := utils.StringToInt(c.PostForm("sort"))
	status, err3 := utils.StringToInt(c.PostForm("status"))

	if err != nil || err3 != nil {
		con.Error(c, ParameterError)
		return
	}
	if err2 != nil {
		con.Error(c, FailedSort)
		return
	}
	// 上传图片
	focusImgSrc, uploadErr := utils.UploadImg(c, "focus_img")
	if uploadErr != nil {
		con.Error(c, FailedUpload)
		return
	}
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(utils.GetUnix()),
	}
	if sqlErr := mysql.DB.Create(&focus).Error; sqlErr != nil {
		con.Error(c, FailedAddFocus)
		return
	}

	con.Success(c)
}

// Update 修改轮播图
func (con FocusController) Update(c *gin.Context) {
	id, err := utils.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	focusType, err2 := utils.StringToInt(c.PostForm("focus_type"))
	link := strings.Trim(c.PostForm("link"), " ")
	sort, err3 := utils.StringToInt(c.PostForm("sort"))
	status, err4 := utils.StringToInt(c.PostForm("status"))

	if err != nil || err2 != nil || err4 != nil {
		con.Error(c, ParameterError)
		return
	}
	if err3 != nil {
		con.Error(c, FailedSort)
		return
	}
	focusImg, uploadErr := utils.UploadImg(c, "focus_img")
	if uploadErr != nil {
		con.Error(c, uploadErr.Error())
		return
	}

	focus := models.Focus{Id: id}
	mysql.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" { // 如果上传的图片不为空再保存
		focus.FocusImg = focusImg
	}
	if sqlErr := mysql.DB.Save(&focus).Error; sqlErr != nil {
		con.Error(c, FailedUpdateFocus)
		return
	}

	con.Success(c)
}

// Delete 删除轮播图
func (con FocusController) Delete(c *gin.Context) {
	id, err := utils.StringToInt(c.PostForm("id"))
	if err != nil {
		con.Error(c, ParameterError)
		return
	}
	focus := models.Focus{Id: id}
	mysql.DB.Delete(&focus)
	// 根据需求 是否需要删除图片
	//os.Remove("static/upload/20210915/1631677671.jpg")
	con.Success(c)
}
