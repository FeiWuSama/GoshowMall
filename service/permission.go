package service

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/vo"
)

type IPermissionService interface {
	SGetAllPermission(ctx context.Context, d *dto.PageDto) (*vo.PageVo[*model.Permission], error)
}
