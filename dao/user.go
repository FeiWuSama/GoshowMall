package dao

import (
	"context"
	"github.com/cnchef/gconv"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/query"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/utils/aes"
	"workspace-goshow-mall/utils/random"
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

func (u UserDao) AddMobileUser(ctx context.Context, id int64, mobile string) error {
	qs := query.Use(u.db).MobileUser
	encryptAES, err := aes.EncryptAES([]byte(mobile), []byte("feiwusama"))
	if err != nil {
		return err
	}
	return qs.WithContext(ctx).Create(&model.MobileUser{
		MobileAes:    string(encryptAES),
		MobileSha256: sha256.NewSHA256Crypto().HashToBase64(mobile),
		UserID:       id,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
	})
}

func (u UserDao) AddUser(ctx context.Context, newUser *model.User) (int64, error) {
	qs := query.Use(u.db).User
	if &newUser.CreateBy == nil {
		newUser.CreateBy = 0
	}
	id := random.GenUserUUId()
	insertUser := &model.User{
		ID:       int64(gconv.ToInt(id)),
		NickName: newUser.NickName,
		Password: newUser.Password,
		Sex:      newUser.Sex,
		Status:   1,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		CreateBy: newUser.CreateBy,
	}
	err := qs.WithContext(ctx).Create(insertUser)
	if err != nil {
		return 0, err
	}
	return insertUser.ID, nil
}

func (u UserDao) GetUserByNickName(ctx context.Context, name string) (*model.User, error) {
	qs := query.Use(u.db).User
	return qs.WithContext(ctx).Where(qs.NickName.Eq(name)).First()
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

func (u UserDao) GetUserByMobile(ctx context.Context, mobile string) (*model.User, error) {
	mobileSha256 := sha256.NewSHA256Crypto().HashToBase64(mobile)
	//qs1 := query.Use(u.db).MobileUser
	//mobileUser, err := qs1.WithContext(ctx).Where(qs1.MobileSha256.Eq(mobileSha256)).First()
	//if err != nil {
	//	return nil, err
	//}
	//if mobileUser == nil {
	//	return nil, result.NewBusinessErrorWithMsg(result.ParamError, "手机号未注册")
	//}
	//qs2 := query.Use(u.db).User
	//user, err := qs2.WithContext(ctx).Where(qs2.ID.Eq(mobileUser.UserID)).First()
	qs1 := query.Use(u.db).User
	qs2 := query.Use(u.db).MobileUser
	return qs1.WithContext(ctx).
		LeftJoin(qs2, qs1.ID.EqCol(qs2.UserID)).
		Where(qs1.Status.Eq(constants.UserActiveStatus), qs2.MobileSha256.Eq(mobileSha256)).First()
}
