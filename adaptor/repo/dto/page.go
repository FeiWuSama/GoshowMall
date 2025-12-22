package dto

type PageDto struct {
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	CreateBy int64  `json:"create_by"`
	UpdateBy int64  `json:"update_by"`
}
