package sms

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
	return nil
}

/*
验证验证码
*/
func VerifyCode(phone, code string) (err error) {
	return nil
}
