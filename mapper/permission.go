package mapper

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
)

type PermissionMapper interface {
	GetPermissionPage(ctx context.Context, d *dto.PageDto) ([]*model.Permission, int64, error)
}
