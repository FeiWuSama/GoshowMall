package user

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/utils/captcha"
	"workspace-goshow-mall/utils/logger"
)

type Ctrl struct {
	adaptor *adaptor.Adaptor
	verify  *redis.Verify
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor: adaptor,
		verify:  redis.NewVerify(adaptor.Redis),
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
