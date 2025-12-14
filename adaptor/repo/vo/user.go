package vo

type UserVo struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Sex      int32  `json:"sex"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type LarkUserVo struct {
	Name            string `json:"name"`
	EnName          string `json:"en_name"`
	AvatarUrl       string `json:"avatar_url"`
	AvatarThumb     string `json:"avatar_thumb"`
	AvatarMiddle    string `json:"avatar_middle"`
	AvatarBig       string `json:"avatar_big"`
	OpenId          string `json:"open_id"`
	UnionId         string `json:"union_id"`
	Email           string `json:"email"`
	EnterpriseEmail string `json:"enterprise_email"`
	UserId          string `json:"user_id"`
	Mobile          string `json:"mobile"`
	TenantKey       string `json:"tenant_key"`
	EmployeeNo      string `json:"employee_no"`
}

type LarkAccessTokenVo struct {
	Code        int64  `json:"code"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int64  `json:"err_code"`
	ErrMsg      string `json:"err_msg"`
}

type LarkTenantTokenVo struct {
	Code              int64  `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}
