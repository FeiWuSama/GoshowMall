package lark

import (
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/rpc"
)

func NewLarkService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		config:      adaptor.Config,
		accessToken: redis.NewAccessToken(adaptor),
		locker:      redis.NewLocker(adaptor),
		larkRpc:     rpc.NewLarkRpc(adaptor),
	}
}
