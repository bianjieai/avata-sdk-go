package models

const (
	CreateAccount        = "/v1beta1/account"          // 创建链账户接口
	BatchCreateAccounts  = "/v1beta1/accounts"         // 批量创建链账户接口
	QueryAccounts        = "/v1beta1/accounts"         // 查询链账户接口
	QueryAccountsHistory = "/v1beta1/accounts/history" // 查询链账户操作记录接口
)

// CreateAccountReq 创建链账户请求参数
type CreateAccountReq struct {
	Name        string `json:"name"`         // 链账户名称，支持 1-20 位汉字、大小写字母及数字组成的字符串
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串组成。此操作 ID 仅限在查询链账户接口中使用，用于查询创建链账户的授权状态。
}

// CreateAccountRes 创建链账户正确返回值
type (
	CreateAccountRes struct {
		Data struct {
			Account     string `json:"account"`      // 链账户地址
			Name        string `json:"name"`         // 链账户名称
			OperationID string `json:"operation_id"` // 操作 ID。此操作 ID 仅限在查询链账户接口中使用，用于查询创建链账户的授权状态。
		} `json:"data"`
	}
)

// BatchCreateAccountsReq 批量创建链账户请求参数
type BatchCreateAccountsReq struct {
	Count       int    `json:"count,omitempty"` // 批量创建链账户的数量
	OperationID string `json:"operation_id"`    // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串。此操作 ID 仅限在查询链账户接口中使用，用于查询创建链账户的授权状态。
}

// BatchCreateAccountsRes 批量创建链账户正确返回值
type BatchCreateAccountsRes struct {
	Data struct {
		Accounts    []string `json:"accounts"`     // 链账户地址列表
		OperationID string   `json:"operation_id"` // 操作 ID。此操作 ID 仅限在查询链账户接口中使用，用于查询创建链账户的授权状态。
	} `json:"data"`
}

// QueryAccountsReq 查询链账户请求参数
type QueryAccountsReq struct {
	Offset      string `json:"offset,omitempty"`       // 游标，默认为 0
	Limit       string `json:"limit,omitempty"`        // 每页记录数，默认为 10，上限为 50
	Account     string `json:"account,omitempty"`      // 链账户地址
	Name        string `json:"name,omitempty"`         // 链账户名称，支持模糊查询
	OperationID string `json:"operation_id,omitempty"` // 操作 ID。此操作 ID 需要填写在请求创建链账户/批量创建链账户接口时，返回的 Operation ID。
	StartDate   string `json:"start_date,omitempty"`   // 创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate     string `json:"end_date,omitempty"`     // 创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy      string `json:"sort_by,omitempty"`      // 排序规则：DATE_ASC / DATE_DESC
}

// QueryAccountsRes 查询链账户正确返回值
type (
	QueryAccountsRes struct {
		Data struct {
			Offset     int64      `json:"offset"`      // 游标
			Limit      int64      `json:"limit"`       // 每页记录数
			TotalCount int64      `json:"total_count"` // 总记录数
			Accounts   []Accounts `json:"accounts"`    // 链账户列表
		} `json:"data"`
	}

	Accounts struct {
		Account     string `json:"account"`      // 链账户地址
		Name        string `json:"name"`         // 链账户名称
		Gas         int64  `json:"gas"`          // 文昌链能量值余额
		BizFee      int64  `json:"biz_fee"`      // 文昌链 DDC 业务费余额，单位：分
		OperationID string `json:"operation_id"` // 操作 ID
		Status      int    `json:"status"`       // Enum: 0，1；链账户的授权状态，0 未授权；1 已授权。链账户授权成功后，可使用该链账户地址发起上链交易请求；未授权时不影响作为交易的接受者地址进行使用（DDC 业务除外）。
	}
)

// QueryAccountsHistoryReq 查询链账户操作记录请求参数
type QueryAccountsHistoryReq struct {
	Offset    string `json:"offset,omitempty"`     // 游标，默认为 0
	Limit     string `json:"limit,omitempty"`      // 每页记录数，默认为 10，上限为 50
	Account   string `json:"account,omitempty"`    // 链账户地址
	Module    string `json:"module,omitempty"`     // 功能模块；Enum: "nft" "mt"
	Operation string `json:"operation,omitempty"`  // 操作类型，仅 module 不为空时有效，默认为 "all"。 module = nft 时，可选：issue_class / transfer_class / mint / edit / transfer / burn； module = mt 时，可选： issue_class / transfer_class / issue / mint / edit / transfer / burn。
	TxHash    string `json:"tx_hash,omitempty"`    // Tx Hash
	StartDate string `json:"start_date,omitempty"` // 日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   // 日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    // 排序规则：DATE_ASC / DATE_DESC
}

// QueryAccountsHistoryRes 查询链账户操作记录返回值
type (
	QueryAccountsHistoryRes struct {
		Data struct {
			Offset           int64              `json:"offset"`            // 游标
			Limit            int64              `json:"limit"`             // 每页记录数
			TotalCount       int64              `json:"total_count"`       // 总记录数
			OperationRecords []OperationRecords `json:"operation_records"` // 操作记录列表
		} `json:"data"`
	}
	OperationRecords struct {
		TxHash      string `json:"tx_hash"`      // 操作 Tx Hash
		Module      string `json:"module"`       // 功能模块；Enum: "nft" "mt"
		Operation   string `json:"operation"`    // 操作类型；Enum: "issue_class" "transfer_class" "mint" "edit" "transfer" "burn" "issue"
		Signer      string `json:"signer"`       // Tx 签名者地址
		Timestamp   string `json:"timestamp"`    // 操作时间戳（UTC 时间）
		GasFee      int64  `json:"gas_fee"`      // 链上交易消耗的能量值；当前支持查询 2022 年 08 月 18 日 11:00:00(UTC 时间) 底层链升级固定 Gas 之后的数据，其它历史数据已归档，暂不支持查询对应结果
		BusinessFee int64  `json:"business_fee"` // 链上交易消耗的业务费
		NFTMsg      NFTMsg `json:"nft_msg"`      // 具体参考接口文档
		MTMsg       MTMsg  `json:"mt_msg"`       // 具体参考接口文档
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
