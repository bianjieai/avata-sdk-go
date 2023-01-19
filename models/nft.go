package models

const (
	CreateNFTClass   = "/v1beta1/nft/classes"                // 创建 NFT 类别
	QueryNFTClasses  = "/v1beta1/nft/classes"                // 查询 NFT 类别
	QueryNFTClass    = "/v1beta1/nft/classes/%s"             // 查询 NFT 类别详情
	TransferNFTClass = "/v1beta1/nft/class-transfers/%s/%s"  // 转让 NFT 类别
	MintNFT          = "/v1beta1/nft/nfts/%s"                // 发行 NFT
	TransferNFT      = "/v1beta1/nft/nft-transfers/%s/%s/%s" // 转让 NFT
	EditNFT          = "/v1beta1/nft/nfts/%s/%s/%s"          // 编辑 NFT
	BurnNFT          = "/v1beta1/nft/nfts/%s/%s/%s"          // 销毁 NFT
	BatchMintNFT     = "/v1beta1/nft/batch/nfts/%s"          // 批量发行 NFT
	BatchTransferNFT = "/v1beta1/nft/batch/nft-transfers/%s" // 批量转让 NFT
	BatchEditNFT     = "/v1beta1/nft/batch/nfts/%s"          // 批量编辑 NFT
	BatchBurnNFT     = "/v1beta1/nft/batch/nfts/%s"          // 批量销毁 NFT
	QueryNFTs        = "/v1beta1/nft/nfts"                   // 查询 NFT
	QueryNFT         = "/v1beta1/nft/nfts/%s/%s"             // 查询 NFT 详情
	QueryNFTHistory  = "/v1beta1/nft/nfts/%s/%s/history"     // 查询 NFT 历史记录
)

