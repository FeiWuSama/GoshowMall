package vo

type UserVo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type CaptchaVo struct {
	Key              string `json:"key"`
	ImageBase64      string `json:"ImageBase64"`
	TitleImageBase64 string `json:"TitleImageBase64"`
	TitleHeight      int    `json:"TitleHeight"`
	TitleWidth       int    `json:"TitleWidth"`
	TitleX           int    `json:"TitleX"`
	TitleY           int    `json:"TitleY"`
}
