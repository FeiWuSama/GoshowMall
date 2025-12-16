package service

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/vo"
)

type IAdminService interface {
	SCreateAdmin(ctx context.Context, dto dto.AddAdminDto, createUserId int64) int64
	SUpdateAdmin(ctx context.Context, adminDto dto.UpdateAdminDto, createUserId int64) bool
	SChangeStatus(ctx context.Context, id string, status string, changeUserId int64) bool
	SLogin(ctx context.Context, loginDto *dto.AdminLoginDto, ticket string) (*vo.AdminVO, error)
}
