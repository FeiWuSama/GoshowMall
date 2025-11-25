package admin

import (
	"context"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/mapper"
)

type Service struct {
	adapter     *adaptor.Adaptor
	adminMapper mapper.AdminMapper
}

func (s *Service) SChangeStatus(ctx context.Context, id string, status string, changeUserId int64) bool {
	return s.adminMapper.MChangeStatus(ctx, id, status, changeUserId)
}

func (s *Service) SUpdateAdmin(ctx context.Context, adminDto dto.UpdateAdminDto, createUserId int64) bool {
	return s.adminMapper.MUpdateAdmin(ctx, adminDto, createUserId)
}

func (s *Service) SCreateAdmin(ctx context.Context, dto dto.AddAdminDto, createUserId int64) int64 {
	return s.adminMapper.CreateAdmin(ctx, dto, createUserId)
}
