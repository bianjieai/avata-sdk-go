package models

const (
	CreateMTClass   = "/v3/native/mt/classes"               // 创建 MT 类别
	QueryMTClasses  = "/v3/native/mt/classes"               // 查询 MT 类别
	QueryMTClass    = "/v3/native/mt/classes/%s"            // 查询 MT 类别详情
	TransferMTClass = "/v3/native/mt/class-transfers/%s/%s" // 转让 MT 类别
	IssueMT         = "/v3/native/mt/mt-issues/%s"          // 发行 MT
	MintMT          = "/v3/native/mt/mt-mints/%s/%s"        // 增发 MT
	TransferMT      = "/v3/native/mt/mt-transfers/%s/%s/%s" // 转让 MT
	EditMT          = "/v3/native/mt/mts/%s/%s/%s"          // 编辑 MT
	BurnMT          = "/v3/native/mt/mts/%s/%s/%s"          // 销毁 MT
	QueryMTs        = "/v3/native/mt/mts"                   // 查询 MT
	QueryMT         = "/v3/native/mt/mts/%s/%s"             // 查询 MT 详情
	QueryMTHistory  = "/v3/native/mt/mts/%s/%s/history"     // 查询 MT 操作记录
	QueryMTBalance  = "/v3/native/mt/mts/%s/%s/balances"    // 查询 MT 余额
)

// CreateMTClassReq 创建 MT 类别请求参数
type CreateMTClassReq struct {
	Name        string `json:"name"`           // MT 类别名称
	Owner       string `json:"owner"`          // MT 类别权属者地址，支持任一 Avata 平台内合法链账户地址
	Data        string `json:"data,omitempty"` // 自定义链上元数据
	OperationId string `json:"operation_id"`   // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// QueryMTClassesReq 查询 MT 类别请求参数
type QueryMTClassesReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 游标，默认为 0
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // MT 类别 ID
	Name       string `json:"name,omitempty"`        // MT 类别名称，支持模糊查询
	Owner      string `json:"owner,omitempty"`       // MT 类别权属者地址
	TxHash     string `json:"tx_hash,omitempty"`     // 创建 MT 类别的 Tx Hash
	StartDate  string `json:"start_date,omitempty"`  // MT 类别创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // MT 类别创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"` //是否查询数据的总数量 0：不查询总数（默认） 1：查询总数
}

// QueryMTClassesRes 查询 MT 类别返回值
type QueryMTClassesRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"` // 游标
		NextPageKey string `json:"next_page_key"` // 游标
		Limit       int    `json:"limit"`         // 每页记录数
		TotalCount  int    `json:"total_count"`   // 总记录数
		Classes     []struct {
			Id        string `json:"id"`        // MT 类别 ID
			Name      string `json:"name"`      // MT 类别名称
			MtCount   int    `json:"mt_count"`  // MT 类别包含的 MT 总量(AVATA 平台内)
			Owner     string `json:"owner"`     // MT 类别权属者地址
			TxHash    string `json:"tx_hash"`   // 创建 MT 类别的 Tx Hash
			Timestamp string `json:"timestamp"` // 创建 MT 类别的时间戳（UTC 时间）
		} `json:"classes"` // 类别列表
	} `json:"data"`
}

// QueryMTClassRes 查询 MT 类别详情返回值
type QueryMTClassRes struct {
	Data struct {
		Id        string `json:"id"`             // MT 类别 ID
		Name      string `json:"name"`           // MT 类别名称
		Data      string `json:"data,omitempty"` // 自定义链上元数据
		MtCount   int    `json:"mt_count"`       // MT 类别包含的 MT 总量(AVATA 平台内)
		Owner     string `json:"owner"`          // MT 类别权属者地址
		TxHash    string `json:"tx_hash"`        // 创建 MT 类别的 Tx Hash
		Timestamp string `json:"timestamp"`      // 创建 MT 类别的时间戳（UTC 时间）
	} `json:"data,omitempty"`
}

