package models

const (
	QueryTxResult       = "/v3/evm/tx/%s"            // 上链交易结果查询/evm
	QueryNativeTxResult = "/v3/native/tx/%s"         // 以原生方式上链交易结果查询接口
	QueryTxTypes        = "/v3/evm/dict/tx_types"    // 查询枚举值列表
	QueryNativeTxTypes  = "/v3/native/dict/tx_types" // 以原生方式查询枚举值列表

)

// QueryTxResultRes 上链交易结果查询返回值
// 交易状态说明：
// status 为 3（未处理），上链请求还在等待处理，请稍等；
// status 为 0（处理中），上链请求正在处理，请等待处理完成；
// status 为 1（成功），交易已上链并执行成功；
// status 为 2（失败），说明该交易执行失败。请在业务侧做容错处理。可以参考接口返回的 message（交易失败的错误描述信息） 对 NFT / MT / 业务接口的请求参数做适当调整后，使用「新的 Operation ID 」重新发起 NFT / MT / 业务接口请求。
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

// QueryNativeTxResultRes 以原生方式查询上链交易结果
type QueryNativeTxResultRes struct {
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
			ID      string `json:"id"`
		} `json:"nft"` // 具体参考接口文档
		Mt struct {
			MtId    string `json:"mt_id"`
			ClassId string `json:"class_id"`
		}
		Record struct {
			RecordId       string `json:"record_id"`
			CertificateUrl string `json:"certificate_url"`
		}
	} `json:"data"`
}

// QueryTxTypesRes 查询枚举值列表
type QueryTxTypesRes struct {
	Data DataItem `json:"data"`
}

// QueryNativeTxTypesRes  以原生方式查询枚举值列表
type QueryNativeTxTypesRes struct {
	Data DataItem `json:"data"`
}
type DataItem struct {
	Data []TxType `json:"data"`
}
type TxType struct {
	Module      int64  `json:"module"`
	Operation   int64  `json:"operation"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
