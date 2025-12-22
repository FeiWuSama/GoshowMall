package permission

import (
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/dao"
)

func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		permissionMapper: dao.NewPermissionDao(*adaptor),
	}
}
