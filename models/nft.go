package models

const (
	CreateNFTClass    = "/v1beta1/nft/classes"             //创建nft类别
	QueryNFTClass     = "/v1beta1/nft/classes"             //查询nft类别
	QueryNFTClassById = "/v1beta1/nft/classes"             //查询nft类别详情
	TransfersNFClass  = "/v1beta1/nft/class-transfers"     //转让nft类别
	CreateNFT         = "/v1beta1/nft/nfts"                //发行nft
	TransfersNFT      = "/v1beta1/nft/nft-transfers"       //转让nft
	EditorNFT         = "/v1beta1/nft/nfts"                //编辑nft
	DeleteNFT         = "/v1beta1/nft/nfts"                //销毁nft
	CreateNFTBatch    = "/v1beta1/nft/batch/nfts"          //批量发行nft
	TransfersNFTBatch = "/v1beta1/nft/batch/nft-transfers" //批量转让nft
	EditorNFTBatch    = "/v1beta1/nft/batch/nfts"          //批量编辑nft
	DeleteNFTBatch    = "/v1beta1/nft/batch/nfts"          //批量销毁nft
	QueryNFT          = "/v1beta1/nft/nfts"                //查询nft
	QueryNFTById      = "/v1beta1/nft/nfts"                //查询nft详情
	QueryNFTHistory   = "/v1beta1/nft/nfts"                //查询nft历史记录
)

//创建nft类别：request
type CreateNFTClassReq struct {
	Name        string            `json:"name"`
	ClassID     string            `json:"class_id"`
	Symbol      string            `json:"symbol"`
	Description string            `json:"description"`
	Uri         string            `json:"uri"`
	UriHash     string            `json:"uri_hash"`
	Data        string            `json:"data"`
	Owner       string            `json:"owner"`
	Tag         map[string]string `json:"tag"`
	OperationID string            `json:"operation_id"`
}

//创建nft类别：response
type CreateNFTClassResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//查询nft类别：request
type QueryNFTClassReq struct {
	Offset    string `json:"offset"`
	Limit     string `json:"limit"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	TxHash    string `json:"tx_hash"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	SortBy    string `json:"sort_by"`
}

//查询nft类别：response
type QueryNFTClassResp struct {
	BaseRes
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

//查询nft类别详情：request
type QueryNFTClassByIdReq struct {
	ID string `json:"id"`
}

//查询nft类别详情：response
type QueryNFTClassByIdResp struct {
	BaseRes
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

//转让nft类别：request
type TransfersNFClassReq struct {
	Recipient   string            `json:"recipient"`
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag"`
}

//转让nft类别：response
type TransfersNFClassResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//发行nft：request
type CreateNFTReq struct {
	Name        string            `json:"name"`
	Uri         string            `json:"uri"`
	UriHash     string            `json:"uri_hash"`
	Data        string            `json:"data"`
	Recipient   string            `json:"recipient"`
	Tag         map[string]string `json:"tag"`
	OperationID string            `json:"operation_id"`
}

//发行nft：response
type CreateNFTResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//转让nft：request
type TransfersNFTReq struct {
	Recipient   string            `json:"recipient"`
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag"`
}

//转让nft：response
type TransfersNFTResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//编辑ft：request
type EditorNFTReq struct {
	Name        string            `json:"name"`
	Uri         string            `json:"uri"`
	Data        string            `json:"data"`
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag"`
}

//编辑nft：response
type EditorNFTResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//销毁nft：request
type DeleteNFTReq struct {
	OperationID string            `json:"operation_id"`
	Tag         map[string]string `json:"tag"`
}

//销毁nft：response
type DeleteNFTResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//批量发行nft：request
type (
	CreateNFTBatchReq struct {
		Name                        string `json:"name"`
		Uri                         string `json:"uri"`
		UriHash                     string `json:"uri_hash"`
		Data                        string `json:"data"`
		CreateNFTBatchReqRecipients `json:"recipients"`
		Tag                         map[string]string `json:"tag"`
		OperationID                 string            `json:"operation_id"`
	}
	CreateNFTBatchReqRecipients []struct {
		Amount    int    `json:"amount"`
		Recipient string `json:"recipient"`
	}
)

//批量发行nft：response
type CreateNFTBatchResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//批量转让nft：request
type (
	TransfersNFTBatchReq struct {
		TransfersNFTBatchReqData `json:"data"`
		Tag                      map[string]string `json:"tag"`
		OperationID              string            `json:"operation_id"`
	}
	TransfersNFTBatchReqData []struct {
		TransfersNFTBatchReqNFTs `json:"nfts"`
		Recipient                string `json:"recipient"`
	}
	TransfersNFTBatchReqNFTs []struct {
		ClassID string `json:"class_id"`
		NFTID   string `json:"nft_id"`
	}
)

//批量转让nft：response
type TransfersNFTBatchResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//批量编辑nft：request
type (
	EditorNFTBatchReq struct {
		EditorNFTBatchReqNFTs `json:"nfts"`
		Tag                   map[string]string `json:"tag"`
		OperationID           string            `json:"operation_id"`
	}
	EditorNFTBatchReqNFTs []struct {
		ClassID string `json:"class_id"`
		NFTID   string `json:"nft_id"`
		Name    string `json:"name"`
		Uri     string `json:"uri"`
		Data    string `json:"data"`
	}
)

//批量编辑nft：response
type EditorNFTBatchResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//批量销毁nft：request
type (
	DeleteNFTBatchReq struct {
		DeleteNFTBatchReqNFTs `json:"nfts"`
		Tag                   map[string]string `json:"tag"`
		OperationID           string            `json:"operation_id"`
	}
	DeleteNFTBatchReqNFTs []struct {
		ClassID string `json:"class_id"`
		NFTID   string `json:"nft_id"`
	}
)

//批量销毁nft：response
type DeleteNFTBatchResp struct {
	BaseRes
	Data struct {
		OperationID string `json:"operation_id"`
	} `json:"data"`
}

//查询nft：request
type QueryNFTReq struct {
	Offset    string `json:"offset"`
	Limit     string `json:"limit"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	ClassID   string `json:"class_id"`
	Owner     string `json:"owner"`
	TxHash    string `json:"tx_hash"`
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	SortBy    string `json:"sort_by"`
}

//查询nft：response
type QueryNFTResp struct {
	BaseRes
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

//查询nft详情：request
type QueryNFTByIdReq struct {
	ClassID string `json:"class_id"`
	NFTID   string `json:"nft_id"`
}

//查询nft详情：response
type QueryNFTByIdResp struct {
	BaseRes
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

//查询nft操作记录：request
type QueryNFTHistoryReq struct {
	Offset    string `json:"offset"`
	Limit     string `json:"limit"`
	Signer    string `json:"signer"`
	TxHash    string `json:"tx_hash"`
	Operation string `json:"operation"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	SortBy    string `json:"sort_by"`
}

//查询nft操作记录：response
type QueryNFTHistoryResp struct {
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
			TimeStamp string `json:"timestamp"`
		} `json:"operation_records"`
	} `json:"data"`
}
