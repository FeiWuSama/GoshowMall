package admin

import (
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor"
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

func (c *Ctrl) HelloWorld(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminDto := c.GetAdminDto(ctx, *c.adaptor, token)
	//helloWorld := c.adminService.HelloWorld(ctx.Request.Context())
	result.NewResultWithOk(ctx, adminDto)
}
