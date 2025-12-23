package service

import (
	"context"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
)

type IPermissionService interface {
	SGetAllPermission(ctx context.Context, d *dto.PageDto) (*paginator.Page[*model.Permission], error)
}
