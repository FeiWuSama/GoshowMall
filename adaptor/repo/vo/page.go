package vo

import (
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor/repo/model"
)

// PermissionPage 专门给 swag 用，避免写泛型语法
type PermissionPage = paginator.Page[*model.Permission]
