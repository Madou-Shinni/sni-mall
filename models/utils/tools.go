package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hunterhug/go_image"
	qrcode "github.com/skip2/go-qrcode"
	"io"
	"io/ioutil"
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

// Float64 将string转换成float64
func Float64(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
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

// QrcodeSava 生成二维码并且保存成文件
func QrcodeSava(c *gin.Context) {
	savePath := "static/upload/qrcode.png"
	err := qrcode.WriteFile("https://www.itying.com", qrcode.Medium, 556, savePath)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	file, _ := ioutil.ReadFile(savePath)
	c.String(200, string(file))
}

// QrcodeByte 生成二维码字节
func QrcodeByte(c *gin.Context) {
	var png []byte
	png, err := qrcode.Encode("https://www.itying.com", qrcode.Medium, 256)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	c.String(200, string(png))
}

// Thumbnail1 图像宽度 进行等比例缩放
func Thumbnail1(c *gin.Context) {
	//按宽度进行比例缩放，输入输出都是文件
	//filename string, savepath string, width int
	filename := "static/upload/0.png"
	savePath := "static/upload/0_600.png"
	err := ScaleF2F(filename, savePath, 600)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail1 成功")
}

// Thumbnail2 图像宽高 都进行等比例缩放
func Thumbnail2(c *gin.Context) {
	filename := "static/upload/tao.jpg"
	savePath := "static/upload/tao_400.png"
	//按宽度和高度进行比例缩放，输入和输出都是文件
	err := ThumbnailF2F(filename, savePath, 400, 400)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail2 成功")
}
