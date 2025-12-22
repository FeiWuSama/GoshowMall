package vo

import "workspace-goshow-mall/adaptor/repo/model"

// PermissionPage 专门给 swag 用，避免写泛型语法
type PermissionPage = PageVo[*model.Permission]

type PageVo[T any] struct {
	Total  int64 `json:"total"`
	Record []T   `json:"record"`
}
