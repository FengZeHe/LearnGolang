package sms

import "fmt"

/*
生成验证码
*/
func GenCode() (code string) {
	return "123456"
}

/*
发送验证码
*/
func SendSMS(phone, code string) (err error) {
	fmt.Println("phone", phone, "code", code, "发送验证码成功")
	return nil
}

/*
验证验证码
*/
func VerifyCode(phone, code string) (err error) {
	return nil
}
