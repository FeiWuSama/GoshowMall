package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

func NewResultWithOk[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, Result[T]{
		Code: OK.Code,
		Data: data,
		Msg:  OK.Msg,
	})
}

func NewResultWithError(ctx *gin.Context, data any, err *BusinessError) {
	ctx.JSON(http.StatusOK, Result[any]{
		Code: err.Code,
		Data: data,
		Msg:  err.Msg,
	})
}
