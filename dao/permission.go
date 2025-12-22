package dao

import (
	"context"
	"gorm.io/gorm"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/query"
)

type PermissionDao struct {
	db *gorm.DB
}

func NewPermissionDao(adaptor adaptor.Adaptor) *PermissionDao {
	return &PermissionDao{
		db: adaptor.Db,
	}
}

func (p PermissionDao) GetPermissionPage(ctx context.Context, d *dto.PageDto) ([]*model.Permission, int64, error) {
	qs := query.Use(p.db).Permission
	do := qs.WithContext(ctx)
	if d.Name != "" {
		do.Where(qs.Name.Like("%" + d.Name + "%"))
	}
	if d.Sort != 0 {
		if d.Sort == 1 {
			do.WithContext(ctx).Order(qs.Sort.Asc())
		} else if d.Sort == -1 {
			do.WithContext(ctx).Order(qs.Sort.Desc())
		}
	}
	if d.CreateBy != 0 {
		do.Where(qs.CreateBy.Eq(d.CreateBy))
	}
	if d.UpdateBy != 0 {
		do.Where(qs.UpdateBy.Eq(d.UpdateBy))
	}
	page, total, err := do.FindByPage(d.PageNum-1, d.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return page, total, err
}
