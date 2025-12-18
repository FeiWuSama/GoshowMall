package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/result"
)

type TokenFunc func(c context.Context, token string) (*vo.UserVo, error)
type AdminTokenFunc func(c context.Context, token string) (*vo.AdminVO, error)

func UserAuthMiddleware(filter func(ctx *gin.Context) bool, tokenFunc TokenFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		if filter != nil && !filter(context) {
			context.Next()
			return
		}
		// 鉴权中间件
		token := context.GetHeader(constants.UserToken)
		if len(token) == 0 {
			result.NewResultWithError(context, nil, result.NewBusinessError(result.Unauthorized))
			context.Abort()
			return
		}
		_, err := tokenFunc(context, token)
		if err != nil {
			result.NewResultWithError(context, nil, result.NewBusinessError(result.PermissionDenied))
			context.Abort()
			return
		}
		if err != nil {
			return
		}
		context.Next()
	}
}

func AdminAuthMiddleware(filter func(ctx *gin.Context) bool, adminTokenFunc AdminTokenFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		if filter != nil && !filter(context) {
			context.Next()
			return
		}
		// 鉴权中间件
		token := context.GetHeader(constants.AdminToken)
		if len(token) == 0 {
			result.NewResultWithError(context, nil, result.NewBusinessError(result.Unauthorized))
			context.Abort()
			return
		}
		_, err := adminTokenFunc(context, token)
		if err != nil {
			result.NewResultWithError(context, nil, result.NewBusinessError(result.PermissionDenied))
			context.Abort()
			return
		}
		context.Next()
	}
}