// TransferMTClassReq 转让 MT 类别请求参数
type TransferMTClassReq struct {
	Recipient   string `json:"recipient"`    // MT 类别接收者地址，支持任一 Avata 内合法链账户地址
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// IssueMTReq 发行 MT 请求参数
type IssueMTReq struct {
	Data        string `json:"data,omitempty"`      // 自定义链上元数据
	Amount      int    `json:"amount,omitempty"`    // MT 数量，不填写数量时，默认发行数量为 1
	Recipient   string `json:"recipient,omitempty"` // MT 接收者地址，支持任一文昌链合法链账户地址，默认为 MT 类别的权属者地址
	OperationID string `json:"operation_id"`        // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// MintMTReq 增发 MT 请求参数
type MintMTReq struct {
	Amount      int    `json:"amount,omitempty"` // MT 数量
	Recipient   string `json:"recipient"`        // MT 接收者地址
	OperationID string `json:"operation_id"`     // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// TransferMTReq 转让 MT 请求参数
type TransferMTReq struct {
	Recipient   string `json:"recipient"`        // MT 接收者地址
	Amount      int    `json:"amount,omitempty"` // 转移的数量（默认为 1 ）
	OperationID string `json:"operation_id"`     // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// EditMTReq 编辑 MT 请求参数
type EditMTReq struct {
	Data        string `json:"data"`         // 自定义链上元数据
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// BurnMTReq 销毁 MT 请求参数
type BurnMTReq struct {
	Amount      int    `json:"amount,omitempty"` // 销毁的数量
	OperationId string `json:"operation_id"`     // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// QueryMTsReq 查询 MT 请求参数
type QueryMTsReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 游标，默认为 0
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // MT ID
	ClassID    string `json:"class_id,omitempty"`    // MT 类别 ID
	Issuer     string `json:"issuer,omitempty"`      // MT 发行者地址
	TxHash     string `json:"tx_hash,omitempty"`     // 创建 MT 的 TX Hash
	StartDate  string `json:"start_date,omitempty"`  // MT 创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // MT 创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"` //是否查询数据的总数量 0：不查询总数（默认） 1：查询总数
}

// QueryMTsRes 查询 MT 返回值
type QueryMTsRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"` // 游标
		NextPageKey string `json:"next_page_key"` // 游标
		Limit       int    `json:"limit"`         // 每页记录数
		TotalCount  int    `json:"total_count"`   // 总记录数
		Mts         []struct {
			Id         string `json:"id"`                    // MT ID
			ClassId    string `json:"class_id"`              // MT 类别 ID
			ClassName  string `json:"class_name"`            // MT 类别名称
			Issuer     string `json:"issuer"`                // 首次发行该 MT 的链账户地址
			OwnerCount int    `json:"owner_count,omitempty"` // MT 拥有者数量(AVATA 平台内)
			Timestamp  string `json:"timestamp"`             // MT 首次发行时间戳（UTC 时间）
		} `json:"mts,omitempty"`
	} `json:"data,omitempty"`
}

// QueryMTRes 查询 MT 详情返回值
type QueryMTRes struct {
	Data struct {
		Id         string `json:"id"`          // MT ID
		ClassId    string `json:"class_id"`    // MT 类别 ID
		ClassName  string `json:"class_name"`  // MT 类别名称
		Data       string `json:"data"`        // 自定义链上元数据
		OwnerCount int    `json:"owner_count"` // MT 拥有者数量(AVATA 平台内)
		IssueData  struct {
			Issuer    string `json:"issuer"`    // 首次发行该 MT 的链账户地址
			Timestamp string `json:"timestamp"` // 首次发行该 MT 的时间戳
			Count     int    `json:"count"`     // 首次发行该 MT 的数量
			TxHash    string `json:"tx_hash"`   // 首次发行该 MT 的交易哈希
		} `json:"issue_data"`
		MtCount   int `json:"mt_count"`   // MT 流通总量(全链)
		MintTimes int `json:"mint_times"` // MT 发行次数(AVATA 平台内累计发行次数(包括首次发行和增发))
	} `json:"data"`
}

// QueryMTHistoryReq 查询 MT 操作记录请求参数
type QueryMTHistoryReq struct {
	PageKey    string `json:"offset,omitempty"`     // 游标，默认为 0
	Limit      string `json:"limit,omitempty"`      // 每页记录数，默认为 10，上限为 50
	Signer     string `json:"signer,omitempty"`     // Tx 签名者地址
	TxHash     string `json:"tx_hash,omitempty"`    // MT 操作 Tx Hash
	Operation  string `json:"operation,omitempty"`  // 操作类型： issue(首发MT) / mint(增发MT) / edit(编辑MT) / transfer(转让MT) / burn(销毁MT)
	StartDate  string `json:"start_date,omitempty"` // MT 操作日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`   // MT 操作日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`    // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"`
}

// QueryMTHistoryRes 查询 MT 操作记录返回值
type QueryMTHistoryRes struct {
	Data struct {
		PrevPageKey      string `json:"prev_page_key"` // 游标
		NextPageKey      string `json:"next_page_key"` // 游标
		Limit            int    `json:"limit"`         // 每页记录数
		TotalCount       int    `json:"total_count"`   // 总记录数
		OperationRecords []struct {
			TxHash    string `json:"tx_hash"`             // MT 操作的 Tx Hash
			Operation int    `json:"operation"`           // MT 操作类型；Enum: "issue" "mint" "edit" "transfer" "burn"
			Signer    string `json:"signer"`              // Tx 签名者地址
			Recipient string `json:"recipient,omitempty"` // MT 接收者地址
			Amount    int    `json:"amount,omitempty"`    // MT 操作数量
			Timestamp string `json:"timestamp"`           // MT 操作时间戳（UTC 时间）
		} `json:"operation_records"`
	} `json:"data,omitempty"`
}

// QueryMTBalanceReq 查询 MT 余额请求参数
type QueryMTBalanceReq struct {
	PageKey    string `json:"offset,omitempty"`      // 游标，默认为 0
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // MT ID
	CountTotal string `json:"count_total,omitempty"` // 是否查询数据的总数量 0：不查询总数（默认） 1：查询总数
}

// QueryMTBalanceRes 查询 MT 余额返回值
type QueryMTBalanceRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"` // 游标
		NextPageKey string `json:"next_page_key"` // 游标
		Limit       int    `json:"limit"`         // 每页记录数
		TotalCount  int    `json:"total_count"`   // 总记录数
		Mts         []struct {
			Id     string `json:"id,omitempty"`     // MT ID
			Amount int    `json:"amount,omitempty"` // MT 数量
		} `json:"mts"`
	} `json:"data"`
}
