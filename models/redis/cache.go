package redis

import (
	"encoding/json"
	"time"
)

var CacheDB = &cacheDB{} // 实例化结构体

type cacheDB struct {
}

// Set 缓存 1.key 2.value 3.过期时间（秒）
func (c cacheDB) Set(key string, value interface{}, expiration int) {
	v, err := json.Marshal(value) // 序列化可以让我们保存结构体等其他数据
	if err == nil {
		rdb.Set(key, string(v), time.Second*time.Duration(expiration))
	}
}

// Get 缓存 1.key 2.value 3.过期时间
func (c cacheDB) Get(key string, obj interface{}) bool {
	result, err := rdb.Get(key).Result()
	if err != nil || result == "" {
		return false
	}

	err2 := json.Unmarshal([]byte(result), obj) // 反序列化可以让我们 获取到数据
	if err2 != nil {
		return false
	}

	return true
}
