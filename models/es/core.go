package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client
var err error

func init() {
	//注意IP和端口
	EsClient, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

	if err != nil {
		fmt.Println(err)
	}
}
