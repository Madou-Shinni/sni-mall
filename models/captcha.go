package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
	"xiaomi-mall/models/redis"
)

// 设置自带的 store
//var store = base64Captcha.DefaultMemStore
// 设置自定义store 让RedisStore实现Store接口
var store base64Captcha.Store = redis.RedisStore{}

// CaptMake 生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	// 配置验证码信息
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125},
		Fonts: []string{"wqy-microhei.ttc"}}

	//ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	//创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	//Generate 生成随机 id、base64 图像字符串
	id, b64s, err = captcha.Generate()
	return id, b64s, err
}

// CaptVerify 验证 captcha 是否正确
func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, true) {
		return true
	} else {
		return false
	}
}