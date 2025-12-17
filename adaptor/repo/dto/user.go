package dto

type UserMobileLoginDto struct {
	Mobile string `json:"mobile"`
}

type UserRegisterDto struct {
	Mobile     string `json:"mobile"`
	NickName   string `json:"nickname"`
	Password   string `json:"password"`
	Sex        int32  `json:"sex"`
	VerifyCode string `json:"verify_code"`
	Scene      string `json:"scene"`
}

type UserMobilePasswordLoginDto struct {
	UserMobileLoginDto
	Ticket   string `json:"ticket"`
	Password string `json:"password"`
}

type UserLarkLoginDto struct {
	AppCode     int32  `json:"app_code"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

type UserLarkMsgDto struct {
	AppCode int32  `json:"app_code"`
	OpenId  string `json:"open_id"`
	IdType  string
	Content string `json:"content"`
}

type UserMobileSmsLoginDto struct {
	UserMobileLoginDto
	VerifyCode string `json:"verify_code"`
	Scene      string `json:"scene"`
}
