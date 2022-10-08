package models

import (
	"context"
	"go-micro.dev/v4/logger"
	pb "xiaomi-mall/proto/captcha"
)

var (
	service = "captcha"
	version = "latest"
)

// 设置自带的 store
//var store = base64Captcha.DefaultMemStore
// 设置自定义store 让RedisStore实现Store接口
//var store base64Captcha.Store = redis.RedisStore{}

// CaptMake 生成验证码
func CaptMake() (id, b64s string, err error) {

	// Create client
	captchaClient := pb.NewCaptchaService(service, CaptchaClient)

	// Call service（调用微服务）
	rsp, err := captchaClient.GetCaptcha(context.Background(), &pb.GetCaptchaRequest{})
	if err != nil {
		logger.Fatal(err)
	}

	return rsp.Id, rsp.B64S, err
}

// CaptVerify 验证 captcha 是否正确
func CaptVerify(id string, capt string) bool {
	// Create client
	captchaClient := pb.NewCaptchaService(service, CaptchaClient)

	// Call service
	rsp, err := captchaClient.VerifyCaptcha(context.Background(), &pb.CaptVerifyRequest{})
	if err != nil {
		logger.Fatal(err)
	}

	return rsp.Flag
}
