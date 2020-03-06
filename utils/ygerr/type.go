package ygerr

import "errors"

type YgErrorCode int

const (
	ClientError YgErrorCode = iota + 1
	ServerError
	UnKnown
)

type YgError interface {
	Code() YgErrorCode
	Error() string
}

// web 端 controller 错误
// 主要借助 code 字段，用于区分错误为客户端产生还是服务端产生
type WebCtlErr struct {
	code YgErrorCode
	err  error
}

func (e *WebCtlErr) Error() string {
	return e.err.Error()
}

func (e *WebCtlErr) Code() YgErrorCode {
	return e.code
}

func NewWebCtl(code YgErrorCode, msg string) *WebCtlErr {
	return &WebCtlErr{
		code: code,
		err:  errors.New(msg),
	}
}
