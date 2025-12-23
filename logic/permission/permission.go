package permission

import (
	"context"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/mapper"
)

type Service struct {
	permissionMapper mapper.PermissionMapper
}

func (s Service) SGetAllPermission(ctx context.Context, d *dto.PageDto) (*paginator.Page[*model.Permission], error) {
	page, err := s.permissionMapper.GetPermissionPage(ctx, d)
	if err != nil {
		return nil, err
	}
	return page, nil
}
