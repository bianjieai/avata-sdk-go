package models

const (
	CreateNFTClass         = "/v3/evm/nft/classes"                   // EVM 模块创建 NFT 类别
	QueryNFTClasses        = "/v3/evm/nft/classes"                   // EVM 模块查询 NFT 类别
	QueryNFTClass          = "/v3/evm/nft/classes/%s"                // EVM 模块查询 NFT 类别详情
	TransferNFTClass       = "/v3/evm/nft/class-transfers/%s/%s"     // EVM 模块转让 NFT 类别
	MintNFT                = "/v3/evm/nft/nfts/%s"                   // EVM 模块发行 NFT
	TransferNFT            = "/v3/evm/nft/nft-transfers/%s/%s/%s"    // EVM 模块转让 NFT
	EditNFT                = "/v3/evm/nft/nfts/%s/%s/%s"             // EVM 模块编辑 NFT
	BurnNFT                = "/v3/evm/nft/nfts/%s/%s/%s"             // EVM 模块销毁 NFT
	QueryNFTs              = "/v3/evm/nft/nfts"                      // EVM 模块查询 NFT
	QueryNFT               = "/v3/evm/nft/nfts/%s/%s"                // EVM 模块查询 NFT 详情
	QueryNFTHistory        = "/v3/evm/nft/nfts/%s/%s/history"        // EVM 模块查询 NFT 操作记录
	CreateNativeNFTClass   = "/v3/native/nft/classes"                // 原生模块创建 NTF 类别
	QueryNativeNFTClasses  = "/v3/native/nft/classes"                // 原生模块查询 NTF 类别
	QueryNativeNFTClass    = "/v3/native/nft/classes/%s"             // 原生模块查询 NFT 类别详情
	TransferNativeNFTClass = "/v3/native/nft/class-transfers/%s/%s"  // 原生模块转让 NFT 类别
	MintNativeNFT          = "/v3/native/nft/nfts/%s"                // 原生模块发行 NFT
	TransferNativeNFT      = "/v3/native/nft/nft-transfers/%s/%s/%s" // 原生模块转让 NFT
	EditNativeNFT          = "/v3/native/nft/nfts/%s/%s/%s"          // 原生模块编辑 NFT
	BurnNativeNFT          = "/v3/native/nft/nfts/%s/%s/%s"          // 原生模块销毁 NFT
	QueryNativeNFTs        = "/v3/native/nft/nfts"                   // 原生模块查询 NFT
	QueryNativeNFT         = "/v3/native/nft/nfts/%s/%s"             // 原生模块查询 NFT 详情
	QueryNativeNFTHistory  = "/v3/native/nft/nfts/%s/%s/history"     // 原生模块查询 NFT 操作记录
)

