package models

const (
	CreateMTClass   = "/v1beta1/mt/classes"               // 创建 MT 类别
	GetMTClasses    = "/v1beta1/mt/classes"               // 查询 MT 类别
	GetMTClass      = "/v1beta1/mt/classes/%s"            // 查询 MT 类别详情
	TransferMTClass = "/v1beta1/mt/class-transfers/%s/%s" // 转让 MT 类别
	IssueMT         = "/v1beta1/mt/mt-issues/%s"          // 发行 MT
	MintMT          = "/v1beta1/mt/mt-mints/%s/%s"        // 增发 MT
	TransferMT      = "/v1beta1/mt/mt-transfers/%s/%s/%s" // 转让 MT
	EditMT          = "/v1beta1/mt/mts/%s/%s/%s"          // 编辑 MT
	BurnMT          = "/v1beta1/mt/mts/%s/%s/%s"          // 销毁 MT
	GetMTs          = "/v1beta1/mt/mts"                   // 查询 MT
	GetMT           = "/v1beta1/mt/mts/%s/%s"             // 查询 MT 详情
	GetMTHistory    = "/v1beta1/mt/mts/%s/%s/history"     // 查询 MT 操作记录
	GetMTBalance    = "/v1beta1/mt/mts/%s/%s/balances"    // 查询 MT 余额
)

// CreateMTClassReq 创建 MT 类别请求参数
type CreateMTClassReq struct {
	Name        string            `json:"name"`
	Owner       string            `json:"owner"`
	Data        string            `json:"data,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationId string            `json:"operation_id"`
}

// GetMTClassesReq 查询 MT 类别请求参数
type GetMTClassesReq struct {
	Offset    string `json:"offset,omitempty"`
	Limit     string `json:"limit,omitempty"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Owner     string `json:"owner,omitempty"`
	TxHash    string `json:"tx_hash,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
}

// GetMTClassesRes 查询 MT 类别返回值
type GetMTClassesRes struct {
	BaseRes
	Data struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		TotalCount int `json:"total_count"`
		Classes    []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			MtCount   int    `json:"mt_count"`
			Owner     string `json:"owner"`
			TxHash    string `json:"tx_hash"`
			Timestamp string `json:"timestamp"`
		} `json:"classes"`
	} `json:"data"`
}

// GetMTClassRes 查询 MT 类别详情返回值
type GetMTClassRes struct {
	BaseRes
	Data struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		MtCount   int    `json:"mt_count"`
		Data      string `json:"data"`
		Owner     string `json:"owner"`
		TxHash    string `json:"tx_hash"`
		Timestamp string `json:"timestamp"`
	} `json:"data"`
}

// TransferMTClassReq 转让 MT 类别请求参数
type TransferMTClassReq struct {
	Recipient   string            `json:"recipient"`
	OperationId string            `json:"operation_id"`
	Tag         map[string]string `json:"tag,omitempty"`
}

// IssueMTReq 发行 MT 请求参数
type IssueMTReq struct {
	Data        string            `json:"data,omitempty"`
	Amount      int               `json:"amount,omitempty"`
	Recipient   string            `json:"recipient,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationId string            `json:"operation_id"`
}

// MintMTReq 增发 MT 请求参数
type MintMTReq struct {
	Amount      int               `json:"amount,omitempty"`
	Recipient   string            `json:"recipient,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationId string            `json:"operation_id"`
}

// TransferMTReq 转让 MT 请求参数
type TransferMTReq struct {
	Amount      int               `json:"amount,omitempty"`
	Recipient   string            `json:"recipient"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationId string            `json:"operation_id"`
}

// EditMTReq 编辑 MT 请求参数
type EditMTReq struct {
	Data        string            `json:"data"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationId string            `json:"operation_id"`
}

// BurnMTReq 销毁 MT 请求参数
type BurnMTReq struct {
	Amount      int               `json:"amount,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationId string            `json:"operation_id"`
}

// GetMTsReq 查询 MT 请求参数
type GetMTsReq struct {
	Offset    string `json:"offset,omitempty"`
	Limit     string `json:"limit,omitempty"`
	ID        string `json:"id,omitempty"`
	ClassID   string `json:"class_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Owner     string `json:"owner,omitempty"`
	TxHash    string `json:"tx_hash,omitempty"`
	Status    string `json:"status,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
}

// GetMTsRes 查询 MT 返回值
type GetMTsRes struct {
	BaseRes
	Data struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		TotalCount int `json:"total_count"`
		Mts        []struct {
			Id         string `json:"id"`
			ClassId    string `json:"class_id"`
			ClassName  string `json:"class_name"`
			Issuer     string `json:"issuer"`
			OwnerCount int    `json:"owner_count"`
			Timestamp  string `json:"timestamp"`
		} `json:"mts"`
	} `json:"data"`
}

// GetMTRes 查询 MT 详情返回值
type GetMTRes struct {
	BaseRes
	Data struct {
		Id         string `json:"id"`
		ClassId    string `json:"class_id"`
		ClassName  string `json:"class_name"`
		Data       string `json:"data"`
		OwnerCount int    `json:"owner_count"`
		IssueData  struct {
			Issuer    string `json:"issuer"`
			Timestamp string `json:"timestamp"`
			Count     int    `json:"count"`
			TxHash    string `json:"tx_hash"`
		} `json:"issue_data"`
		MtCount   int `json:"mt_count"`
		MintTimes int `json:"mint_times"`
	} `json:"data"`
}

// GetMTHistoryReq 查询 MT 操作记录请求参数
type GetMTHistoryReq struct {
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

// GetMTHistoryRes 查询 MT 操作记录返回值
type GetMTHistoryRes struct {
	BaseRes
	Data struct {
		Offset           int `json:"offset"`
		Limit            int `json:"limit"`
		TotalCount       int `json:"total_count"`
		OperationRecords []struct {
			TxHash    string `json:"tx_hash"`
			Operation string `json:"operation"`
			Signer    string `json:"signer"`
			Recipient string `json:"recipient"`
			Amount    int    `json:"amount"`
			Timestamp string `json:"timestamp"`
		} `json:"operation_records"`
	} `json:"data"`
}

// GetMTBalanceReq 查询 MT 余额请求参数
type GetMTBalanceReq struct {
	Offset string `json:"offset,omitempty"`
	Limit  string `json:"limit,omitempty"`
	ID     string `json:"id,omitempty"`
}

// GetMTBalanceRes 查询 MT 余额返回值
type GetMTBalanceRes struct {
	BaseRes
	Data struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		TotalCount int `json:"total_count"`
		Mts        []struct {
			Id     string `json:"id"`
			Amount int    `json:"amount"`
		} `json:"mts"`
	} `json:"data"`
}
