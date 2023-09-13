package models

// 验证签名时传入的 API 版本号
const (
	APIVersionV1     = "V1" // V1 版本 AVATA Open API
	APIVersionsOther = ""   // 其它版本 AVATA Open API
)

// 区分调用不同服务模块的接口
const (
	Native = "native"
	EVM    = "evm"
)

// OnCallbackResV1 AVATA Open API V1 版本回调服务参数
type (
	OnCallbackResV1 struct {
		OperationId string `json:"operation_id"`           // 操作 ID
		Module      string `json:"module"`                 // 交易模块
		Type        string `json:"type"`                   // 用户操作类型
		Status      int32  `json:"status"`                 // 交易状态，1 成功; 2 失败
		TxHash      string `json:"tx_hash,omitempty"`      // 交易哈希
		Message     string `json:"message,omitempty"`      // 交易失败的错误描述信息
		BlockHeight int64  `json:"block_height,omitempty"` // 交易上链的区块高度
		Timestamp   string `json:"timestamp,omitempty"`    // 交易上链时间（UTC 时间）
		Nft         NftV1  `json:"nft,omitempty"`          // 对应不同操作类型的消息体
		Mt          Mt     `json:"mt,omitempty"`           // 对应不同操作类型的消息体
		Record      Record `json:"record,omitempty"`       // 对应不同操作类型的消息体
	}
)

// OnCallbackResNative AVATA Open API V2 及以上版本 Native 模块接口回调服务参数
type OnCallbackResNative struct {
	Kind        string    `json:"kind"`                   // 区分服务，native / evm
	OperationId string    `json:"operation_id"`           // 操作 ID
	Module      int32     `json:"module"`                 // 交易模块
	Operation   int32     `json:"operation"`              // 用户操作类型
	TxHash      string    `json:"tx_hash,omitempty"`      // 交易哈希
	Status      int32     `json:"status"`                 // 交易状态，1 成功; 2 失败
	Message     string    `json:"message,omitempty"`      // 交易失败的错误描述信息
	BlockHeight int64     `json:"block_height,omitempty"` // 交易上链的区块高度
	Timestamp   string    `json:"timestamp,omitempty"`    // 交易上链时间（UTC 时间）
	Nft         NftNative `json:"nft,omitempty"`          // 对应不同操作类型的消息体
	Mt          Mt        `json:"mt,omitempty"`           // 对应不同操作类型的消息体
	Record      Record    `json:"record,omitempty"`       // 对应不同操作类型的消息体
}

// OnCallbackResEVM V2 及以上版本 EVM 模块接口回调服务参数
type OnCallbackResEVM struct {
	Kind        string `json:"kind"`                   // 区分服务，native / evm
	OperationId string `json:"operation_id"`           // 操作 ID
	Module      int32  `json:"module"`                 // 交易模块
	Operation   int32  `json:"operation"`              // 用户操作类型
	TxHash      string `json:"tx_hash,omitempty"`      // 交易哈希
	Status      int32  `json:"status"`                 // 交易状态，1 成功; 2 失败
	Message     string `json:"message,omitempty"`      // 交易失败的错误描述信息
	BlockHeight int64  `json:"block_height,omitempty"` // 交易上链的区块高度
	Timestamp   string `json:"timestamp,omitempty"`    // 交易上链时间（UTC 时间）
	Nft         NftEVM `json:"nft,omitempty"`          // 对应不同操作类型的消息体
}

type Mt struct {
	ClassId string `json:"class_id,omitempty"` // MT 类别 ID
	MtId    string `json:"mt_id,omitempty"`    // MT ID
}

type Record struct {
	RecordId       string `json:"record_id"`       // 区块链存证 ID
	CertificateUrl string `json:"certificate_url"` // 区块链存证证书的下载链接；证书下载链接并非长期有效，请您尽快将证书文件下载至本地并妥善保管。
}

type NftV1 struct {
	ClassId string `json:"class_id,omitempty"` // 类别 ID
	NftId   string `json:"nft_id,omitempty"`   // NFT ID
}

type NftNative struct {
	ClassId string `json:"class_id,omitempty"` // 类别 ID
	Id      string `json:"id,omitempty"`       // NFT ID
}

type NftEVM struct {
	ClassId string `json:"class_id,omitempty"` // 类别 ID
	Id      int64  `json:"id,omitempty"`       // NFT ID
}

type KindRes struct {
	Kind string `json:"kind"` // 区分服务，native / evm
}