// CreateNFTClassReq 创建 EVM 模块 NFT 类别请求参数
type CreateNFTClassReq struct {
	Name                 string `json:"name"`                              // NFT 类别名称
	Uri                  string `json:"uri,omitempty"`                     // 链外数据链接
	Symbol               string `json:"symbol"`                            // 标识
	Owner                string `json:"owner"`                             // NFT 类别权属者地址，拥有在该 NFT 类别中发行 NFT 的权限和转让该 NFT 类别的权限。 支持任一 Avata 平台内合法链账户地址
	EditableByOwner      int    `json:"editable_by_owner,omitempty"`       // NFT 类别权限的控制功能， 此类别下某一 NFT 的持有者可以编辑该 NFT ：1可编辑 0不可编辑
	EditableByClassOwner int    `json:"editable_by_class_owner,omitempty"` // NFT 类别权限的控制功能， 此 NFT 类别的权属者可以编辑这个类别下所有的 NFT ：1可编辑 0不可编辑
	OperationID          string `json:"operation_id"`                      // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// CreateNativeNFTClassReq 创建原生模块 NFT 类别请求参数
type CreateNativeNFTClassReq struct {
	Name            string `json:"name"`                        // NFT 类别名称
	ClasID          string `json:"clas_id,omitempty"`           // NFT类别ID
	Symbol          string `json:"symbol,omitempty"`            // 标识
	Description     string `json:"description,omitempty"`       // 描述
	Uri             string `json:"uri,omitempty"`               // 链外数据链接
	UriHash         string `json:"uri_hash,omitempty"`          // 链外数据 Hash
	Data            string `json:"data,omitempty"`              // 自定义链上元数据
	EditableByOwner int    `json:"editable_by_owner,omitempty"` // NFT 拥有者是否可编辑 NFT,1可编辑，0不可编辑，默认1
	Owner           string `json:"owner"`                       // NFT 类别权属者地址，拥有在该 NFT 类别中发行 NFT 的权限和转让该 NFT 类别的权限。 支持任一 Avata 平台内合法链账户地址
	OperationID     string `json:"operation_id"`                // 操作 ID，保证幂等性；由接入方生成的唯一的、大小写敏感、不超过 64 个 ASCII 字符的字符串
}

// QueryNFTClassesReq 查询 EVM 模块 NFT 类别请求参数
type QueryNFTClassesReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 分页数据的 Key，AVATA 会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // NFT 类别 ID
	Name       string `json:"name,omitempty"`        // NFT 类别名称，支持模糊查询
	Owner      string `json:"owner,omitempty"`       // NFT 类别权属者地址
	TxHash     string `json:"tx_hash,omitempty"`     // 创建 NFT 类别的 Tx Hash
	StartDate  string `json:"start_date,omitempty"`  // NFT 类别创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // NFT 类别创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"` // 是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryNativeNFTClassesReq 查询原生模块 NFT 类别请求参数
type QueryNativeNFTClassesReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 分页数据的 Key，AVATA 会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // NFT 类别 ID
	Name       string `json:"name,omitempty"`        // NFT 类别名称，支持模糊查询
	Owner      string `json:"owner,omitempty"`       // NFT 类别权属者地址
	TxHash     string `json:"tx_hash,omitempty"`     // 创建 NFT 类别的 Tx Hash
	StartDate  string `json:"start_date,omitempty"`  // NFT 类别创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // NFT 类别创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"` // 是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryNFTClassesRes 查询 EVM 模块 NFT 类别返回值
type QueryNFTClassesRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"` // 上一页数据的 Key，AVATA 会根据该值进行上一页数据的查询
		NextPageKey string `json:"next_page_key"` // 下一页数据的 Key，AVATA 会根据该值进行下一页数据的查询
		Limit       int    `json:"limit"`         // 每页记录数
		TotalCount  int    `json:"total_count"`   // 总记录数
		Classes     []struct {
			ID        string `json:"id"`        // NFT 类别 ID
			Name      string `json:"name"`      // NFT 类别名称
			Uri       string `json:"uri"`       // 链外数据链接
			Symbol    string `json:"symbol"`    // NFT 类别标识
			Owner     string `json:"owner"`     // NFT 类别权属者地址
			TxHash    string `json:"tx_hash"`   // 创建 NFT 类别的 Tx Hash
			TimeStamp string `json:"timestamp"` // 创建 NFT 类别的时间戳（UTC 时间）
		} `json:"classes"`
	} `json:"data"`
}

// QueryNativeNFTClassesRes 查询原生模块 NFT 类别返回值
type QueryNativeNFTClassesRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"` // 上一页数据的 Key，AVATA 会根据该值进行上一页数据的查询
		NextPageKey string `json:"next_page_key"` // 下一页数据的 Key，AVATA 会根据该值进行下一页数据的查询
		Limit       int    `json:"limit"`         // 每页记录数
		TotalCount  int    `json:"total_count"`   // 总记录数
		Classes     []struct {
			ID        string `json:"id"`               // NFT 类别 ID
			Name      string `json:"name"`             // NFT 类别名称
			Symbol    string `json:"symbol,omitempty"` // NFT 类别标识
			Uri       string `json:"uri"`              // 链外数据链接
			Owner     string `json:"owner"`            // NFT 类别权属者地址
			TxHash    string `json:"tx_hash"`          // 创建 NFT 类别的 Tx Hash
			TimeStamp string `json:"timestamp"`        // 创建 NFT 类别的时间戳（UTC 时间）
		} `json:"classes"`
	} `json:"data"`
}

