package itying

import (
	"net/http"
	"strings"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
	"xiaomi-mall/models/redis"
	"xiaomi-mall/models/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DefaultController struct{}

// Index 首页数据渲染
func (con DefaultController) Index(c *gin.Context) {
	//1、获取顶部导航
	var topNavList []models.Nav
	if b := redis.CacheDB.Get("topNavList", &topNavList); !b { // 先获取缓存数据
		// redis没获取到数据，mysql获取
		mysql.DB.Where("status=1 AND position=1").Find(&topNavList)
		// 缓存 5~10分钟
		expir := utils.Rand(6) + 5
		redis.CacheDB.Set("topNavList", topNavList, expir*60)
	}

	//2、获取轮播图数据
	var focusList []models.Focus
	if b := redis.CacheDB.Get("focusList", &focusList); !b { // 先获取缓存数据
		// redis没获取到数据，mysql获取
		mysql.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		// 缓存 5~10分钟
		expir := utils.Rand(6) + 5
		redis.CacheDB.Set("focusList", focusList, expir*60)
	}

	//3、获取分类的数据
	var goodsCateList []models.GoodsCate
	//https://gorm.io/zh_CN/docs/preload.html
	if b := redis.CacheDB.Get("goodsCateList", &goodsCateList); !b { // 先获取缓存数据
		// redis没获取到数据，mysql获取
		mysql.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)
		// 缓存 5~10分钟
		expir := utils.Rand(6) + 5
		redis.CacheDB.Set("goodsCateList", goodsCateList, expir*60)
	}

	//4、获取中间导航
	var middleNavList []models.Nav
	if b := redis.CacheDB.Get("middleNavList", &middleNavList); !b { // 先获取缓存数据
		// redis没获取到数据，mysql获取
		mysql.DB.Where("status=1 AND position=2").Find(&middleNavList)
		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
			relationIds := strings.Split(relation, ",")
			var goodsList []models.Goods
			mysql.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		// 缓存 5~10分钟
		expir := utils.Rand(6) + 5
		redis.CacheDB.Set("middleNavList", middleNavList, expir*60)
	}

	//手机
	var phoneList []models.Goods
	if b := redis.CacheDB.Get("phoneList", &phoneList); !b { // 先获取缓存数据
		phoneList = models.GetGoodsByCategory(1, "best", 8)
		// 缓存 5~10分钟
		expir := utils.Rand(6) + 5
		redis.CacheDB.Set("phoneList", phoneList, expir*60)
	}

	//配件
	var otherList []models.Goods
	if b := redis.CacheDB.Get("otherList", &otherList); !b { // 先获取缓存数据
		otherList = models.GetGoodsByCategory(9, "all", 1)
		// 缓存 5~10分钟
		expir := utils.Rand(6) + 5
		redis.CacheDB.Set("otherList", otherList, expir*60)
	}

	c.JSON(http.StatusOK, gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		"otherList":     otherList,
	})

}
