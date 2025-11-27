package constants

const (
	UserToken  = "token"
	AdminToken = "token"
)

const (
	UserTokenKey     = "goshow:user:token:"
	AdminTokenKey    = "goshow:admin:token:"
	CaptchaKey       = "goshow:captcha:"
	CaptchaTicketKey = "goshow:ticket:"

	TokenExpire   = 60 * 60 * 24 * 3
	CaptchaExpire = 60 * 5
)

const (
	UserActiveStatus = 1
	UserBanStatus    = -1
)
