package dto

type UserMobileLoginDto struct {
	Mobile string `json:"mobile"`
	Ticket string `json:"ticket"`
}

type UserMobilePasswordLoginDto struct {
	UserMobileLoginDto
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
