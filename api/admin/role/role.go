package role

import (
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/api"
	"workspace-goshow-mall/logic/permission"
	"workspace-goshow-mall/service"
)

type Ctrl struct {
	api.BaseCtrl
	adaptor           *adaptor.Adaptor
	permissionService service.IPermissionService
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor:           adaptor,
		permissionService: permission.NewService(adaptor),
	}
}
