package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	. "github.com/hunterhug/go_image"
	qrcode "github.com/skip2/go-qrcode"
	"gopkg.in/ini.v1"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
	"xiaomi-mall/models"
	mysql "xiaomi-mall/models/mysql"
)

const (
	FailedFileUpload = "文件上传失败！"
)

// Rand 生成随机数 生成0到v-1的随机数
func Rand(v int) int {
	//将时间戳设置成种子数（真正随机）
	rand.Seed(time.Now().UnixNano())
	//过期时间生成0-99之间的随机数 rand.Intn(100)固定随机，需要配合rand.Seed使用
	return rand.Intn(v)
}

// RandCode 随机验证码4位
func RandCode() string {
	var str string
	for i := 0; i < 4; i++ {
		v := IntToString(Rand(10))
		str += v
	}
	return str
}

// GetOrderId 获取订单随机数
func GetOrderId() string {
	template := "20160102150405"
	return time.Now().Format(template) + RandCode()
}

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

// Substr 截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])

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

// LocalUploadImg 上传图片到本地
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
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

// ResizeGoodsImage 生成商品缩略图
func ResizeGoodsImage(filename string) {
	extname := path.Ext(filename)                          // 获取后缀
	thumbnailSize := GetSettingFromColumn("ThumbnailSize") // 获取列名的属性值
	thumbnailSizeSlice := strings.Split(thumbnailSize, ",")
	// static/upload/tao_400.png
	// 生成后 static/upload/tao_400.png_100x100.png
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		w, _ := StringToInt(thumbnailSizeSlice[i])
		// 拼接字符串
		savePath := "static/upload/tao_400.png" + "_" + thumbnailSizeSlice[i] + "x" + thumbnailSizeSlice[i] + extname
		//按宽度和高度进行比例缩放，输入和输出都是文件
		err := ThumbnailF2F(filename, savePath, w, w)
		if err != nil {
			fmt.Println(err) // 日志
		}
	}
}

// GetSettingFromColumn 通过列获取值
func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := models.Setting{}
	mysql.DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

// GetOssStatus 获取Oss的状态
func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := StringToInt(config.Section("oss").Key("status").String())
	return ossStatus
}

// UploadImg 上传图片(自动判断oss还是本地)
func UploadImg(c *gin.Context, picName string) (string, error) {
	ossStatus := GetOssStatus()
	// 查看oss是否开启
	if ossStatus == 1 {
		return OssUploadImg(c, picName)
	} else {
		return LocalUploadImg(c, picName)
	}

}

// OssUploadImg 上传图片到Oss（内部调用了OssUpload）
func OssUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、定义图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "static/upload/" + day

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)

	OssUpload(file, dst)
	return dst, nil

}

// OssUpload Oss上传
func OssUpload(file *multipart.FileHeader, dst string) (string, error) {

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "GJoqWHXB2c9S9gwP", "Lgf3weXuWITUUb17vDJfveg1jmKEe9")
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket("beego")
	if err != nil {
		return "", err
	}

	// 上传文件流。
	err = bucket.PutObject(dst, f)
	if err != nil {
		return "", err
	}
	return dst, nil
}
