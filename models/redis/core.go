package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"os"
)

// 声明一个rdb全局变量
var rdb *redis.Client

// Init 初始化连接
func init() {

	//读取.ini里面的数据库配置
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	password := config.Section("redis").Key("password").String()
	db, _ := config.Section("redis").Key("db").Int()

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", ip, port),
		Password: password,
		DB:       db,
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