// QueryNFTClassRes 查询 EVM 模块 NFT 类别详情返回值
type QueryNFTClassRes struct {
	Data struct {
		ID                   string `json:"id"`                      // NFT 类别 ID
		Name                 string `json:"name"`                    // NFT 类别名称
		Uri                  string `json:"uri,omitempty"`           // 链外数据链接
		Symbol               string `json:"symbol"`                  // NFT 类别标识
		NftCount             int    `json:"nft_count"`               // NFT 类别包含的 NFT 总量
		Owner                string `json:"owner"`                   // NFT 类别权属者地址
		EditableByOwner      int    `json:"editable_by_owner"`       // NFT 类别权限的控制功能， 此类别下某一 NFT 的持有者可以编辑该 NFT ：1可编辑 0不可编辑
		EditableByClassOwner int    `json:"editable_by_class_owner"` // NFT 类别权限的控制功能， 此 NFT 类别的权属者可以编辑这个类别下所有的 NFT 1：可编辑 0：不可编辑（默认）
		TxHash               string `json:"tx_hash"`                 // 创建 NFT 类别的 Tx Hash
		TimeStamp            string `json:"timestamp"`               // 创建 NFT 类别的时间戳（UTC 时间）
	} `json:"data"`
}

// QueryNativeNFTClassRes 查询原生模块 NFT 类别详情返回值
type QueryNativeNFTClassRes struct {
	Data struct {
		ID              string `json:"id,omitempty"`                // NFT 类别 ID
		Name            string `json:"name,omitempty"`              // NFT 类别名称
		Symbol          string `json:"symbol,omitempty"`            // NFT 类别标识
		Description     string `json:"description,omitempty"`       // NFT 类别描述
		NftCount        int    `json:"nft_count,omitempty"`         // NFT 类别包含的 NFT 总量
		Uri             string `json:"uri,omitempty"`               // 链外数据链接
		UriHash         string `json:"uri_hash,omitempty"`          // 链外数据 Hash
		Data            string `json:"data,omitempty"`              // 自定义链上元数据
		Owner           string `json:"owner,omitempty"`             // NFT 类别权属者地址
		TxHash          string `json:"tx_hash,omitempty"`           // 创建 NFT 类别的 Tx Hash
		TimeStamp       string `json:"timestamp,omitempty"`         // 创建 NFT 类别的时间戳（UTC 时间）
		EditableByOwner int    `json:"editable_by_owner,omitempty"` // NFT 类别权限的控制功能， 此类别下某一 NFT 的持有者可以编辑该 NFT ：1可编辑 0不可编辑
	} `json:"data"`
}

// TransferNFClassReq 转让 EVM 模块 NFT 类别请求参数
type TransferNFClassReq struct {
	Recipient   string `json:"recipient"`    // NFT 合约接收者地址
	OperationID string `json:"operation_id"` // 保证幂等性，避免重复请求
}

// TransferNativeNFClassReq 转让原生模块 NFT 类别请求参数
type TransferNativeNFClassReq struct {
	Recipient   string `json:"recipient"`    // NFT 合约接收者地址
	OperationID string `json:"operation_id"` // 保证幂等性，避免重复请求
}

