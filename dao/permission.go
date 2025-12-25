package dao

import (
	"context"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"gorm.io/gorm"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/query"
)

func NewPermissionDao(adaptor adaptor.Adaptor) *PermissionDao {
	return &PermissionDao{
		db: adaptor.Db,
	}
}

type PermissionDao struct {
	db *gorm.DB
}

func (p PermissionDao) GetPermissionByParentId(ctx context.Context, id int64) ([]*model.Permission, error) {
	qs := query.Use(p.db).Permission
	return qs.WithContext(ctx).Where(qs.ParentID.Eq(id)).Find()
}

func (p PermissionDao) GetPermissionByAdminId(ctx context.Context, id int64) ([]*model.Permission, error) {
	qs1 := query.Use(p.db).Permission
	qs2 := query.Use(p.db).RolePermission
	qs3 := query.Use(p.db).AdminRole
	return qs1.WithContext(ctx).
		LeftJoin(qs2, qs2.PermissionID.EqCol(qs1.ID)).
		LeftJoin(qs3, qs3.RoleID.EqCol(qs2.RoleID)).
		Where(qs3.AdminID.Eq(id), qs1.Status.Eq(1)).
		Find()
}

func (p PermissionDao) GetPermissionPageByRoleId(ctx context.Context, roleId int64) ([]*model.Permission, error) {
	qs1 := query.Use(p.db).Permission
	qs2 := query.Use(p.db).RolePermission
	return qs1.WithContext(ctx).LeftJoin(qs2, qs2.PermissionID.EqCol(qs1.ID)).Where(qs2.RoleID.Eq(roleId), qs1.Status.Eq(1)).Find()
}

func (p PermissionDao) GetPermissionPage(ctx context.Context, d *dto.PageDto) (*paginator.Page[*model.Permission], error) {
	qs := query.Use(p.db).Permission
	do := qs.WithContext(ctx)
	do.Where(qs.Status.Eq(1))
	if d.Name != "" {
		do = do.Where(qs.Name.Like("%" + d.Name + "%"))
	}
	if d.Sort != 0 {
		if d.Sort == 1 {
			do = do.Order(qs.Sort.Asc())
		} else if d.Sort == -1 {
			do = do.Order(qs.Sort.Desc())
		}
	}
	if d.CreateBy != 0 {
		do = do.Where(qs.CreateBy.Eq(d.CreateBy))
	}
	if d.UpdateBy != 0 {
		do = do.Where(qs.UpdateBy.Eq(d.UpdateBy))
	}
	page := paginator.Page[*model.Permission]{CurrentPage: d.PageNum, PageSize: d.PageSize}
	err := page.SelectPages(do.UnderlyingDB())
	if err != nil {
		return nil, err
	}
	return &page, nil
}
