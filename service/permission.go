package service

import (
	"context"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/vo"
)

type IPermissionService interface {
	SGetAllPermission(ctx context.Context, d *dto.PageDto) (*paginator.Page[*model.Permission], error)
	SGetPermissionByRoleId(ctx context.Context, roleId int64) ([]*model.Permission, error)
	SGetPermissionByAdminId(ctx context.Context, userId int64) ([]*model.Permission, error)
	ConvertPermissionList2Tree(ctx context.Context, id int64) ([]*vo.PermissionVo, error)
}
