package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Md5 加密
func Md5(str string) string {
	h := md5.New()         // 创建md5实例
	io.WriteString(h, str) // 传入实例和加密参数
	return fmt.Sprintf("%x", h.Sum(nil))
}

// StringToInt 将string转换成int
func StringToInt(str string) (int, error) {
	int, err := strconv.Atoi(str)
	return int, err
}

// IntToString 将int转换成string
func IntToString(int int) string {
	str := strconv.Itoa(int)
	return str
}

// StringToInt64 将string转换成int64
func StringToInt64(str string) (int64, error) {
	int, err := strconv.ParseInt(str, 10, 64)
	return int, err
}

// Int64ToString 将int64转换成string
func Int64ToString(int64 int64) string {
	str := strconv.FormatInt(int64, 10)
	return str
}

// GetUnix 获取当前时间
func GetUnix() int64 {
	return time.Now().Unix()
}
