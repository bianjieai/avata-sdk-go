package models

const (
	SetResolves          = "/v3/evm/ns/resolves/%s/%s"      // 设置域名解析
	QueryResolves        = "/v3/evm/ns/resolves/%s"         // 查询域名解析
	SetReverseResolves   = "/v3/evm/ns/reverse-resolves/%s" // 设置域名反向解析
	QueryReverseResolves = "/v3/evm/ns/reverse-resolves/%s" // 查询域名反向解析
)

// SetResolvesReq 设置域名解析请求参数
type SetResolvesReq struct {
	ResolveType int `json:"resolve_type"` // 域名解析类型 1：链账户 2：文本
	Addr        struct {
		BlockChain int    `json:"block_chain"`          // 底层区块链 1000：天舟链  1001：天和链 1002: 神舟链
		AddrValue  string `json:"addr_value,omitempty"` // 链账户地址，默认值为当前域名 owner
	} `json:"addr,omitempty"` // 链账户 resolve_type 为 1 必填
	Text struct {
		Key       string `json:"key,omitempty"` // 文本数据 key 可选：email、avatar、description、notice、keywords
		TextValue string `json:"text_value"`    // 文本数据值
	} `json:"text,omitempty"` // 文本数据 resolve_type 为 2 必填
	OperationID string `json:"operation_id"` // 操作 ID
}

// QueryResolvesReq 查询域名解析请求参数
type QueryResolvesReq struct {
	ResolveType string `json:"resolve_type,omitempty"` // 域名解析类型 0：全部 1：链账户 2：文本
}

// QueryResolvesRes 查询域名解析返回值
type QueryResolvesRes struct {
	Data struct {
		Addrs []struct {
			BlockChain int    `json:"block_chain"` // 底层区块链 1000：天舟链  1001：天和链 1002: 神舟链
			AddrValue  string `json:"addr_value"`  // 链账户地址，默认值为当前域名 owner
		} `json:"addrs"` //链账户resolve_type为1必填
		Texts []struct {
			Key       string `json:"key"`        // 文本数据 key 可选：email、avatar、description、notice、keywords
			TextValue string `json:"text_value"` // 文本数据值
		} `json:"texts"` // 文本数据 resolve_type 为 2 必填
	} `json:"data"`
}

// SetReverseResolvesReq 设置域名反向解析请求参数
type SetReverseResolvesReq struct {
	Name        string `json:"name"`         // 域名名称
	OperationID string `json:"operation_id"` // 操作 ID
}

// QueryReverseResolvesRes 查询域名反向解析返回值
type QueryReverseResolvesRes struct {
	Data struct {
		Name string `json:"name"` // 域名
	} `json:"data"`
}
