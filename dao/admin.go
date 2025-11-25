package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/query"
	"workspace-goshow-mall/constants"
)

func NewAdminDao(adaptor adaptor.Adaptor) *AdminDao {
	return &AdminDao{
		db:          adaptor.Db,
		redisClient: adaptor.Redis,
	}
}

type AdminDao struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func (a *AdminDao) MChangeStatus(ctx context.Context, id string, status string, changeUserId int64) bool {
	qs := query.Use(a.db).Admin
	userId, err1 := strconv.Atoi(id)
	if err1 != nil {
		return false
	}
	userStatus, err2 := strconv.Atoi(status)
	if err2 != nil {
		return false
	}
	result, err := qs.WithContext(ctx).
		Where(qs.ID.Eq(int64(userId))).
		Updates(model.Admin{
			Status:   int32(userStatus),
			UpdateBy: changeUserId,
			UpdateAt: time.Now(),
		})
	if err != nil || result.Error != nil {
		return false
	}
	return result.RowsAffected > 0
}

func (a *AdminDao) MUpdateAdmin(ctx context.Context, adminDto dto.UpdateAdminDto, createUserId int64) bool {
	qs := query.Use(a.db).Admin
	result, err := qs.WithContext(ctx).
		Where(qs.ID.Eq(int64(adminDto.Id))).
		Updates(model.Admin{
			Name:     adminDto.Name,
			NickName: adminDto.NickName,
			Mobile:   adminDto.Mobile,
			Sex:      adminDto.Sex,
		})
	if err != nil || result.Error != nil {
		return false
	}
	return result.RowsAffected > 0
}

func (a *AdminDao) CreateAdmin(ctx context.Context, dto dto.AddAdminDto, createUseId int64) int64 {
	qs := query.Use(a.db).Admin
	err := qs.WithContext(ctx).Create(&model.Admin{
		Name:     dto.Name,
		NickName: dto.NickName,
		Mobile:   dto.Mobile,
		Sex:      dto.Sex,
		Status:   constants.UserActiveStatus,
		CreateBy: createUseId,
		CreateAt: time.Now(),
		UpdateBy: createUseId,
		UpdateAt: time.Now(),
	})
	if err != nil {
		return 0
	}
	return 1
}
