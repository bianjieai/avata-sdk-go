package models

import "fmt"

const CodeSpace = "AVATA-SDK-GO"

//code
const (
	UnknownError = "UNKNOWN_ERROR" // 未知错误
	BadRequest   = "BAD_REQUEST"   // 参数错误
)

// Message
const (
	ErrDomain    = "the avata domain address is required"
	ErrAPIKey    = "the api key for the project is required"
	ErrAPISecret = "the api secret the project is required"
	ErrAmount    = "amount should be integer yuan"

	ErrParam = "%s is required"
)

type (
	Error interface {
		Error() string
		CodeSpace() string
		Code() string
		Msg() string
	}

	//ErrorRes Avata 错误提示信息
	ErrorRes struct {
		codeSpace string `json:"code_space"` // 命名空间
		code      string `json:"code"`       // 错误码
		message   string `json:"message"`    // 错误描述
	}
)

func (e ErrorRes) Error() string {
	return fmt.Sprintf("code_space: %s, code: %s, message: %s", e.codeSpace, e.code, e.message)
}

func (e ErrorRes) CodeSpace() string {
	return e.codeSpace
}

func (e ErrorRes) Code() string {
	return e.code
}

func (e ErrorRes) Msg() string {
	return e.message
}

func NewSDKError(message string) Error {
	return ErrorRes{
		codeSpace: CodeSpace,
		code:      UnknownError,
		message:   message,
	}
}

func NewAvataError(avataError AvataError) Error {
	return ErrorRes{
		codeSpace: avataError.CodeSpace,
		code:      avataError.Code,
		message:   avataError.Message,
	}
}

func InvalidParam(message string) Error {
	return ErrorRes{
		codeSpace: CodeSpace,
		code:      BadRequest,
		message:   message,
	}
}
