package service

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/vo"
)

type IUserService interface {
	SLogin(context context.Context, dto interface{}) (*vo.UserVo, error)
	SPostMobileSmsCode(context context.Context, ticket string, mobile string, scene string) error
}
