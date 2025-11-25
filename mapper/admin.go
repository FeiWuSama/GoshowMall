package mapper

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/dto"
)

type AdminMapper interface {
	CreateAdmin(ctx context.Context, dto dto.AddAdminDto, userId int64) int64
	MUpdateAdmin(ctx context.Context, adminDto dto.UpdateAdminDto, id int64) bool
	MChangeStatus(ctx context.Context, id string, status string, changeUserId int64) bool
}