// MintNFTReq 发行 EVM 模块 NFT 请求参数
type MintNFTReq struct {
	Uri         string `json:"uri"`                // 链外数据链接
	UriHash     string `json:"uri_hash,omitempty"` // 链外数据 Hash
	Recipient   string `json:"recipient"`          // NFT 接收者地址，支持任一文昌链合法链账户地址，默认为 NFT 类别的权属者地址
	OperationID string `json:"operation_id"`       // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// MintNativeNFTReq 发行原生模块 NFT 请求参数
type MintNativeNFTReq struct {
	Name        string `json:"name,omitempty"`      // NFT 名称
	Uri         string `json:"uri,omitempty"`       // 链外数据链接
	UriHash     string `json:"uri_hash,omitempty"`  // 链外数据 Hash
	Data        string `json:"data,omitempty"`      // 自定义链上元数据
	Recipient   string `json:"recipient,omitempty"` // NFT 接收者地址，支持任一文昌链合法链账户地址，默认为 NFT 类别的权属者地址
	OperationID string `json:"operation_id"`        // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// TransferNFTReq 转让 EVM 模块 NFT 请求参数
type TransferNFTReq struct {
	Recipient   string `json:"recipient"`    // NFT 接收者地址
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// TransferNativeNFTReq 转让原生模块 NFT 请求参数
type TransferNativeNFTReq struct {
	Recipient   string `json:"recipient"`    // NFT 接收者地址
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// EditNFTReq 编辑 EVM 模块 NFT 请求参数
type EditNFTReq struct {
	Uri         string `json:"uri"`                // 链外数据链接
	UriHash     string `json:"uri_hash,omitempty"` // 链外数据 Hash
	OperationID string `json:"operation_id"`       // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// EditNativeNFTReq 编辑原生模块 NFT 请求参数
type EditNativeNFTReq struct {
	Name        string `json:"name,omitempty"`     // NFT 名称
	Uri         string `json:"uri,omitempty"`      // 链外数据链接
	Data        string `json:"data,omitempty"`     // 自定义链上元数据
	UriHash     string `json:"uri_hash,omitempty"` // 链外数据 Hash
	OperationID string `json:"operation_id"`       // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// BurnNFTReq 销毁 EVM 模块 NFT 请求参数
type BurnNFTReq struct {
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// BurnNativeNFTReq 销毁原生模块 NFT 请求参数
type BurnNativeNFTReq struct {
	OperationID string `json:"operation_id"` // 操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// QueryNFTsReq 查询 EVM 模块 NFT 请求参数
type QueryNFTsReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 分页数据的 Key，AVATA 会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // NFT ID
	ClassID    string `json:"class_id,omitempty"`    // NFT 类别 ID
	Owner      string `json:"owner,omitempty"`       // NFT 持有者地址
	TxHash     string `json:"tx_hash,omitempty"`     // 创建 NFT 的 Tx Hash
	Status     string `json:"status,omitempty"`      // NFT 状态 1：active 2：burned
	StartDate  string `json:"start_date,omitempty"`  // NFT 创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // NFT 创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"` // 是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryNativeNFTsReq 查询原生模块 NFT 请求参数
type QueryNativeNFTsReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 分页数据的 Key， AVATA 会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	ID         string `json:"id,omitempty"`          // NFT ID
	ClassID    string `json:"class_id,omitempty"`    // NFT 类别 ID
	Owner      string `json:"owner,omitempty"`       // NFT 持有者地址
	TxHash     string `json:"tx_hash,omitempty"`     // 创建 NFT 的 Tx Hash
	StartDate  string `json:"start_date,omitempty"`  // NFT 创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // NFT 创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	Status     string `json:"status,omitempty"`      // NFT 状态：active / burned，默认为 active
	Name       string `json:"name,omitempty"`        // NFT 名称，支持模糊查询
	CountTotal string `json:"count_total,omitempty"` // 是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryNFTsRes 查询 EVM 模块 NFT 返回值
type QueryNFTsRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"`         // 上一页数据的 Key，AVATA 会根据该值进行上一页数据的查询
		NextPageKey string `json:"next_page_key"`         // 下一页数据的 Key，AVATA 会根据该值进行下一页数据的查询
		Limit       int    `json:"limit"`                 // 每页记录数
		TotalCount  int    `json:"total_count,omitempty"` // 总记录数
		NFTs        []struct {
			ID          int    `json:"id"`                 // NFT ID
			ClassID     string `json:"class_id"`           // NFT 类别 ID
			ClassName   string `json:"class_name"`         // NFT 类别名称
			ClassSymbol string `json:"class_symbol"`       // NFT 类别标识
			Uri         string `json:"uri"`                // 链外数据链接
			UriHash     string `json:"uri_hash,omitempty"` // 链外数据 Hash
			Owner       string `json:"owner"`              // NFT 持有者地址
			Status      int    `json:"status"`             // NFT 状态 1：active 2：burned
			TxHash      string `json:"tx_hash"`            // NFT 发行 Tx Hash
			TimeStamp   string `json:"timestamp"`          // NFT 发行时间戳（UTC 时间）
		} `json:"nfts,omitempty"` // NFT 列表
	} `json:"data"`
}

// QueryNativeNFTsRes 查询原生模块 NFT 返回值
type QueryNativeNFTsRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"`         // 上一页数据的 Key，AVATA 会根据该值进行上一页数据的查询
		NextPageKey string `json:"next_page_key"`         // 下一页数据的 Key，AVATA 会根据该值进行下一页数据的查询
		Limit       int    `json:"limit"`                 // 每页记录数
		TotalCount  int    `json:"total_count,omitempty"` // 总记录数
		NFTs        []struct {
			ID          string `json:"id"`                     // NFT ID
			Name        string `json:"name,omitempty"`         // NFT 名称
			ClassID     string `json:"class_id"`               // NFT 类别 ID
			ClassName   string `json:"class_name"`             // NFT 类别名称
			ClassSymbol string `json:"class_symbol,omitempty"` // NFT 类别标识
			Uri         string `json:"uri,omitempty"`          // 链外数据链接
			Owner       string `json:"owner"`                  // NFT 持有者地址
			Status      int    `json:"status"`                 // NFT 状态：active / burned;
			TxHash      string `json:"tx_hash"`                // NFT 发行 Tx Hash
			TimeStamp   string `json:"timestamp"`              // NFT 发行时间戳（UTC 时间）
		} `json:"nfts,omitempty"` // NFT 列表
	} `json:"data"`
}

// QueryNFTRes 查询 EVM 模块 NFT 详情返回值
type QueryNFTRes struct {
	Data struct {
		ID          int    `json:"id"`           // NFT ID
		ClassID     string `json:"class_id"`     // NFT 类别 ID
		ClassName   string `json:"class_name"`   // NFT 类别名称
		ClassSymbol string `json:"class_symbol"` // NFT 类别标识
		Uri         string `json:"uri"`          // 链外数据链接
		UriHash     string `json:"uri_hash"`     // 链外数据 Hash
		Owner       string `json:"owner"`        // NFT 持有者地址
		Status      int    `json:"status"`       // NFT 状态：active / burned;
		TxHash      string `json:"tx_hash"`      // NFT 发行 Tx Hash
		TimeStamp   string `json:"timestamp"`    // NFT 发行时间戳（UTC 时间）
	} `json:"data"`
}

// QueryNativeNFTRes 查询原生模块 NFT 详情返回值
type QueryNativeNFTRes struct {
	Data struct {
		ID          string `json:"id"`                     // NFT ID
		Name        string `json:"name"`                   // NFT 名称
		ClassID     string `json:"class_id"`               // NFT 类别 ID
		ClassName   string `json:"class_name"`             // NFT 类别名称
		ClassSymbol string `json:"class_symbol,omitempty"` // NFT 类别标识
		Uri         string `json:"uri,omitempty"`          // 链外数据链接
		UriHash     string `json:"uri_hash,omitempty"`     // 链外数据 Hash
		Data        string `json:"data,omitempty"`         // 自定义链上元数据
		Owner       string `json:"owner"`                  // NFT 持有者地址
		Status      int    `json:"status"`                 // NFT 状态：active / burned;
		TxHash      string `json:"tx_hash"`                // NFT 发行 Tx Hash
		TimeStamp   string `json:"timestamp"`              // NFT 发行时间戳（UTC 时间）
	} `json:"data"`
}

// QueryNFTHistoryReq 查询 EVM 模块 NFT 操作记录请求参数
type QueryNFTHistoryReq struct {
	PageKey    string `json:"page_key,omitempty"`    // 分页数据的 Key，AVATA 会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	Signer     string `json:"signer,omitempty"`      // Tx 签名者地址
	TxHash     string `json:"tx_hash,omitempty"`     // NFT 操作 Tx Hash
	Operation  string `json:"operation,omitempty"`   // 操作类型：mint / edit / transfer / burn
	StartDate  string `json:"start_date,omitempty"`  // NFT 操作日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // NFT 操作日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total,omitempty"` //是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryNativeNFTHistoryReq 查询原生模块 NFT 操作记录请求参数
type QueryNativeNFTHistoryReq struct {
	PageKey    string `json:"page_key,omitempty"`   // 分页数据的 Key，AVATA 会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`      // 每页记录数，默认为 10，上限为 50
	Signer     string `json:"signer,omitempty"`     // Tx 签名者地址
	TxHash     string `json:"tx_hash,omitempty"`    // NFT 操作 Tx Hash
	Operation  string `json:"operation,omitempty"`  // 操作类型：mint / edit / transfer / burn
	StartDate  string `json:"start_date,omitempty"` // NFT 操作日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`   // NFT 操作日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`    // 排序规则：DATE_ASC / DATE_DESC
	CountTotal string `json:"count_total"`          // 是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryNFTHistoryRes 查询 EVM 模块 NFT 操作记录返回值
type QueryNFTHistoryRes struct {
	Data struct {
		PrevPageKey      string `json:"prev_page_key"` // 上一页数据的 Key，AVATA 会根据该值进行上一页数据的查询
		NextPageKey      string `json:"next_page_key"` // 下一页数据的 Key，AVATA 会根据该值进行下一页数据的查询
		Limit            int    `json:"limit"`         // 每页记录数
		TotalCount       int    `json:"total_count"`   // 总记录数
		OperationRecords []struct {
			TxHash    string `json:"tx_hash,omitempty"`   // NFT 操作的 Tx Hash
			Operation int    `json:"operation,omitempty"` // NFT 操作类型；Enum: "mint" "edit" "transfer" "burn"
			Signer    string `json:"signer,omitempty"`    // Tx 签名者地址
			Recipient string `json:"recipient,omitempty"` // NFT 接收者地址
			TimeStamp string `json:"timestamp,omitempty"` // NFT 操作时间戳（UTC 时间）
		} `json:"operation_records"`
	} `json:"data"`
}

// QueryNativeNFTHistoryRes 查询原生模块 NFT 操作记录返回值
type QueryNativeNFTHistoryRes struct {
	Data struct {
		PrevPageKey      string `json:"prev_page_key"`         // 上一页数据的 Key，AVATA 会根据该值进行上一页数据的查询
		NextPageKey      string `json:"next_page_key"`         // 下一页数据的 Key，AVATA 会根据该值进行下一页数据的查询
		Limit            int    `json:"limit"`                 // 每页记录数
		TotalCount       int    `json:"total_count,omitempty"` // 总记录数
		OperationRecords []struct {
			TxHash    string `json:"tx_hash,omitempty"`   // NFT 操作的 Tx Hash
			Operation int    `json:"operation,omitempty"` // NFT 操作类型；Enum: "mint" "edit" "transfer" "burn"
			Signer    string `json:"signer,omitempty"`    // Tx 签名者地址
			Recipient string `json:"recipient,omitempty"` // NFT 接收者地址
			TimeStamp string `json:"timestamp,omitempty"` // NFT 操作时间戳（UTC 时间）
		} `json:"operation_records,omitempty"`
	} `json:"data"`
}
