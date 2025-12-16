package admin

import (
	"context"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/mapper"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/utils/md5"
	"workspace-goshow-mall/utils/random"
)

type Service struct {
	adapter     *adaptor.Adaptor
	adminMapper mapper.AdminMapper
	verify      *redis.Verify
}

func (s *Service) SLogin(ctx context.Context, loginDto *dto.AdminLoginDto, ticket string) (*vo.AdminVO, error) {
	_, err := s.verify.GetCaptchaTicket(ctx, ticket)
	if err != nil {
		return nil, err
	}
	admin, err := s.adminMapper.GetAdminByMobile(ctx, loginDto.Mobile)
	if err != nil {
		return nil, err
	}
	if !md5.MD5Verify(loginDto.Password, admin.Password) || admin.Status == constants.UserBanStatus {
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "手机号或密码错误")
	}
	return &vo.AdminVO{
		Id:       admin.ID,
		Name:     admin.Name,
		Nickname: admin.NickName,
		Token:    random.GenUUId(),
	}, nil
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
