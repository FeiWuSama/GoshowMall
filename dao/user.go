package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/query"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/utils/sha256"
)

func NewUserDao(adaptor adaptor.Adaptor) *UserDao {
	return &UserDao{
		db:          adaptor.Db,
		redisClient: adaptor.Redis,
	}
}

type UserDao struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func (u UserDao) GetUserByOpenIdAndCode(ctx context.Context, id string, code int32) (*model.User, error) {
	qs1 := query.Use(u.db).AppUser
	appUser, err := qs1.WithContext(ctx).Where(qs1.OpenID.Eq(id), qs1.AppCode.Eq(code)).First()
	if err != nil {
		return nil, err
	}
	if appUser == nil {
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "用户不存在")
	}
	qs2 := query.Use(u.db).User
	user, err := qs2.WithContext(ctx).Where(qs2.ID.Eq(appUser.UserID)).First()
	return user, nil
}

func (u UserDao) GetUserByMobile(ctx context.Context, loginDto *dto.UserMobilePasswordLoginDto) (*model.User, error) {
	mobileSha256 := sha256.NewSHA256Crypto().HashToBase64(loginDto.Mobile)
	qs1 := query.Use(u.db).MobileUser
	mobileUser, err := qs1.WithContext(ctx).Where(qs1.MobileSha256.Eq(mobileSha256)).First()
	if err != nil {
		return nil, err
	}
	if mobileUser == nil {
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "手机号未注册")
	}
	qs2 := query.Use(u.db).User
	user, err := qs2.WithContext(ctx).Where(qs2.ID.Eq(mobileUser.UserID)).First()
	return user, nil
}
