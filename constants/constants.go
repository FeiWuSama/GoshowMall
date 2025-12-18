package constants

const (
	UserToken  = "token"
	AdminToken = "admin-token"
)

const (
	UserTokenKey     = "goshow:user:token:"
	AdminTokenKey    = "goshow:admin:token:"
	SlideCaptchaKey  = "goshow:slide:captcha:"
	CaptchaTicketKey = "goshow:ticket:"
	PasswordErrorKey = "goshow:error:password:"
	SmsCodeKey       = "goshow:sms:"

	TokenExpire    = 60 * 60 * 24 * 3
	CaptchaExpire  = 60 * 5
	SmsLoginExpire = 60 * 5

	CaptchaTicketExpire = 60
	PasswordErrorExpire = 10
	LockerExpire        = 5
)

const (
	UserActiveStatus = 1
	UserBanStatus    = -1
)

const PasswordErrorCount = 5

const (
	WechatAppCode = 1001
	LarkAppCode   = 1002
)

const (
	OpenIdType  = "open_id"
	UnionIdType = "union_id"
	UserIdType  = "user_id"
	ChatIdType  = "chat_id"
)
