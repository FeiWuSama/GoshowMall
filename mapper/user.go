package mapper

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/model"
)

type UserMapper interface {
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
	GetUserByOpenIdAndCode(ctx context.Context, id string, code int32) (*model.User, error)
}
