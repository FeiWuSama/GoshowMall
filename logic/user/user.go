package user

import (
	"context"
	"fmt"
	"github.com/cnchef/gconv"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"workspace-goshow-mall/adaptor"
	myRedis "workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/mapper"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/rpc"
	"workspace-goshow-mall/service"
	"workspace-goshow-mall/utils/logger"
	"workspace-goshow-mall/utils/md5"
)

type Service struct {
	adapter     *adaptor.Adaptor
	userMapper  mapper.UserMapper
	verify      *myRedis.Verify
	larkService service.ILarkService
	larkRpc     *rpc.LarkRpc
}

func (s Service) SLogin(context context.Context, userMobileLoginDto interface{}) (*vo.UserVo, error) {
	var user *model.User
	var err error
	switch userMobileLoginDto.(type) {
	case *dto.UserMobilePasswordLoginDto:
		user, err = s.getUserByPassword(context, userMobileLoginDto.(*dto.UserMobilePasswordLoginDto))
		if err != nil {
			return nil, err
		}
	case *dto.UserLarkLoginDto:
		user, err = s.getUserByLark(context, userMobileLoginDto.(*dto.UserLarkLoginDto))
		if err != nil {
			return nil, err
		}
	}
	token := uuid.New().String()
	userVo := &vo.UserVo{
		Token:    token,
		Id:       user.ID,
		Nickname: user.NickName,
		Avatar:   user.Avatar,
		Sex:      user.Sex,
	}
	err = s.verify.SaveUserToken(context, token, gconv.ToString(userVo))
	if err != nil {
		return nil, err
	}
	return userVo, nil
}

func (s Service) getUserByPassword(context context.Context, loginDto *dto.UserMobilePasswordLoginDto) (*model.User, error) {
	_, err := s.verify.GetCaptchaTicket(context, loginDto.Ticket)
	if err != nil {
		logger.Error("verify error", zap.Error(err))
		return nil, err
	}
	count, err := s.verify.IncrPasswordErrorCount(context, loginDto.Mobile)
	if err != nil {
		logger.Error("redis error", zap.Error(err))
		return nil, err
	}
	if count > constants.PasswordErrorCount {
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, fmt.Sprintf("密码错误次数过多,请在%d分钟后重试", constants.PasswordErrorCount))
	}
	user, err := s.userMapper.GetUserByMobile(context, loginDto)
	if err != nil {
		logger.Error("not found user error", zap.Error(err))
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "手机号或密码错误")
	}
	if !md5.MD5Verify(loginDto.Password, user.Password) || user.Status == constants.UserBanStatus {
		logger.Error("password error")
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "手机号或密码错误")
	}
	_ = s.verify.DeletePasswordErrorCount(context, loginDto.Mobile)
	return user, nil
}

func (s Service) getUserByLark(ctx context.Context, dto *dto.UserLarkLoginDto) (*model.User, error) {
	accessTokenVo, err := s.larkService.SLarkGetToken(ctx, dto.AppCode, dto.Code, dto.RedirectUri, "")
	if err != nil {
		logger.Error("get lark user access token err", zap.Error(err))
		return nil, err
	}
	larkUserInfo, err := s.larkRpc.GetLarkUserInfo(ctx, accessTokenVo.AccessToken)
	if err != nil {
		logger.Error("get lark user info err", zap.Error(err))
		return nil, err
	}
	user, err := s.userMapper.GetUserByOpenIdAndCode(ctx, larkUserInfo.OpenId, dto.AppCode)
	if err != nil {
		logger.Error("get user by open id and code error", zap.Error(err))
		return nil, err
	}
	if user.Status == constants.UserBanStatus {
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "飞书登录失败")
	}
	return user, nil
}
