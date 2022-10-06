package admin

import (
	"github.com/gin-gonic/gin"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/utils"
)

const (
	TopCate                              = 0 // 顶级分类
	FailedAddGoodsCate                   = "添加商品分类失败！"
	FailedUpdateGoodsCate                = "修改商品分类失败！"
	FailedDeleteChildrenGoodsCateAtFirst = "删除失败，请先删除子分类！"
)

type GoodsCateController struct {
	BaseController
}

// TopList 商品顶级分类列表
func (con GoodsCateController) TopList(c *gin.Context) {
	// 顶级分类
	var goodsCateList []models.GoodsCate
	mysql.DB.Where("pid = ?", TopCate).Find(&goodsCateList)
	con.SuccessAndData(c, goodsCateList)
}

// List 商品分类列表
func (con GoodsCateController) List(c *gin.Context) {
	var goodsCateList []models.GoodsCate
	mysql.DB.Preload("GoodsCateItems").Find(&goodsCateList)
	con.SuccessAndData(c, goodsCateList)
}

// Add 添加商品分类
func (con GoodsCateController) Add(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := utils.StringToInt(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := utils.StringToInt(c.PostForm("sort"))
	status, err3 := utils.StringToInt(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, ParameterError)
		return
	}
	if err2 != nil {
		con.Error(c, FailedSort)
		return
	}
	cateImgDir, _ := utils.UploadImg(c, "cate_img")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(utils.GetUnix()),
	}
	err := mysql.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(c, FailedAddGoodsCate)
		return
	}
	con.Success(c)
}

// Update 修改商品分类
func (con GoodsCateController) Update(c *gin.Context) {
	id, err1 := utils.StringToInt(c.PostForm("id"))
	title := c.PostForm("title")
	pid, err2 := utils.StringToInt(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err3 := utils.StringToInt(c.PostForm("sort"))
	status, err4 := utils.StringToInt(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, ParameterError)
		return
	}
	if err3 != nil {
		con.Error(c, FailedSort)
		return
	}
	cateImgDir, _ := utils.UploadImg(c, "cate_img")

	goodsCate := models.GoodsCate{Id: id}
	mysql.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}
	err := mysql.DB.Save(&goodsCate).Error
	if err != nil {
		con.Error(c, FailedUpdateGoodsCate)
		return
	}
	con.Success(c)
}

// Delete 删除商品分类
func (con GoodsCateController) Delete(c *gin.Context) {
	id, err := utils.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, ParameterError)
	} else {
		//获取我们要删除的数据
		goodsCate := models.GoodsCate{Id: id}
		mysql.DB.Find(&goodsCate)
		if goodsCate.Pid == 0 { //顶级分类
			var goodsCateList []models.GoodsCate
			// 查询是否有子分类
			mysql.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)
			if len(goodsCateList) > 0 {
				con.Error(c, FailedDeleteChildrenGoodsCateAtFirst)
			} else {
				mysql.DB.Delete(&goodsCate)
				con.Success(c)
			}
		} else { //操作 或者菜单
			mysql.DB.Delete(&goodsCate)
			con.Success(c)
		}

	}
	con.Success(c)
}
