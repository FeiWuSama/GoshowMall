package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/constants"
)

type BaseCtrl struct {
}

func (c *BaseCtrl) GetUserDto(ctx *gin.Context, adaptor adaptor.Adaptor, token string) *dto.UserDto {
	result := adaptor.Redis.Get(ctx, constants.UserTokenKey+token)
	if result == nil {
		return nil
	}
	val := result.Val()
	var userDto *dto.UserDto
	if err := json.Unmarshal([]byte(val), &userDto); err != nil {
		return nil
	}
	return userDto
}

func (c *BaseCtrl) GetAdminDto(ctx *gin.Context, adaptor adaptor.Adaptor, token string) *dto.UserDto {
	result := adaptor.Redis.Get(ctx, constants.AdminTokenKey+token)
	if result == nil {
		return nil
	}
	val := result.Val()
	var userDto *dto.UserDto
	if err := json.Unmarshal([]byte(val), &userDto); err != nil {
		return nil
	}
	return userDto
}
