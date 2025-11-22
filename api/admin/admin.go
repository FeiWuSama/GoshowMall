package admin

import (
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/logic/admin"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/service"
)

type Ctrl struct {
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
	helloWorld := c.adminService.HelloWorld(ctx.Request.Context())
	result.NewResultWithOk(ctx, helloWorld)
}
