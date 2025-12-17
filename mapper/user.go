package mapper

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/model"
)

type UserMapper interface {
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
	GetUserByOpenIdAndCode(ctx context.Context, id string, code int32) (*model.User, error)
	GetUserByNickName(ctx context.Context, name string) (*model.User, error)
	AddUser(ctx context.Context, newUser *model.User) (int64, error)
	AddMobileUser(ctx context.Context, id int64, mobile string) error
}
