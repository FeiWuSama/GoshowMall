package vo

type AdminVO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Sex      int32  `json:"sex"`
	Token    string `json:"token"`
}
