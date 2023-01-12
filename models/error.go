package models

// Code
const (
	CodeSuccess = 0  // 成功
	CodeFailed  = -1 // 失败
)

// Message
const (
	ErrDomain    = "the avata domain address is required"
	ErrAPIKey    = "the api key for the project is required"
	ErrAPISecret = "the api secret the project is required"
	ErrAmount    = "amount should be integer yuan"

	ErrParam = "%s is required"
)
