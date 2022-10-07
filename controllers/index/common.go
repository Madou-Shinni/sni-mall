package index

import (
	"net/http"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/redis"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct{}

// Render 封装首页和商品页都有的数据
func (con BaseController) Render(c *gin.Context, data map[string]interface{}) {

	//1、获取顶部导航
	var topNavList []models.Nav
	if hasTopNavList := redis.CacheDB.Get("topNavList", &topNavList); !hasTopNavList {
		mysql.DB.Where("status=1 AND position=1").Find(&topNavList)
		redis.CacheDB.Set("topNavList", topNavList, 60*60)
	}

	//2、获取分类的数据
	var goodsCateList []models.GoodsCate

	if hasGoodsCateList := redis.CacheDB.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		//https://gorm.io/zh_CN/docs/preload.html
		mysql.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		redis.CacheDB.Set("goodsCateList", goodsCateList, 60*60)
	}

	//3、获取中间导航
	var middleNavList []models.Nav
	if hasMiddleNavList := redis.CacheDB.Get("middleNavList", &middleNavList); !hasMiddleNavList {
		mysql.DB.Where("status=1 AND position=2").Find(&middleNavList)
		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
			relationIds := strings.Split(relation, ",")
			var goodsList []models.Goods
			mysql.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		redis.CacheDB.Set("middleNavList", middleNavList, 60*60)
	}

	renderData := gin.H{
		"topNavList":    topNavList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
	}

	for key, v := range data {
		renderData[key] = v
	}

	c.JSON(http.StatusOK, renderData)

}
