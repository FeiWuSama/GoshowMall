package role

import (
	"github.com/cnchef/gconv"
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/api"
	"workspace-goshow-mall/logic/permission"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/service"
)

type Ctrl struct {
	api.BaseCtrl
	adaptor           *adaptor.Adaptor
	permissionService service.IPermissionService
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor:           adaptor,
		permissionService: permission.NewService(adaptor),
	}
}

// GetPermissionByRoleId
// @Summary 根据角色ID获取权限列表
// @Tags adminPermission
// @Accept json
// @Produce json
// @host localhost:8080
// @param roleId query int true "角色ID"
// @Router /api/admin/role/permission [post]
// @Success 200 {object} result.Result[[]*model.Permission]
func (c *Ctrl) GetPermissionByRoleId(ctx *gin.Context) {
	roleId := ctx.Query("roleId")
	permissions, err := c.permissionService.SGetPermissionByRoleId(ctx.Request.Context(), int64(gconv.ToInt(roleId)))
	if errorIf := result.ErrorIf(ctx, err); errorIf {
		return
	}
	result.NewResultWithOk[[]*model.Permission](ctx, permissions)
}
