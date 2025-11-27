package router

var whiteList = map[string]bool{
	"/health":        true,
	"/check":         true,
	"/admin/captcha": true,
}
