package models

// BaseParams 初始化 SDK 客户端参数
type BaseParams struct {
	Domain    string // 域名
	APIKey    string // 项目参数 API KEY
	APISecret string // 项目参数 API SECRET
}

// 通用返回参数
type (
	BaseRes struct {
		Code    int    `json:"code"`    // 调用 SDK 方法是否成功
		Http    Http   `json:"http"`    // HTTP 响应状态码
		Message string `json:"message"` // 调用 SDK 方法失败提示信息
		Error   Error  `json:"error"`   // 接口返回的错误提示信息
	}

	Http struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Error struct {
		CodeSpace string `json:"code_space"`
		Code      string `json:"code"`
		Message   string `json:"message"`
	}
)

// TxRes 发起上链交易接口返回参数
type TxRes struct {
	BaseRes
	Data struct {
		OperationId string `json:"operation_id"`
	} `json:"data"`
}
