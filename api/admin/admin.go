package admin

import (
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/api"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/logic/admin"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/service"
)

type Ctrl struct {
	api.BaseCtrl
	adaptor      *adaptor.Adaptor
	adminService service.IAdminService
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor:      adaptor,
		adminService: admin.NewService(adaptor),
	}
}

// CreateAdmin
// @Summary 创建用户
// @Tags 管理员接口
// @Accept json
// @Produce json
// @host localhost:8080
// @param adminDto body dto.AddAdminDto true "用户信息"
// @Router /api/admin/create [post]
func (c *Ctrl) CreateAdmin(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo := c.GetAdminVo(ctx, *c.adaptor, token)
	if adminVo == nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.Unauthorized))
		ctx.Abort()
		return
	}
	adminDto := dto.AddAdminDto{}
	if err := ctx.ShouldBindJSON(&adminDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		ctx.Abort()
		return
	}
	rows := c.adminService.SCreateAdmin(ctx.Request.Context(), adminDto, int64(adminVo.Id))
	if rows == 0 {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		ctx.Abort()
		return
	}
	result.NewResultWithOk(ctx, nil)
}

// UpdateAdmin
// @Summary 更新用户
// @Tags 管理员接口
// @Accept json
// @Produce json
// @host localhost:8080
// @param adminDto body dto.UpdateAdminDto true "用户信息"
// @Router /api/admin/update [post]
func (c *Ctrl) UpdateAdmin(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo := c.GetAdminVo(ctx, *c.adaptor, token)
	if adminVo == nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.Unauthorized))
		ctx.Abort()
		return
	}
	adminDto := dto.UpdateAdminDto{}
	if err := ctx.ShouldBindJSON(&adminDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		ctx.Abort()
		return
	}
	isSuccess := c.adminService.SUpdateAdmin(ctx.Request.Context(), adminDto, int64(adminVo.Id))
	if !isSuccess {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		ctx.Abort()
		return
	}
	result.NewResultWithOk(ctx, nil)
}

// ChangeStatus
// @Summary 改变用户状态
// @Tags 管理员接口
// @Accept json
// @Produce json
// @host localhost:8080
// @param id path int true "用户id"
// @param status path int true "用户状态"
// @Router /api/admin/status/{id}/{status} [post]
func (c *Ctrl) ChangeStatus(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo := c.GetAdminVo(ctx, *c.adaptor, token)
	if adminVo == nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.Unauthorized))
		ctx.Abort()
		return
	}
	id := ctx.Param("id")
	status := ctx.Param("status")
	isSuccess := c.adminService.SChangeStatus(ctx.Request.Context(), id, status, int64(adminVo.Id))
	if !isSuccess {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		ctx.Abort()
		return
	}
	result.NewResultWithOk(ctx, nil)
}
