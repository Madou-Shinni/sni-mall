package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个rdb全局变量
var rdb *redis.Client

// Init 初始化连接
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 8,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Close() {
	_ = rdb.Close()
}
