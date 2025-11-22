package admin

import (
	"context"
	"workspace-goshow-mall/adaptor"
)

type Service struct {
	adapter *adaptor.Adaptor
}

func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		adapter: adaptor,
	}
}

func (s *Service) HelloWorld(c context.Context) string {
	return "Hello World"
}
