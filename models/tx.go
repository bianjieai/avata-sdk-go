package models

const (
	QueryTxResult       = "/v3/evm/tx/%s"            // EVM 模块上链交易结果查询接口
	QueryNativeTxResult = "/v3/native/tx/%s"         // 原生模块上链交易结果查询接口
	QueryTxTypes        = "/v3/evm/dict/tx_types"    // EVM 模块查询枚举值列表接口
	QueryNativeTxTypes  = "/v3/native/dict/tx_types" // 原生模块查询枚举值列表接口

)

// QueryTxResultRes EVM 模块上链交易结果查询接口返回值
// 交易状态说明：
// status 为 0（处理中），上链请求正在处理，请等待处理完成；
// status 为 1（成功），交易已上链并执行成功；
// status 为 2（失败），说明该交易执行失败。请在业务侧做容错处理。可以参考接口返回的 message（交易失败的错误描述信息） 对 NFT / MT / 业务接口的请求参数做适当调整后，使用「新的 Operation ID 」重新发起 NFT / MT / 业务接口请求。
// status 为 3（未处理），上链请求还在等待处理，请稍等；
type QueryTxResultRes struct {
	Data struct {
		Module      int    `json:"module"`       // 交易模块；Enum: 1 nft  2 ns(域名)  3 record(存证) 4 合约调用；
		Operation   int    `json:"operation"`    // 用户操作类型；
		TxHash      string `json:"tx_hash"`      // 交易哈希
		Status      int    `json:"status"`       // 交易状态， 0 处理中； 1 成功； 2 失败； 3 未处理；Enum: 0 1 2 3
		Message     string `json:"message"`      // 交易失败的错误描述信息
		BlockHeight int    `json:"block_height"` // 交易上链的区块高度
		Timestamp   string `json:"timestamp"`    // 交易上链时间（UTC 时间）
		Nft         struct {
			ClassId string `json:"class_id"` // NFT 类别 ID
			ID      int    `json:"id"`       // NFT ID
		} `json:"nft"` // 具体参考接口文档
		Ns struct {
			Name    string `json:"name"`    // 域名
			Owner   string `json:"owner"`   // 域名拥有者的链账户地址
			Expires int    `json:"expires"` // 当前域名过期时间戳
			Node    string `json:"node"`    // 域名 node key
			Addr    struct {
				BlockChain int    `json:"block_chain"` // 底层区块链
				AddrValue  string `json:"addr_value"`  // 链账户地址
			} `json:"addr"`
			Text struct {
				Key       string `json:"key"`        // 文本数据 key
				TextValue string `json:"text_value"` // 文本数据值
			} `json:"text"`
			Address string `json:"addr"` // 链账户地址
		} `json:"ns"` // 具体参考接口文档

	} `json:"data"`
}

// QueryNativeTxResultRes 原生模块上链交易结果查询接口返回值
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
			ClassId string `json:"class_id"` // NFT 类别 ID
			ID      string `json:"id"`       // NFT ID
		} `json:"nft"` // 具体参考接口文档
		Mt struct {
			MtId    string `json:"mt_id"`    // MT ID
			ClassId string `json:"class_id"` // MT 类别 ID
		} `json:"mt"`
		Record struct {
			RecordId       string `json:"record_id"`       // 区块链存证 ID
			CertificateUrl string `json:"certificate_url"` // 区块链存证证书的下载链接；证书下载链接并非长期有效，请您尽快将证书文件下载至本地并妥善保管。
		} `json:"record"`
	} `json:"data"`
}

// QueryTxTypesRes 查询枚举值列表接口返回值
type QueryTxTypesRes struct {
	Data DataItem `json:"data"`
}

// QueryNativeTxTypesRes  原生模块查询枚举值列表接口返回值
type (
	QueryNativeTxTypesRes struct {
		Data DataItem `json:"data"`
	}

	DataItem struct {
		Data []TxType `json:"data"`
	}

	TxType struct {
		Module      int64  `json:"module"`      // 交易模块：请通过 查询枚举值列表 接口查看
		Operation   int64  `json:"operation"`   // 用户操作类型 请通过 查询枚举值列表 接口查看
		Code        string `json:"code"`        // 标识
		Name        string `json:"name"`        // 名称
		Description string `json:"description"` // 描述
	}
)
