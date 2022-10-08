package index

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
	"xiaomi-mall/models"
	"xiaomi-mall/models/es"
	mysql "xiaomi-mall/models/mysql"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	BaseController
}

// Index 初始化的时候判断goods是否存在  创建索引配置映射
func (con SearchController) Index(c *gin.Context) {
	// 查看goods索引是否存在
	exists, err := es.EsClient.IndexExists("goods").Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	print(exists)
	if !exists {
		// 配置映射
		mapping := `
		{
			"settings": {
			  "number_of_shards": 1,
			  "number_of_replicas": 0
			},
			"mappings": {
			  "properties": {
				"Content": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				},
				"Title": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				}
			  }
			}
		  }
		`
		//注意：增加的写法
		_, err := es.EsClient.CreateIndex("goods").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Println(err)
		}
	}

	c.String(200, "创建索引配置映射成功")
}

// AddGoods 增加商品数据
func (con SearchController) AddGoods(c *gin.Context) {
	var goods []models.Goods
	mysql.DB.Find(&goods)

	addResult, err := es.EsClient.Index().
		Index("goods").
		Type("_doc").
		Id(strconv.Itoa(goods[1].Id)).
		BodyJson(goods[1]).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", addResult.Id, addResult.Index, addResult.Type)

	c.String(200, "AddGoods success")
}

// UpdateGoods 更新数据
func (con SearchController) UpdateGoods(c *gin.Context) {

	var goods []models.Goods
	mysql.DB.Find(&goods)
	goods[0].Title = "我是修改后的数据"
	goods[0].GoodsContent = "我是修改后的数据GoodsContent"

	_, err := es.EsClient.Update().
		Index("goods").
		Type("_doc").
		Id("19").
		Doc(goods[0]).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	c.String(200, "修改数据 success")
}

// DeleteGoods 删除
func (con SearchController) DeleteGoods(c *gin.Context) {

	_, err := es.EsClient.Delete().
		Index("goods").
		Type("_doc").
		Id("19").
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	c.String(200, "删除成功 success")
}

// GetOne 查询一条数据
func (con SearchController) GetOne(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "GetOne Error")
		}
	}()
	result, err := es.EsClient.Get().
		Index("goods").
		Type("_doc").
		Id("19").
		Do(context.Background())
	if err != nil {
		// Some other kind of error
		panic(err)
	}
	goods := models.Goods{}
	fmt.Printf("%#v", result.Source)
	json.Unmarshal(result.Source, &goods)

	c.JSON(200, gin.H{
		"goods": goods,
	})
}

// Query 模糊查询
func (con SearchController) Query(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "Query Error")
		}
	}()
	query := elastic.NewMatchQuery("Title", "手机")
	searchResult, err := es.EsClient.Search().
		Index("goods").          // search in index "twitter"
		Query(query).            // specify the query
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	goods := models.Goods{}
	c.JSON(200, gin.H{
		"searchResult": searchResult.Each(reflect.TypeOf(goods)),
	})
}

// FilterQuery 条件筛选查询
func (con SearchController) FilterQuery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "Query Error")
		}
	}()

	//筛选
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("Title", "小米"))
	boolQ.Filter(elastic.NewRangeQuery("Id").Gt(19))
	boolQ.Filter(elastic.NewRangeQuery("Id").Lt(42))
	searchResult, err := es.EsClient.Search().
		Index("goods").
		Type("_doc").
		Sort("Id", true).
		Query(boolQ).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
	}
	var goodsList []models.Goods
	var goods models.Goods
	for _, item := range searchResult.Each(reflect.TypeOf(goods)) {
		t := item.(models.Goods)
		fmt.Printf("Id:%v 标题：%v\n", t.Id, t.Title)
		goodsList = append(goodsList, t)
	}

	c.JSON(200, gin.H{
		"goodsList": goodsList,
	})

}

// PagingQuery 分页查询
func (con SearchController) PagingQuery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "Query Error")
		}
	}()

	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	pageSize := 2
	query := elastic.NewMatchQuery("Title", "手机")
	searchResult, err := es.EsClient.Search().
		Index("goods").                             // search in index "twitter"
		Query(query).                               // specify the query
		Sort("Id", true).                           // true 表示升序   false 降序
		From((page - 1) * pageSize).Size(pageSize). // take documents 0-9
		Do(context.Background())                    // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	goods := models.Goods{}
	c.JSON(200, gin.H{
		"searchResult": searchResult.Each(reflect.TypeOf(goods)),
	})
}
