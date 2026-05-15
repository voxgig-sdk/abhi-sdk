package core

type AbhiError struct {
	IsAbhiError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewAbhiError(code string, msg string, ctx *Context) *AbhiError {
	return &AbhiError{
		IsAbhiError: true,
		Sdk:              "Abhi",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *AbhiError) Error() string {
	return e.Msg
}
