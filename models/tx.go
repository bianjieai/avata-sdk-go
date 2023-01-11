package models

const (
	GetTxResult    = "/v1beta1/tx/%s"         // 上链交易结果查询接口
	GetTxQueueInfo = "/v1beta1/tx/queue/info" // 上链交易排队状态查询接口
)

// GetTxResultRes 上链交易结果查询返回值
type GetTxResultRes struct {
	BaseRes
	Data struct {
		Type        string `json:"type"`
		Module      string `json:"module"`
		TxHash      string `json:"tx_hash"`
		Status      int    `json:"status"`
		Message     string `json:"message"`
		BlockHeight int    `json:"block_height"`
		Timestamp   string `json:"timestamp"`
		Tag         struct {
			Key1 string `json:"key1"`
			Key2 string `json:"key2"`
			Key3 string `json:"key3"`
		} `json:"tag"`
		Nft struct {
			ClassId string `json:"class_id"`
			NftId   string `json:"nft_id"`
		} `json:"nft"`
		Mt struct {
			ClassId string `json:"class_id"`
			MtId    string `json:"mt_id"`
		} `json:"mt"`
		Record struct {
			RecordId       string `json:"record_id"`
			CertificateUrl string `json:"certificate_url"`
		} `json:"record"`
	} `json:"data"`
}

// GetTxQueueInfoReq 上链交易排队状态查询请求参数
type GetTxQueueInfoReq struct {
	OperationID string `json:"operation_id"`
}

// GetTxQueueInfoRes 上链交易排队状态查询返回值
type GetTxQueueInfoRes struct {
	BaseRes
	Data struct {
		QueueTotal       int    `json:"queue_total"`
		QueueRequestTime string `json:"queue_request_time"`
		QueueCostTime    int    `json:"queue_cost_time"`
		TxQueuePosition  int    `json:"tx_queue_position"`
		TxRequestTime    string `json:"tx_request_time"`
		TxCostTime       int    `json:"tx_cost_time"`
		TxMessage        string `json:"tx_message"`
	} `json:"data"`
}
