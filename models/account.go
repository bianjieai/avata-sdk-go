package models

const (
	CreateAccount       = "/v1beta1/account"          // 创建链账户接口
	BatchCreateAccounts = "/v1beta1/accounts"         // 批量创建链账户接口
	GetAccounts         = "/v1beta1/accounts"         // 查询链账户接口
	GetAccountsHistory  = "/v1beta1/accounts/history" // 查询链账户操作记录接口
)

// CreateAccountReq 创建链账户请求参数
type CreateAccountReq struct {
	Name        string `json:"name"`
	OperationID string `json:"operation_id"`
}

// CreateAccountRes 创建链账户返回值
type CreateAccountRes struct {
	Data struct {
		Account     string `json:"account"`
		Name        string `json:"name"`
		OperationID string `json:"operation_id"`
	} `json:"data"`
}
