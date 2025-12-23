package mapper

import (
	"context"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
)

type PermissionMapper interface {
	GetPermissionPage(ctx context.Context, d *dto.PageDto) (*paginator.Page[*model.Permission], error)
	//GetPermissionPageByRoleId(ctx context.Context, roleId int64) ([]*model.Permission, error)
}
