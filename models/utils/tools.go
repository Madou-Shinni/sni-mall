package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
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

// IntToString 将string转换成int
func IntToString(int int) string {
	str := strconv.Itoa(int)
	return str
}
