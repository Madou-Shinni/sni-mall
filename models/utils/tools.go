package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(str string) string {
	h := md5.New()         // 创建md5实例
	io.WriteString(h, str) // 传入实例和加密参数
	return fmt.Sprintf("%x", h.Sum(nil))
}
