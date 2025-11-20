package result

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
