package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// DateToUnix 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetUnix 获取时间戳毫秒
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetUnixNano 获取时间戳纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// GetDate 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// GetDay 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

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

// UploadImg 上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err1 := c.FormFile(picName)
	if err1 != nil {
		return "", err1
	}
	// 2、获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{".jpg": true, ".png": true, ".gif": true, ".jpeg": true}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	// 3、创建图片保存的目录 static/upload/20210624
	day := GetDay()
	dir := "./static/upload/" + day
	err := os.MkdirAll(dir, 0666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// 4、生成文件名称和文件保存的目录 111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName
	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}
