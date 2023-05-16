package models

const (
	RegisterDomain = "/v2/ns/domains"         // 注册域名
	QueryDomain    = "/v2/ns/domains"         // 查询域名
	TransferDomin  = "/v2/ns/transfers/%s/%s" //转让域名
	QueryDomains   = "/v2/ns/domains/%s"      //查询用户域名
)

//注册域名请求参数
type RegisterDomainReq struct {
	Name        string `json:"name"`         // 一级域名名称
	Owner       string `json:"owner"`        // 域名拥有者的链账户地址
	OperationID string `json:"operation_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
	Duration    int    `json:"duration"`     //枚举：1,2,3,4,5
}

//查询域名请求参数
type QueryDomainReq struct {
	Name string `json:"name"`          // 域名关键字
	Tld  string `json:"tld,omitempty"` // 根域名
}

//查询域名返回结果
type QueryDomainRes struct {
	Data struct {
		Domains []struct {
			Name            string `json:"name"`
			Status          int    `json:"status"`
			Owner           string `json:"owner"`
			Expire          int    `json:"expire"`
			ExpireTimestamp int    `json:"expire_timestamp"`
			Msg             string `json:"msg"`
		} `json:"domains"`
	} `json:"data"`
}

//注册域名请求参数
type TransferDominReq struct {
	Recipient   string `json:"recipient"`    //域名接收者地址
	OperationID string `json:"operation_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
}

//查询用户域名
type QueryDomainsReq struct {
	Name       string `json:"name,omitempty"`        // 域名关键字
	Tld        string `json:"tld,omitempty"`         // 根域名
	PageKey    string `json:"page_key,omitempty"`    //分页数据的Key， Avata会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       //每页记录数，默认为 10，上限为 50
	CountTotal string `json:"count_total,omitempty"` //是否查询数据的总数量 0：不查询总数（默认）1：查询总数
}

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
		PrevPageKey string `json:"prev_page_key"` //上一页数据的Key
		NextPageKey string `json:"next_page_key"` //下一页数据的Key
		Limit       int    `json:"limit"`         //每页记录数
		TotalCount  int    `json:"total_count"`   //总记录数
	} `json:"data"`
}
