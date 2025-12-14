package lark

import (
	"context"
	"encoding/json"
	"github.com/cnchef/gconv"
	"go.uber.org/zap"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/config"
	"workspace-goshow-mall/rpc"
	"workspace-goshow-mall/service"
	"workspace-goshow-mall/utils/logger"
)

type Service struct {
	config      config.Config
	locker      *redis.Locker
	accessToken *redis.AccessToken
	larkRpc     *rpc.LarkRpc
}

func (s Service) SLarkGetTenantToken(ctx context.Context, appCode int32) (*vo.LarkTenantTokenVo, error) {
	tokenVo, err := s.larkRpc.GetLarkTenantToken(ctx, appCode)
	if err != nil {
		logger.Error("getLarkTenantToken err", zap.Error(err))
		return nil, err
	}
	return tokenVo, nil
}

func (s Service) SLarkGetToken(ctx context.Context, appCode int32, code string, redirectUrl string, scope string) (*vo.LarkAccessTokenVo, error) {
	tokenVo, err := s.getAccessToken(ctx, appCode, code, redirectUrl, scope)
	if err != nil {
		logger.Error("get lark user access token err", zap.Error(err))
		return nil, err
	}
	return tokenVo, nil
}

func (s Service) SLarkUpdateToken(ctx context.Context, tokenFunc service.GetTokenFunc, lockKey string, cacheKey string) (*vo.LarkAccessTokenVo, error) {
	locker, err := s.locker.GetLocker(ctx, lockKey)
	if err != nil {
		logger.Error("get locker err", zap.Error(err))
		return nil, err
	}
	if locker {
		token, err := tokenFunc()
		if err != nil {
			logger.Error("get token err", zap.Error(err))
			return nil, err
		}
		_, err = s.accessToken.SetAccessToken(ctx, cacheKey, gconv.ToString(token), token.ExpiresIn)
		if err != nil {
			logger.Error("set access token err", zap.Error(err))
			return nil, err
		}
		return token, nil
	}
	err = s.locker.AwaitLock(ctx, lockKey)
	if err != nil {
		return nil, err
	}
	return s.getCache(ctx, cacheKey)
}

func (s Service) getCache(ctx context.Context, key string) (*vo.LarkAccessTokenVo, error) {
	accessToken, expire, err := s.accessToken.GetAccessToken(ctx, key)
	if err != nil {
		logger.Error("get access token err", zap.Error(err))
		return nil, err
	}
	accessTokenVo := &vo.LarkAccessTokenVo{}
	err = json.Unmarshal([]byte(accessToken), accessTokenVo)
	if err != nil {
		logger.Error("parse access token err", zap.Error(err))
		return nil, err
	}
	accessTokenVo.ExpiresIn = expire
	return accessTokenVo, nil
}

func (s Service) getAccessToken(ctx context.Context, appCode int32, code string, redirectUrl string, scope string) (*vo.LarkAccessTokenVo, error) {
	getTokenFunc := func() (*vo.LarkAccessTokenVo, error) {
		tokenVo, err := s.larkRpc.GetLarkUserAccessToken(ctx, appCode, code, redirectUrl, scope)
		if err != nil {
			logger.Error("get lark user access token err", zap.Error(err))
			return nil, err
		}
		return tokenVo, nil
	}
	tokenVo, err := getTokenFunc()
	if err != nil {
		logger.Error("rpc err", zap.Error(err))
		return nil, err
	}
	return tokenVo, nil
}
