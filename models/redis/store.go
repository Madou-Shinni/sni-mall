package redis

import (
	"fmt"
	"time"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// Set 存放验证码
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := rdb.Set(key, value, time.Minute*5).Err()
	return err
}

// Get 获取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := rdb.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := rdb.Del(key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 验证验证码
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