// CreateNFTClassReq 创建 NFT 类别：request
type CreateNFTClassReq struct {
	Name        string            `json:"name"`                  // NFT 类别名称
	ClassID     string            `json:"class_id,omitempty"`    // NFT 类别 ID，仅支持小写字母及数字，以字母开头
	Symbol      string            `json:"symbol,omitempty"`      // 标识
	Description string            `json:"description,omitempty"` // 描述
	Uri         string            `json:"uri,omitempty"`         // 链外数据链接
	UriHash     string            `json:"uri_hash,omitempty"`    // 链外数据 Hash
	Data        string            `json:"data,omitempty"`        // 自定义链上元数据
	Owner       string            `json:"owner"`                 // NFT 类别权属者地址，拥有在该 NFT 类别中发行 NFT 的权限和转让该 NFT 类别的权限。 支持任一 Avata 平台内合法链账户地址
	Tag         map[string]string `json:"tag,omitempty"`         // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
	OperationID string            `json:"operation_id"`          // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// QueryNFTClassesReq 查询 NFT 类别：request
type QueryNFTClassesReq struct {
	Offset    string `json:"offset,omitempty"`     // 游标，默认为 0
	Limit     string `json:"limit,omitempty"`      // 每页记录数，默认为 10，上限为 50
	ID        string `json:"id,omitempty"`         // NFT 类别 ID
	Name      string `json:"name,omitempty"`       // NFT 类别名称，支持模糊查询
	Owner     string `json:"owner,omitempty"`      // NFT 类别权属者地址
	TxHash    string `json:"tx_hash,omitempty"`    // 创建 NFT 类别的 Tx Hash
	StartDate string `json:"start_date,omitempty"` // NFT 类别创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   // NFT 类别创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    // 排序规则：DATE_ASC / DATE_DESC
}

// QueryNFTClassesRes 查询 NFT 类别：Response
type QueryNFTClassesRes struct {
	Data struct {
		Offset     int `json:"offset"`      // 游标
		Limit      int `json:"limit"`       // 每页记录数
		TotalCount int `json:"total_count"` // 总记录数
		Classes    []struct {
			ID        string `json:"id"`        // NFT 类别 ID
			Name      string `json:"name"`      // NFT 类别名称
			Symbol    string `json:"symbol"`    // NFT 类别标识
			NFTCount  int    `json:"nft_count"` // NFT 类别包含的 NFT 总量
			Uri       string `json:"uri"`       // 链外数据链接
			Owner     string `json:"owner"`     // NFT 类别权属者地址
			TxHash    string `json:"tx_hash"`   // 创建 NFT 类别的 Tx Hash
			TimeStamp string `json:"timestamp"` // 创建 NFT 类别的时间戳（UTC 时间）
		} `json:"classes"`
	} `json:"data"`
}

// QueryNFTClassRes 查询 NFT 类别详情：Response
type QueryNFTClassRes struct {
	Data struct {
		ID          string `json:"id"`          // NFT 类别 ID
		Name        string `json:"name"`        // NFT 类别名称
		Symbol      string `json:"symbol"`      // NFT 类别标识
		Description string `json:"description"` // NFT 类别描述
		NFTCount    int    `json:"nft_count"`   // NFT 类别包含的 NFT 总量
		Uri         string `json:"uri"`         // 链外数据链接
		UriHash     string `json:"uri_hash"`    // 链外数据 Hash
		Data        string `json:"data"`        // 自定义链上元数据
		Owner       string `json:"owner"`       // NFT 类别权属者地址
		TxHash      string `json:"tx_hash"`     // 创建 NFT 类别的 Tx Hash
		TimeStamp   string `json:"timestamp"`   // 创建 NFT 类别的时间戳（UTC 时间）
	} `json:"data"`
}

// TransferNFClassReq 转让 NFT 类别：request
type TransferNFClassReq struct {
	Recipient   string            `json:"recipient"`     // NFT 类别接收者地址，支持任一 Avata 平台内合法链账户地址
	Tag         map[string]string `json:"tag,omitempty"` // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
	OperationID string            `json:"operation_id"`  // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// MintNFTReq 发行 NFT：request
type MintNFTReq struct {
	Name        string            `json:"name"`                // NFT 名称
	Uri         string            `json:"uri,omitempty"`       // 链外数据链接
	UriHash     string            `json:"uri_hash,omitempty"`  // 链外数据 Hash
	Data        string            `json:"data,omitempty"`      // 自定义链上元数据
	Recipient   string            `json:"recipient,omitempty"` // NFT 接收者地址，支持任一文昌链合法链账户地址，默认为 NFT 类别的权属者地址
	Tag         map[string]string `json:"tag,omitempty"`       // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
	OperationID string            `json:"operation_id"`        // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// TransferNFTReq 转让 NFT ：request
type TransferNFTReq struct {
	Recipient   string            `json:"recipient"`     // NFT 接收者地址
	Tag         map[string]string `json:"tag,omitempty"` // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
	OperationID string            `json:"operation_id"`  // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// EditNFTReq 编辑 NFT ：request
type EditNFTReq struct {
	Name        string            `json:"name"`           // NFT 名称
	Uri         string            `json:"uri,omitempty"`  // 链外数据链接
	Data        string            `json:"data,omitempty"` // 自定义链上元数据
	Tag         map[string]string `json:"tag,omitempty"`  // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
	OperationID string            `json:"operation_id"`   // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// BurnNFTReq 销毁 NFT ：request
type BurnNFTReq struct {
	Tag         map[string]string `json:"tag,omitempty"` // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
	OperationID string            `json:"operation_id"`  // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// BatchMintNFTReq 批量发行nft：request
type (
	BatchMintNFTReq struct {
		Name        string            `json:"name"`               // NFT 名称
		Uri         string            `json:"uri,omitempty"`      // 链外数据链接
		UriHash     string            `json:"uri_hash,omitempty"` // 链外数据 Hash
		Data        string            `json:"data,omitempty"`     // 自定义链上元数据
		Recipients  []Recipients      `json:"recipients"`         // NFT 接收者地址和发行数量。以数组的方式进行组合，可以自定义多个组合，可面向多地址批量发行 NFT。
		Tag         map[string]string `json:"tag,omitempty"`      // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
		OperationID string            `json:"operation_id"`       // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	}

	Recipients struct {
		Amount    int    `json:"amount"`    // NFT 发行数量
		Recipient string `json:"recipient"` // NFT 接收者地址，支持任一文昌链合法链账户地址。
	}
)

// BatchTransferNFTReq 批量转让 NFT ：request
type (
	BatchTransferNFTReq struct {
		Data        []BatchTransferNFTData `json:"data"`
		Tag         map[string]string      `json:"tag,omitempty"` // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
		OperationID string                 `json:"operation_id"`  // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	}

	BatchTransferNFTData struct {
		NFTs      []BatchTransferNFTs `json:"nfts"`
		Recipient string              `json:"recipient"` // NFT 接收者地址
	}
	BatchTransferNFTs struct {
		ClassID string `json:"class_id"` // NFT 类别 ID
		NFTID   string `json:"nft_id"`   // NFT ID
	}
)

// BatchEditNFTReq 批量编辑 NFT ：request
type (
	BatchEditNFTReq struct {
		NFTs        []BatchEditNfts   `json:"nfts"`
		Tag         map[string]string `json:"tag,omitempty"` // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
		OperationID string            `json:"operation_id"`  // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	}
	BatchEditNfts struct {
		ClassID string `json:"class_id"`       // NFT 类别 ID
		NFTID   string `json:"nft_id"`         // NFT ID
		Name    string `json:"name"`           // NFT 名称
		Uri     string `json:"uri,omitempty"`  // 链外数据链接
		Data    string `json:"data,omitempty"` // 自定义链上元数据
	}
)

// BatchBurnNFTReq 批量销毁 NFT ：request
type (
	BatchBurnNFTReq struct {
		NFTs        []NFTs            `json:"nfts"`
		Tag         map[string]string `json:"tag,omitempty"` // 交易标签， 自定义 key：支持大小写英文字母和汉字和数字，长度 6-12 位，自定义 value：长度限制在 64 位字符，支持大小写字母和数字
		OperationID string            `json:"operation_id"`  // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	}
	NFTs struct {
		ClassID string `json:"class_id"` // NFT 类别 ID
		NFTID   string `json:"nft_id"`   // NFT ID
	}
)

// QueryNFTsReq 查询 NFT ：request
type QueryNFTsReq struct {
	Offset    string `json:"offset,omitempty"`
	Limit     string `json:"limit,omitempty"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ClassID   string `json:"class_id,omitempty"`
	Owner     string `json:"owner,omitempty"`
	TxHash    string `json:"tx_hash,omitempty"`
	Status    string `json:"status,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
}

// QueryNFTsRes 查询 NFT ：Response
type QueryNFTsRes struct {
	Data struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		TotalCount int `json:"total_count"`
		NFTs       []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			ClassID     string `json:"class_id"`
			ClassName   string `json:"class_name"`
			ClassSymbol string `json:"class_symbol"`
			NFTCount    int    `json:"nft_count"`
			Uri         string `json:"uri"`
			Owner       string `json:"owner"`
			Status      string `json:"status"`
			TxHash      string `json:"tx_hash"`
			TimeStamp   string `json:"timestamp"`
		} `json:"nfts"`
	} `json:"data"`
}

// QueryNFTRes 查询 NFT 详情：Response
type QueryNFTRes struct {
	Data struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		ClassID     string `json:"class_id"`
		ClassName   string `json:"class_name"`
		ClassSymbol string `json:"class_symbol"`
		Uri         string `json:"uri"`
		UriHash     string `json:"uri_hash"`
		Data        string `json:"data"`
		Owner       string `json:"owner"`
		Status      string `json:"status"`
		TxHash      string `json:"tx_hash"`
		TimeStamp   string `json:"timestamp"`
	} `json:"data"`
}

// QueryNFTHistoryReq 查询 NFT 操作记录：request
type QueryNFTHistoryReq struct {
	Offset    string `json:"offset,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Signer    string `json:"signer,omitempty"`
	TxHash    string `json:"tx_hash,omitempty"`
	Operation string `json:"operation,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
}

// QueryNFTHistoryRes 查询 NFT 操作记录：Response
type QueryNFTHistoryRes struct {
	Data struct {
		Offset           int `json:"offset"`
		Limit            int `json:"limit"`
		TotalCount       int `json:"total_count"`
		OperationRecords []struct {
			TxHash    string `json:"tx_hash"`
			Operation string `json:"operation"`
			Signer    string `json:"signer"`
			Recipient string `json:"recipient"`
			TimeStamp string `json:"timestamp"`
		} `json:"operation_records"`
	} `json:"data"`
}
