package permission

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/mapper"
)

type Service struct {
	permissionMapper mapper.PermissionMapper
}

func (s Service) SGetAllPermission(ctx context.Context, d *dto.PageDto) (*vo.PageVo[*model.Permission], error) {
	permissions, total, err := s.permissionMapper.GetPermissionPage(ctx, d)
	if err != nil {
		return nil, err
	}
	return &vo.PageVo[*model.Permission]{
		Record: permissions,
		Total:  total,
	}, nil
}
