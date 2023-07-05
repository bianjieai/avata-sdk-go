package models

const (
	RegisterDomain = "/v3/evm/ns/domains"         // 注册域名
	QueryDomain    = "/v3/evm/ns/domains"         // 查询域名
	TransferDomin  = "/v3/evm/ns/transfers/%s/%s" //转让域名
	QueryDomains   = "/v3/evm/ns/domains/%s"      //查询用户域名
)

// 注册域名请求参数
type RegisterDomainReq struct {
	Name        string `json:"name"`         // 一级域名名称
	Owner       string `json:"owner"`        // 域名拥有者的链账户地址
	Duration    int    `json:"duration"`     //枚举：1,2,3,4,5
	OperationID string `json:"operation_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
}

// 查询域名请求参数
type QueryDomainReq struct {
	Name string `json:"name"`                    // 域名关键字
	Tld  string `json:"tld,omitempty,omitempty"` // 根域名
}

// 查询域名返回结果
type QueryDomainRes struct {
	Data struct {
		Domains []struct {
			Name            string `json:"name"`                       // 一级域名
			Status          int    `json:"status"`                     // 当前域名状态
			Owner           string `json:"owner,omitempty"`            // 当前域名拥有者的链账户地址
			Expire          int    `json:"expire"`                     // 当前域名过期状态
			ExpireTimestamp int    `json:"expire_timestamp,omitempty"` // 当前域名过期时间戳
			Msg             string `json:"msg,omitempty"`              // 提示信息
		} `json:"domains"`
	} `json:"data"`
}

// 转让域名请求参数
type TransferDominReq struct {
	Recipient   string `json:"recipient"`    // 域名接收者地址
	OperationID string `json:"operation_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
}

// 查询用户域名
type QueryDomainsReq struct {
	Name string `json:"name"`          // 域名关键字
	Tld  string `json:"tld,omitempty"` // 根域名
}

// 查询用户域名返回值
type QueryDomainsRes struct {
	Data struct {
		Domains []struct {
			Name            string `json:"name"`             //一级域名， 如：test.wallet
			Status          int    `json:"status"`           //当前域名状态 0：未注册； 1：已注册；
			Owner           string `json:"owner"`            //当前域名拥有者的链账户地址
			Expire          int    `json:"expire"`           //当前域名过期状态 0 ：未过期； 1：已过期
			ExpireTimestamp int    `json:"expire_timestamp"` //当前域名过期时间戳
			Msg             string `json:"msg"`              //提示信息
		} `json:"domains"`
	} `json:"data"`
}
