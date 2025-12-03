package mapper

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
)

type UserMapper interface {
	GetUserByMobile(ctx context.Context, loginDto *dto.UserMobilePasswordLoginDto) (*model.User, error)
}
