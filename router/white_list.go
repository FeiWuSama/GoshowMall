package router

var whiteList = map[string]bool{
	"/health":                     true,
	"/check":                      true,
	"/admin/captcha/slide":        true,
	"/user/info":                  true,
	"/user/captcha/slide":         true,
	"/user/captcha/slide/verify":  true,
	"/user/mobile/login/password": true,
	"/user/lark/login":            true,
	"/user/mobile/smsCode":        true,
	"/admin/login":                true,
	"/admin/captcha/slide/verify": true,
	"/admin/info":                 true,
}
