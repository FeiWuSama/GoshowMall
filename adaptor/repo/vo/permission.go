package vo

import "workspace-goshow-mall/adaptor/repo/model"

type PermissionVo struct {
	model.Permission
	Children []*PermissionVo `json:"children"`
}

type PermissionVoList []*PermissionVo

type PermissionList []*model.Permission
