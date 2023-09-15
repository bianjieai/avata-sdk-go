package models

// BaseParams 初始化 SDK 客户端参数
type BaseParams struct {
	Domain    string // 域名
	APIKey    string // 项目参数 API KEY
	APISecret string // 项目参数 API SECRET
}

// TxRes 通用正确返回值(所有发起上链交易接口)
type TxRes struct {
	Data struct {
	} `json:"data"`
}

// Response Avata 错误提示信息
type (
	Response struct {
		AvataError AvataError `json:"error"`
	}
	AvataError struct {
		CodeSpace string `json:"code_space"` // 命名空间
		Code      string `json:"code"`       // 错误码
		Message   string `json:"message"`    // 错误描述
	}
)
