package models

const (
	CreateOrder      = "/v1beta1/orders"       // 购买能量值/业务费接口
	QueryOrder       = "/v1beta1/orders/%s"    // 查询能量值/业务费购买结果接口
	QueryOrders      = "/v1beta1/orders"       // 查询能量值/业务费购买结果列表接口
	BatchCreateOrder = "/v1beta1/orders/batch" // 批量购买能量值接口
)

// CreateOrderReq 购买能量值/业务费接口请求参数
type CreateOrderReq struct {
	Account   string `json:"account"`    // 链账户地址
	Amount    int    `json:"amount"`     // 购买金额 ，只能购买整数元金额；单位：分
	OrderType string `json:"order_type"` // 充值类型：gas：能量值；business：业务费；Enum: "gas" "business"
	OrderId   string `json:"order_id"`   // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
}

// OrderRes 购买能量值/业务费接口返回值/批量购买能量值接口返回值
type OrderRes struct {
	Data struct {
		OrderId string `json:"order_id"` // 交易流水号（用户发起交易时传入的交易流水号)
	} `json:"data"`
}

// QueryOrdersReq 查询能量值/业务费购买结果列表接口请求参数
type QueryOrdersReq struct {
	Offset    string `json:"offset,omitempty"`     // 游标，默认为 0
	Limit     string `json:"limit,omitempty"`      // 每页记录数，默认为 10，上限为 50
	Status    string `json:"status,omitempty"`     // 订单状态：success 充值成功 / failed 充值失败 / pending 正在充值
	StartDate string `json:"start_date,omitempty"` // 充值订单创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   // 充值订单创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    // 排序规则：DATE_ASC / DATE_DESC，默认为 DATE_DESC
}

// QueryOrdersRes 查询能量值/业务费购买结果列表接口返回值
type QueryOrdersRes struct {
	Data struct {
		Offset     int `json:"offset"`      // 游标
		Limit      int `json:"limit"`       // 每页记录数
		TotalCount int `json:"total_count"` // 总记录数
		OrderInfos []struct {
			OrderId    string `json:"order_id"`    // 订单流水号
			Status     string `json:"status"`      // 订单状态：success 充值成功 / failed 充值失败 / pending 正在充值
			Message    string `json:"message"`     // 订单失败的错误描述信息
			Account    string `json:"account"`     // 链账户地址 （调用「批量购买能量值」接口不展示此字段）
			Amount     string `json:"amount"`      // 充值金额，为整数元金额；单位：分 （调用「批量购买能量值」接口不展示此字段）
			Number     string `json:"number"`      // 充值的数量，充值 gas 该值单位为 ugas，充值业务费单位为分（调用「批量购买能量值」接口不展示此字段）
			CreateTime string `json:"create_time"` // 创建时间（UTC 时间）
			UpdateTime string `json:"update_time"` // 最后操作时间（UTC 时间）
			OrderType  string `json:"order_type"`  // 订单类型，gas / business
		} `json:"order_infos"`
	} `json:"data"`
}

// QueryOrderRes 查询能量值/业务费购买结果接口返回值
type QueryOrderRes struct {
	Data struct {
		OrderId    string `json:"order_id"`    // 订单流水号
		Status     string `json:"status"`      // 订单状态：success 充值成功 / failed 充值失败 / pending 正在充值
		Message    string `json:"message"`     // 订单失败的错误描述信息
		Account    string `json:"account"`     // 链账户地址 （调用「批量购买能量值」接口不展示此字段）
		Amount     string `json:"amount"`      // 充值金额，为整数元金额；单位：分 （调用「批量购买能量值」接口不展示此字段）
		Number     string `json:"number"`      // 充值的数量，充值 gas 该值单位为 ugas，充值业务费单位为分（调用「批量购买能量值」接口不展示此字段）
		CreateTime string `json:"create_time"` // 创建时间（UTC 时间）
		UpdateTime string `json:"update_time"` // 最后操作时间（UTC 时间）
		OrderType  string `json:"order_type"`  // 订单类型，gas / business
	} `json:"data"`
}

// BatchCreateOrderReq 批量购买能量值接口请求参数
type (
	BatchCreateOrderReq struct {
		List    []List `json:"list"`     // 充值信息
		OrderId string `json:"order_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
	}
	List struct {
		Account string `json:"account"` // 链账户地址
		Amount  int    `json:"amount"`  // 购买金额 ，只能购买整数元金额；单位：分
	}
)
