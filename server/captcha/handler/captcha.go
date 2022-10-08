package handler

import (
	pb "captcha/proto/captcha"
	"context"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

//创建store
var store = base64Captcha.DefaultMemStore

type Captcha struct{}

// GetCaptcha 获取验证码
func (e *Captcha) GetCaptcha(ctx context.Context, req *pb.GetCaptchaRequest, rsp *pb.GetCaptchaResponse) error {
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
	id, b64s, err := captcha.Generate()
	// 返回数据
	rsp.Id = id
	rsp.B64S = b64s

	return err
}

// VerifyCaptcha 验证验证码
func (e *Captcha) VerifyCaptcha(ctx context.Context, req *pb.CaptVerifyRequest, rsp *pb.CaptVerifyResponse) error {
	if store.Verify(req.Id, req.Capt, true) {
		rsp.Flag = true
	} else {
		rsp.Flag = false
	}
	return nil
}
