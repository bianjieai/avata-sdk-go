package models

// BaseParams 初始化 SDK 客户端参数
type BaseParams struct {
	Domain    string // 域名
	APIKey    string // 项目参数 API KEY
	APISecret string // 项目参数 API SECRET
}

// 通用返回参数
type (
	Response struct {
		Code    int         `json:"code"`    // SDK 响应状态码
		Http    Http        `json:"http"`    // HTTP 响应状态码
		Message string      `json:"message"` // SDK 返回的错误提示信息
		Error   Error       `json:"error"`   // 接口返回的错误提示信息
		Data    interface{} `json:"data"`    // 接口返回的参数
	}

	Http struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Error struct {
		CodeSpace string `json:"code_space"` // 命名空间
		Code      string `json:"code"`       // 错误码
		Message   string `json:"message"`    // 错误描述
	}
)

// TxRes 通用正确返回值(所有发起上链交易接口)
type TxRes struct {
	Data struct {
		OperationId string `json:"operation_id"` // 操作 ID
	} `json:"data"`
}
