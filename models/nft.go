package models

const (
	CreateNFTClass    = "/v1beta1/nft/classes"                // 创建 NFT 类别
	QueryNFTClass     = "/v1beta1/nft/classes"                // 查询 NFT 类别
	QueryNFTClassById = "/v1beta1/nft/classes/%s"             // 查询 NFT 类别详情
	TransfersNFClass  = "/v1beta1/nft/class-transfers/%s/%s"  // 转让 NFT 类别
	CreateNFT         = "/v1beta1/nft/nfts/%s"                // 发行 NFT
	TransferNFT       = "/v1beta1/nft/nft-transfers/%s/%s/%s" // 转让 NFT
	EditNFT           = "/v1beta1/nft/nfts/%s/%s/%s"          // 编辑 NFT
	DeleteNFT         = "/v1beta1/nft/nfts/%s/%s/%s"          // 销毁 NFT
	BatchCreateNFT    = "/v1beta1/nft/batch/nfts/%s"          // 批量发行 NFT
	BatchTransferNFT  = "/v1beta1/nft/batch/nft-transfers/%s" // 批量转让 NFT
	BatchEditNFT      = "/v1beta1/nft/batch/nfts/%s"          // 批量编辑 NFT
	BatchDeleteNFT    = "/v1beta1/nft/batch/nfts/%s"          // 批量销毁 NFT
	QueryNFT          = "/v1beta1/nft/nfts"                   // 查询 NFT
	QueryNFTById      = "/v1beta1/nft/nfts/%s/%s"             // 查询 NFT 详情
	QueryNFTHistory   = "/v1beta1/nft/nfts/%s/%s/history"     // 查询 NFT 历史记录
)

// CreateNFTClassReq 创建 NFT 类别：request
type CreateNFTClassReq struct {
	Name        string            `json:"name"`
	ClassID     string            `json:"class_id,omitempty"`
	Symbol      string            `json:"symbol,omitempty"`
	Description string            `json:"description,omitempty"`
	Uri         string            `json:"uri,omitempty"`
	UriHash     string            `json:"uri_hash,omitempty"`
	Data        string            `json:"data,omitempty"`
	Owner       string            `json:"owner"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationID string            `json:"operation_id"`
}

// QueryNFTClassReq 查询 NFT 类别：request
type QueryNFTClassReq struct {
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

// QueryNFTClassRes 查询 NFT 类别：Response
type QueryNFTClassRes struct {
	Data struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		TotalCount int `json:"total_count"`
		Classes    []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Symbol    string `json:"symbol"`
			NFTCount  int    `json:"nft_count"`
			Uri       string `json:"uri"`
			Owner     string `json:"owner"`
			TxHash    string `json:"tx_hash"`
			TimeStamp string `json:"timestamp"`
		} `json:"classes"`
	} `json:"data"`
}

// QueryNFTClassByIdReq 查询 NFT 类别详情：request
type QueryNFTClassByIdReq struct {
	ID string `json:"id"`
}

// QueryNFTClassByIdRes 查询 NFT 类别详情：Response
type QueryNFTClassByIdRes struct {
	Data struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		Description string `json:"description"`
		NFTCount    int    `json:"nft_count"`
		Uri         string `json:"uri"`
		UriHash     string `json:"uri_hash"`
		Data        string `json:"data"`
		Owner       string `json:"owner"`
		TxHash      string `json:"tx_hash"`
		TimeStamp   string `json:"timestamp"`
	} `json:"data"`
}

// TransfersNFClassReq 转让 NFT 类别：request
type TransfersNFClassReq struct {
	Recipient   string            `json:"recipient"`
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag,omitempty"`
}

// CreateNFTReq 发行nft：request
type CreateNFTReq struct {
	Name        string            `json:"name"`
	Uri         string            `json:"uri,omitempty"`
	UriHash     string            `json:"uri_hash,omitempty"`
	Data        string            `json:"data,omitempty"`
	Recipient   string            `json:"recipient,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
	OperationID string            `json:"operation_id"`
}

// TransferNFTReq 转让 NFT ：request
type TransferNFTReq struct {
	Recipient   string            `json:"recipient"`
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag,omitempty"`
}

// EditNFTReq 编辑 NFT ：request
type EditNFTReq struct {
	Name        string            `json:"name"`
	Uri         string            `json:"uri,omitempty"`
	Data        string            `json:"data,omitempty"`
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag,omitempty"`
}

// DeleteNFTReq 销毁 NFT ：request
type DeleteNFTReq struct {
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag,omitempty"`
}

// BatchCreateNFTReq 批量发行nft：request
type (
	BatchCreateNFTReq struct {
		Name        string            `json:"name"`
		Uri         string            `json:"uri,omitempty"`
		UriHash     string            `json:"uri_hash,omitempty"`
		Data        string            `json:"data,omitempty"`
		Recipients  []Recipients      `json:"recipients"`
		Tag         map[string]string `json:"tag,omitempty"`
		OperationID string            `json:"operation_id"`
	}

	Recipients struct {
		Amount    int    `json:"amount"`
		Recipient string `json:"recipient"`
	}
)

// BatchTransferNFTReq 批量转让 NFT ：request
type (
	BatchTransferNFTReq struct {
		Data        []BatchTransferNFTData `json:"data"`
		Tag         map[string]string      `json:"tag,omitempty"`
		OperationID string                 `json:"operation_id"`
	}
	BatchTransferNFTData struct {
		NFTs      []BatchTransferNFTs `json:"nfts"`
		Recipient string              `json:"recipient"`
	}
	BatchTransferNFTs struct {
		ClassID string `json:"class_id"`
		NFTID   string `json:"nft_id"`
	}
)

// BatchEditNFTReq 批量编辑 NFT ：request
type (
	BatchEditNFTReq struct {
		NFTs        []BatchEditNfts   `json:"nfts"`
		Tag         map[string]string `json:"tag,omitempty"`
		OperationID string            `json:"operation_id"`
	}
	BatchEditNfts struct {
		ClassID string `json:"class_id"`
		NFTID   string `json:"nft_id"`
		Name    string `json:"name"`
		Uri     string `json:"uri,omitempty"`
		Data    string `json:"data,omitempty"`
	}
)

// BatchDeleteNFTReq 批量销毁 NFT ：request
type (
	BatchDeleteNFTReq struct {
		NFTs        []NFTs            `json:"nfts"`
		Tag         map[string]string `json:"tag,omitempty"`
		OperationID string            `json:"operation_id"`
	}
	NFTs struct {
		ClassID string `json:"class_id"`
		NFTID   string `json:"nft_id"`
	}
)

// QueryNFTReq 查询 NFT ：request
type QueryNFTReq struct {
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

// QueryNFTRes 查询 NFT ：Response
type QueryNFTRes struct {
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

// QueryNFTByIdReq 查询 NFT 详情：request
type QueryNFTByIdReq struct {
	ClassID string `json:"class_id"`
	NFTID   string `json:"nft_id"`
}

// QueryNFTByIdRes 查询 NFT 详情：Response
type QueryNFTByIdRes struct {
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
