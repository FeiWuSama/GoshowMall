package user

import (
	"errors"
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
	"workspace-goshow-mall/logic/user"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/service"
	"workspace-goshow-mall/utils/captcha"
	"workspace-goshow-mall/utils/logger"
)

type Ctrl struct {
	api.BaseCtrl
	adaptor     *adaptor.Adaptor
	verify      *redis.Verify
	userService service.IUserService
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor:     adaptor,
		verify:      redis.NewVerify(adaptor.Redis),
		userService: user.NewService(adaptor),
	}
}

// GetSlideCaptcha
// @Summary 获取滑块验证码
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} result.Result[vo.SlideCaptchaVo]
// @host localhost:8080
// @Router /api/user/captcha/slide [get]
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
	if err != nil {
		ctx.Abort()
		logger.Error("captcha error", zap.Error(err))
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
// @Tags user
// @Accept json
// @Produce json
// @param SlideCaptchaCheckDto body dto.SlideCaptchaCheckDto true "校验信息"
// @Success 200 {object} result.Result[vo.SlideCaptchaCheckVo]
// @host localhost:8080
// @Router /api/user/captcha/slide/verify [post]
func (c *Ctrl) VerifySlideCaptcha(ctx *gin.Context) {
	slideCaptchaCheckDto := &dto.SlideCaptchaCheckDto{}
	if err := ctx.ShouldBindJSON(slideCaptchaCheckDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		logger.Error("captcha error", zap.Error(err))
		ctx.Abort()
		return
	}
	captchaData, err := c.verify.GetCaptcha(ctx.Request.Context(), slideCaptchaCheckDto.Key)
	if err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessErrorWithMsg(result.ParamError, "验证码已过期"))
		ctx.Abort()
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
	if err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		ctx.Abort()
		return
	}
	result.NewResultWithOk[vo.SlideCaptchaCheckVo](ctx, vo.SlideCaptchaCheckVo{
		Ticket: ticket,
		Expire: constants.CaptchaTicketExpire,
	})
}

// MobileLoginByPassword
// @Summary 手机号登录
// @Tags user
// @Accept json
// @Produce json
// @param userMobileLoginDto body dto.UserMobilePasswordLoginDto true "手机号登录信息"
// @Success 200 {object} result.Result[vo.UserVo]
// @host localhost:8080
// @Router /api/user/mobile/login/password [post]
func (c *Ctrl) MobileLoginByPassword(ctx *gin.Context) {
	userMobileLoginDto := &dto.UserMobilePasswordLoginDto{}
	if err := ctx.ShouldBindJSON(userMobileLoginDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		ctx.Abort()
		return
	}
	userMobileLoginDto.Ticket = ctx.Request.Header.Get("captcha-ticket")
	userVo, err := c.userService.SLogin(ctx.Request.Context(), userMobileLoginDto)
	if err != nil {
		if errors.As(err, &result.BusinessError{}) {
			result.NewResultWithError(ctx, nil, err.(*result.BusinessError))
		} else {
			result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		}
		ctx.Abort()
		return
	}
	result.NewResultWithOk[vo.UserVo](ctx, *userVo)
}

// GetUserInfo
// @Summary 获取用户信息
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} result.Result[vo.UserVo]
// @host localhost:8080
// @Router /api/user/info [get]
func (c *Ctrl) GetUserInfo(ctx *gin.Context) {
	userVo, err := c.GetUserVo(ctx.Request.Context(), c.adaptor, ctx.Request.Header.Get(constants.UserToken))
	if err != nil {
		logger.Error("get user info error", zap.Error(err))
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		ctx.Abort()
		return
	}
	result.NewResultWithOk[vo.UserVo](ctx, *userVo)
}

// LoginByLark
// @Summary 飞书登录
// @Tags user
// @Accept json
// @Produce json
// @param UserLarkLoginDto body dto.UserLarkLoginDto true "飞书登录信息"
// @Success 200 {object} result.Result[vo.UserVo]
// @host localhost:8080
// @Router /api/user/lark/login [post]
func (c *Ctrl) LoginByLark(ctx *gin.Context) {
	userLarkLoginDto := &dto.UserLarkLoginDto{}
	if err := ctx.ShouldBindJSON(userLarkLoginDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		ctx.Abort()
		return
	}
	userVo, err := c.userService.SLogin(ctx.Request.Context(), userLarkLoginDto)
	if err != nil {
		if errors.As(err, &result.BusinessError{}) {
			result.NewResultWithError(ctx, nil, err.(*result.BusinessError))
		} else {
			result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		}
		ctx.Abort()
		return
	}
	result.NewResultWithOk[vo.UserVo](ctx, *userVo)
}

// PostMobileSmsCode
// @Summary 手机号短信验证码登录
// @Tags user
// @Accept json
// @Produce json
// @param ticket query string true "ticket"
// @param mobile query string true "mobile"
// @param scene query string true "场景"
// @Success 200 {object} result.Result[any]
// @host localhost:8080
// @Router /api/user/mobile/smsCode [post]
func (c *Ctrl) PostMobileSmsCode(ctx *gin.Context) {
	ticket := ctx.Query("ticket")
	mobile := ctx.Query("mobile")
	scene := ctx.Query("scene")
	err := c.userService.SPostMobileSmsCode(ctx.Request.Context(), ticket, mobile, scene)
	if err != nil {
		if errors.As(err, &result.BusinessError{}) {
			result.NewResultWithError(ctx, nil, err.(*result.BusinessError))
		} else {
			result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		}
		ctx.Abort()
		return
	}
	result.NewResultWithOk[any](ctx, nil)
}

// MobileLoginBySmsCode
// @Summary 手机号短信验证码登录
// @Tags user
// @Accept json
// @Produce json
// @param userMobileSmsLoginDto body dto.UserMobileSmsLoginDto true "手机号短信验证码登录信息"
// @Success 200 {object} result.Result[vo.UserVo]
// @host localhost:8080
// @Router /api/user/mobile/smsCode/verify [post]
func (c *Ctrl) MobileLoginBySmsCode(ctx *gin.Context) {
	userMobileSmsLoginDto := &dto.UserMobileSmsLoginDto{}
	if err := ctx.ShouldBindJSON(userMobileSmsLoginDto); err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
	}
	userVo, err := c.userService.SLogin(ctx.Request.Context(), userMobileSmsLoginDto)
	if err != nil {
		if errors.As(err, &result.BusinessError{}) {
			result.NewResultWithError(ctx, nil, err.(*result.BusinessError))
		} else {
			result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ServerError))
		}
		ctx.Abort()
		return
	}
	result.NewResultWithOk[vo.UserVo](ctx, *userVo)
}
