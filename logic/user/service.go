package user

import (
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/dao"
	"workspace-goshow-mall/logic/lark"
	"workspace-goshow-mall/rpc"
)

func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		adapter:     adaptor,
		userMapper:  dao.NewUserDao(*adaptor),
		verify:      redis.NewVerify(adaptor.Redis),
		larkService: lark.NewLarkService(adaptor),
		larkRpc:     rpc.NewLarkRpc(adaptor),
	}
}
