package service

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/vo"
)

type GetTokenFunc func() (*vo.LarkAccessTokenVo, error)

type ILarkService interface {
	SLarkUpdateToken(ctx context.Context, tokenFunc GetTokenFunc, lockKey string, cacheKey string) (*vo.LarkAccessTokenVo, error)
	SLarkGetToken(ctx context.Context, appCode int32, code string, redirectUrl string, scope string) (*vo.LarkAccessTokenVo, error)
	SLarkGetTenantToken(ctx context.Context, appCode int32) (*vo.LarkTenantTokenVo, error)
}
