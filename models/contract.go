package models

const (
	UseContract   = "/v3/evm/contract/calls" //调用合约
	QueryContract = "/v3/evm/contract/calls" //查询合约
)

type UseContractReq struct {
	From        string `json:"from"`                 //签名交易链账户地址
	To          string `json:"to"`                   //合约地址
	Data        string `json:"data"`                 //ABI 编码待签名交易的 hex 字符串
	GasLimit    int    `json:"gas_limit"`            //gas 使用上限大小
	Estimation  int    `json:"estimation,omitempty"` //是否模拟执行交易 0 不模拟 1 模拟
	OperationID string `json:"operation_id"`         //操作 ID
}

type QueryContractReq struct {
	To   string `json:"to"`   //合约地址
	Data string `json:"data"` //ABI 编码待签名交易的 hex 字符串
}

type QueryContractRes struct {
	Data struct {
		Result string `json:"result"`
	} `json:"data"`
}
