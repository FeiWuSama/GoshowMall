package permission

import (
	"github.com/cnchef/gconv"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/vo"
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

// GetPermissionPage
// @Summary 分页获取权限列表
// @Tags adminPermission
// @Accept json
// @Produce json
// @host localhost:8080
// @param adminDto body dto.PageDto true "分页请求"
// @Success 200 {object} result.Result[vo.PermissionPage]
// @Router /api/admin/permission/page [post]
func (c *Ctrl) GetPermissionPage(ctx *gin.Context) {
	d := &dto.PageDto{}
	err := ctx.ShouldBindJSON(d)
	if err != nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.ParamError))
		return
	}
	page, err := c.permissionService.SGetAllPermission(ctx.Request.Context(), d)
	if errorIf := result.ErrorIf(ctx, err); errorIf {
		return
	}
	result.NewResultWithOk[paginator.Page[*model.Permission]](ctx, *page)
}

// GetPermissionTree
// @Summary 获取权限树
// @Tags adminPermission
// @Accept json
// @Produce json
// @Success 200 {object} result.Result[vo.PermissionVoList]
// @host localhost:8080
// @Router /api/admin/permission/tree [get]
func (c *Ctrl) GetPermissionTree(ctx *gin.Context) {
	permissionTree, err := c.permissionService.ConvertPermissionList2Tree(ctx, 0)
	if errorIf := result.ErrorIf(ctx, err); errorIf {
		return
	}
	result.NewResultWithOk[[]*vo.PermissionVo](ctx, permissionTree)
}

// GetPermissionByRoleId
// @Summary 根据角色ID获取权限列表
// @Tags adminPermission
// @Accept json
// @Produce json
// @host localhost:8080
// @param roleId query int true "角色ID"
// @Router /api/admin/role/permission [post]
// @Success 200 {object} result.Result[vo.PermissionList]
func (c *Ctrl) GetPermissionByRoleId(ctx *gin.Context) {
	roleId := ctx.Param("roleId")
	permissions, err := c.permissionService.SGetPermissionByRoleId(ctx.Request.Context(), int64(gconv.ToInt(roleId)))
	if errorIf := result.ErrorIf(ctx, err); errorIf {
		return
	}
	result.NewResultWithOk[[]*model.Permission](ctx, permissions)
}
