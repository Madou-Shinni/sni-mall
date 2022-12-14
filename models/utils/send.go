package utils

import (
	"fmt"
	"xiaomi-mall/models/redis"

	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

// SMS 发送短信验证码
func SMS(phone string, code string) {
	clnt := ypclnt.New("62c24eee15fxxxxxxxxxxxxxxxe0beca9e0a") //apikey https://www.yunpian.com/官网获取
	param := ypclnt.NewParam(2061690)                          //模板id
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TEXT] = "【xiaomi-mall】您的验证码是" + code
	r := clnt.Sms().SingleSend(param)
	fmt.Println(r)
}

// SMSByVoice 发送语音
func SMSByVoice(phone string, code string) {
	clnt := ypclnt.New("62c24eee15fxxxxxxxxxxxxxxxe0beca9e0a") //apikey
	param := ypclnt.NewParam(2061690)                          //模板id
	param[ypclnt.MOBILE] = phone
	param[ypclnt.CODE] = code
	r := clnt.Voice().Send(param)
	fmt.Println(r)
}

// VerifyCode 验证短信验证码 失败返回false成功返回true
func VerifyCode(phone string, codeStr string) bool {
	code, err := redis.CacheDB.GetString(phone)
	if err != nil || code == "" || code != codeStr {
		return false
	}
	return true
}
