package result

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type BusinessError struct {
	Code int
	Msg  string
}

func (e BusinessError) Error() string {
	return e.Msg
}

func NewBusinessError(code Code) *BusinessError {
	return &BusinessError{
		Code: code.Code,
		Msg:  code.Msg,
	}
}

func NewBusinessErrorWithMsg(code Code, msg string) *BusinessError {
	return &BusinessError{
		Code: code.Code,
		Msg:  msg,
	}
}

func ErrorIf(ctx *gin.Context, err error) bool {
	if err != nil {
		if errors.As(err, &BusinessError{}) {
			NewResultWithError(ctx, nil, err.(*BusinessError))
		} else {
			NewResultWithError(ctx, nil, NewBusinessError(ServerError))
		}
		ctx.Abort()
		return true
	}
	return false
}
