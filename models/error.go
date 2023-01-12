package models

// Code
const (
	CodeSuccess = 0  // 成功
	CodeFailed  = -1 // 失败
)

// SDK 错误描述
const (
	ErrDomain    = "the avata domain address is required"
	ErrAPIKey    = "the api key for the project is required"
	ErrAPISecret = "the api secret the project is required"

	ErrParam = "%s is required"
)
