package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/wenlng/go-captcha/v2/slide"
	"go.uber.org/zap"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/api"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/logic/admin"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/service"
	"workspace-goshow-mall/utils/captcha"
	"workspace-goshow-mall/utils/logger"
)

type Ctrl struct {
	api.BaseCtrl
	adaptor      *adaptor.Adaptor
	adminService service.IAdminService
	verify       *redis.Verify
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor:      adaptor,
		adminService: admin.NewService(adaptor),
		verify:       redis.NewVerify(adaptor.Redis),
	}
}

// CreateAdmin
// @Summary 创建用户
// @Tags admin
// @Accept json
// @Produce json
// @host localhost:8080
// @param adminDto body dto.AddAdminDto true "用户信息"
// @Router /api/admin/create [post]
func (c *Ctrl) CreateAdmin(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo, err := c.GetAdminVo(ctx.Request.Context(), c.adaptor, token)
	if err != nil {
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
	result.NewResultWithOk[any](ctx, nil)
}

// UpdateAdmin
// @Summary 更新用户
// @Tags admin
// @Accept json
// @Produce json
// @host localhost:8080
// @param adminDto body dto.UpdateAdminDto true "用户信息"
// @Router /api/admin/update [post]
func (c *Ctrl) UpdateAdmin(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo, err := c.GetAdminVo(ctx.Request.Context(), c.adaptor, token)
	if err != nil {
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
	result.NewResultWithOk[any](ctx, nil)
}

// ChangeStatus
// @Summary 改变用户状态
// @Tags admin
// @Accept json
// @Produce json
// @host localhost:8080
// @param id path int true "用户id"
// @param status path int true "用户状态"
// @Router /api/admin/status/{id}/{status} [post]
func (c *Ctrl) ChangeStatus(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo, err := c.GetAdminVo(ctx.Request.Context(), c.adaptor, token)
	if err != nil {
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
	result.NewResultWithOk[any](ctx, nil)
}

// GetSlideCaptcha
// @Summary 获取滑块验证码
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} result.Result[vo.SlideCaptchaVo]
// @host localhost:8080
// @Router /api/admin/captcha/slide [get]
func (c *Ctrl) GetSlideCaptcha(ctx *gin.Context) {
	captchaDto := &dto.SlideCaptchaDto{}
	if err := ctx.ShouldBindQuery(captchaDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		ctx.Abort()
		return
	}
	newCaptcha := captcha.NewCaptcha()
	var mbs64Data, tbs64Data string
	captchaData, err := newCaptcha.Generate()
	if err != nil {
		ctx.Abort()
		logger.Error("captcha error", zap.Error(err))
		return
	}
	dotData, err := json.Marshal(captchaData.GetData())
	if err != nil {
		ctx.Abort()
		logger.Error("captcha error", zap.Error(err))
		return
	}
	mbs64Data, err = captchaData.GetMasterImage().ToBase64()
	if err != nil {
		ctx.Abort()
		logger.Error("captcha error", zap.Error(err))
		return
	}
	tbs64Data, err = captchaData.GetTileImage().ToBase64()
	if err != nil {
		ctx.Abort()
		logger.Error("captcha error", zap.Error(err))
		return
	}
	key := uuid.New().String()
	err = c.verify.SaveCaptcha(ctx, key, string(dotData))
	errorIf := result.ErrorIf(ctx, err)
	if errorIf {
		return
	}
	result.NewResultWithOk[vo.SlideCaptchaVo](ctx, vo.SlideCaptchaVo{
		Key:              key,
		ImageBase64:      mbs64Data,
		TitleImageBase64: tbs64Data,
		TitleHeight:      captchaData.GetData().Width,
		TitleWidth:       captchaData.GetData().Height,
		TitleX:           captchaData.GetData().DY,
		TitleY:           captchaData.GetData().DY,
	})
}

// VerifySlideCaptcha
// @Summary 验证滑块验证码
// @Tags admin
// @Accept json
// @Produce json
// @param SlideCaptchaCheckDto body dto.SlideCaptchaCheckDto true "校验信息"
// @Success 200 {object} result.Result[vo.SlideCaptchaCheckVo]
// @host localhost:8080
// @Router /api/admin/captcha/slide/verify [post]
func (c *Ctrl) VerifySlideCaptcha(ctx *gin.Context) {
	slideCaptchaCheckDto := &dto.SlideCaptchaCheckDto{}
	if err := ctx.ShouldBindJSON(slideCaptchaCheckDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		logger.Error("captcha error", zap.Error(err))
		ctx.Abort()
		return
	}
	captchaData, err := c.verify.GetCaptcha(ctx.Request.Context(), slideCaptchaCheckDto.Key)
	errorIf := result.ErrorIf(ctx, err)
	if errorIf {
		return
	}
	dot := slide.Block{}
	err = json.Unmarshal([]byte(captchaData), &dot)
	if err != nil {
		logger.Error("json paste error", zap.Error(err))
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		ctx.Abort()
		return
	}
	validate := slide.Validate(slideCaptchaCheckDto.SlideX, slideCaptchaCheckDto.SlideY, dot.X, dot.Y, 5)
	if !validate {
		result.NewResultWithError(ctx, nil, result.NewBusinessErrorWithMsg(result.ParamError, "验证码错误"))
		ctx.Abort()
		return
	}
	ticket := uuid.New().String()
	jsonData, err := json.Marshal(slideCaptchaCheckDto)
	if err != nil {
		logger.Error("convert json error", zap.Error(err))
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		ctx.Abort()
		return
	}
	err = c.verify.SaveCaptchaTicket(ctx.Request.Context(), constants.CaptchaTicketKey+ticket, string(jsonData))
	errorIf = result.ErrorIf(ctx, err)
	if errorIf {
		return
	}
	result.NewResultWithOk[vo.SlideCaptchaCheckVo](ctx, vo.SlideCaptchaCheckVo{
		Ticket: ticket,
		Expire: constants.CaptchaTicketExpire,
	})
}

// Login
// @Summary 管理员登录
// @Tags admin
// @Accept json
// @Produce json
// @param loginDto body dto.AdminLoginDto true "管理员登录信息"
// @Success 200 {object} result.Result[vo.AdminVO]
// @host localhost:8080
// @Router /api/admin/login [post]
func (c *Ctrl) Login(ctx *gin.Context) {
	loginDto := &dto.AdminLoginDto{}
	if err := ctx.ShouldBindJSON(loginDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
	}
	ticket := ctx.Request.Header.Get("captcha-ticket")
	adminVo, err := c.adminService.SLogin(ctx.Request.Context(), loginDto, ticket)
	errorIf := result.ErrorIf(ctx, err)
	if errorIf {
		return
	}
	result.NewResultWithOk[vo.AdminVO](ctx, *adminVo)
}
