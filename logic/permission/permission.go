package permission

import (
	"context"
	"github.com/jinzhu/copier"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/mapper"
)

type Service struct {
	permissionMapper mapper.PermissionMapper
}

func (s Service) SGetPermissionByAdminId(ctx context.Context, userId int64) ([]*model.Permission, error) {
	return s.permissionMapper.GetPermissionByAdminId(ctx, userId)
}

func (s Service) SGetPermissionByRoleId(ctx context.Context, roleId int64) ([]*model.Permission, error) {
	return s.permissionMapper.GetPermissionPageByRoleId(ctx, roleId)
}

func (s Service) SGetAllPermission(ctx context.Context, d *dto.PageDto) (*paginator.Page[*model.Permission], error) {
	page, err := s.permissionMapper.GetPermissionPage(ctx, d)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s Service) ConvertPermissionList2Tree(ctx context.Context, id int64) ([]*vo.PermissionVo, error) {
	permissions, err := s.permissionMapper.GetPermissionByParentId(ctx, id)
	if err != nil {
		return nil, err
	}
	permissionVos := make([]*vo.PermissionVo, len(permissions))
	err = copier.Copy(&permissionVos, permissions)
	if err != nil {
		return nil, err
	}
	for _, permission := range permissionVos {
		children, err := s.ConvertPermissionList2Tree(ctx, permission.ID)
		if err != nil {
			return nil, err
		}
		permission.Children = children
	}
	return permissionVos, nil
}
