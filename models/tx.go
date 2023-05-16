package models

const (
	QueryTxResult    = "/v2/tx/%s"         // 上链交易结果查询接口
	QueryTxQueueInfo = "/v2/tx/queue/info" // 上链交易排队状态查询接口
)

// QueryTxResultRes 上链交易结果查询返回值
//交易状态说明：
//status 为 3（未处理），上链请求还在等待处理，请稍等；
//status 为 0（处理中），上链请求正在处理，请等待处理完成；
//status 为 1（成功），交易已上链并执行成功；
//status 为 2（失败），说明该交易执行失败。请在业务侧做容错处理。可以参考接口返回的 message（交易失败的错误描述信息） 对 NFT / MT / 业务接口的请求参数做适当调整后，使用「新的 Operation ID 」重新发起 NFT / MT / 业务接口请求。
type QueryTxResultRes struct {
	Data struct {
		Module      int    `json:"module"`       // 交易模块；Enum: 1 nft  2 ns(域名)  3 record(存证)
		Operation   int    `json:"operation"`    // 用户操作类型；Enum: 1：issue_class； 2：transfer_class； 3：mint_nft； 4：edit_nft； 5：transfer_nft； 6：burn_nft
		TxHash      string `json:"tx_hash"`      // 交易哈希
		Status      int    `json:"status"`       // 交易状态， 0 处理中； 1 成功； 2 失败； 3 未处理；Enum: 0 1 2 3
		Message     string `json:"message"`      // 交易失败的错误描述信息
		BlockHeight int    `json:"block_height"` // 交易上链的区块高度
		Timestamp   string `json:"timestamp"`    // 交易上链时间（UTC 时间）
		Nft         struct {
			ClassId string `json:"class_id"`
			ID      int    `json:"id"`
		} `json:"nft"` // 具体参考接口文档
		Ns struct {
			Name    string `json:"name"`
			Owner   string `json:"owner"`
			Expires int    `json:"expires"`
		} `json:"ns"` // 具体参考接口文档
	} `json:"data"`
}

// QueryTxQueueInfoReq 上链交易排队状态查询请求参数
type QueryTxQueueInfoReq struct {
	OperationID string `json:"operation_id,omitempty"` // 操作 ID，是指用户在进行具体的NFT/MT/业务接口请求时，自定义的操作ID。注意：不支持创建链账户/批量创建链账户的操作 ID 查询。
}

// QueryTxQueueInfoRes 上链交易排队状态查询返回值
type QueryTxQueueInfoRes struct {
	Data struct {
		QueueTotal       int    `json:"queue_total"`        // 当前队列中待处理交易总数
		QueueRequestTime string `json:"queue_request_time"` // 当前队列即将被处理交易的请求时间（UTC 时间）
		QueueCostTime    int    `json:"queue_cost_time"`    // 当前队列中所有交易处理完预估时间（秒）
		TxQueuePosition  int    `json:"tx_queue_position"`  // Operation ID 对应交易所处队列中的位置；若交易存在队列中，0 则表示正在重试
		TxRequestTime    string `json:"tx_request_time"`    // Operation ID 对应交易的请求时间（UTC 时间）
		TxCostTime       int    `json:"tx_cost_time"`       // Operation ID 对应交易预估处理所需时间（秒）
		TxMessage        string `json:"tx_message"`         // Operation ID 对应交易排队描述信息
	} `json:"data"`
}
