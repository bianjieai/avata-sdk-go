package models

const (
	CreateAccount        = "/v1beta1/account"          // 创建链账户接口
	BatchCreateAccounts  = "/v1beta1/accounts"         // 批量创建链账户接口
	QueryAccounts        = "/v1beta1/accounts"         // 查询链账户接口
	QueryAccountsHistory = "/v1beta1/accounts/history" // 查询链账户操作记录接口
)

// CreateAccountReq 创建链账户请求参数
type CreateAccountReq struct {
	Name        string `json:"name"`
	OperationID string `json:"operation_id"`
}

// CreateAccountRes 创建链账户返回值
type (
	CreateAccountRes struct {
		Data    struct {
			Account     string `json:"account"`
			Name        string `json:"name"`
			OperationID string `json:"operation_id"`
		} `json:"data"`
	}
)

// BatchCreateAccountsReq 批量创建链账户请求参数
type BatchCreateAccountsReq struct {
	Count       int    `json:"count,omitempty"`
	OperationID string `json:"operation_id"`
}

// BatchCreateAccountsRes 批量创建链账户返回值
type BatchCreateAccountsRes struct {
	Data struct {
		Accounts    []string `json:"accounts"`
		OperationID string   `json:"operation_id"`
	} `json:"data"`
}

// QueryAccountsReq 查询链账户请求参数
type QueryAccountsReq struct {
	Offset      string `json:"offset,omitempty"`
	Limit       string `json:"limit,omitempty"`
	Account     string `json:"account,omitempty"`
	Name        string `json:"name,omitempty"`
	OperationID string `json:"operation_id,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	SortBy      string `json:"sort_by,omitempty"`
}

// QueryAccountsRes 查询链账户返回值
type (
	QueryAccountsRes struct {
		Data struct {
			Offset     int64      `json:"offset"`
			Limit      int64      `json:"limit"`
			TotalCount int64      `json:"total_count"`
			Accounts   []Accounts `json:"accounts"`
		} `json:"data"`
	}

	Accounts struct {
		Account     string `json:"account"`
		Name        string `json:"name"`
		Gas         int64  `json:"gas"`
		BizFee      int64  `json:"biz_fee"`
		OperationID string `json:"operation_id"`
		Status      int    `json:"status"`
	}
)

// QueryAccountsHistoryReq 查询链账户操作记录请求参数
type QueryAccountsHistoryReq struct {
	Offset    string `json:"offset,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Account   string `json:"account,omitempty"`
	Module    string `json:"module,omitempty"`
	Operation string `json:"operation,omitempty"`
	TxHash    string `json:"tx_hash,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
}

// QueryAccountsHistoryRes 查询链账户操作记录返回值
type (
	QueryAccountsHistoryRes struct {
		Data struct {
			Offset           int64              `json:"offset"`
			Limit            int64              `json:"limit"`
			TotalCount       int64              `json:"total_count"`
			OperationRecords []OperationRecords `json:"operation_records"`
		} `json:"data"`
	}
	OperationRecords struct {
		TxHash      string `json:"tx_hash"`
		Module      string `json:"module"`
		Operation   string `json:"operation"`
		Signer      string `json:"signer"`
		Timestamp   string `json:"timestamp"`
		GasFee      int64  `json:"gas_fee"`
		BusinessFee int64  `json:"business_fee"`
		NFTMsg      NFTMsg `json:"nft_msg"`
		MTMsg       MTMsg  `json:"mt_msg"`
	}
	NFTMsg struct {
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Symbol      string `json:"symbol,omitempty"`
		URI         string `json:"uri,omitempty"`
		Recipient   string `json:"recipient,omitempty"`
		ClassId     string `json:"class_id,omitempty"`
		ClassName   string `json:"class_name,omitempty"`
		ClassSymbol string `json:"class_symbol,omitempty"`
	}
	MTMsg struct {
		Id        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Recipient string `json:"recipient,omitempty"`
		ClassId   string `json:"class_id,omitempty"`
		ClassName string `json:"class_name,omitempty"`
		Amount    int    `json:"amount,omitempty"`
	}
)
