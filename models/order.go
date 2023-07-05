package models

const (
	CreateOrder      = "/v3/orders"       // 购买能量值/业务费接口
	QueryOrder       = "/v2/orders/%s"    // 查询能量值/业务费购买结果接口
	QueryOrders      = "/v3/orders"       // 查询能量值/业务费购买结果列表接口
	BatchCreateOrder = "/v3/orders/batch" // 批量购买能量值接口
)

// CreateOrderReq 购买能量值/业务费接口请求参数
type CreateOrderReq struct {
	Account     string `json:"account"`      // 链账户地址
	Amount      int    `json:"amount"`       // 购买金额 ，只能购买整数元金额；单位：分
	OrderType   int    `json:"order_type"`   // 充值类型：1 gas：能量值；2 business：业务费；Enum: 1   2
	OperationID string `json:"operation_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
}

// OrderRes 购买能量值/业务费接口返回值/批量购买能量值接口返回值
type OrderRes struct {
	Data struct {
	} `json:"data"`
}

// QueryOrdersReq 查询能量值/业务费购买结果列表接口请求参数
type QueryOrdersReq struct {
	PageKey    string `json:"page_key,omitempty"`    //分页数据的Key， Avata会根据该值进行上下页的查询， 该值请从返回的数据体中获取，首页查询可以不传该参数
	Limit      string `json:"limit,omitempty"`       // 每页记录数，默认为 10，上限为 50
	Status     string `json:"status,omitempty"`      // 订单状态：success 充值成功 / failed 充值失败 / pending 正在充值
	StartDate  string `json:"start_date,omitempty"`  // 充值订单创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate    string `json:"end_date,omitempty"`    // 充值订单创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy     string `json:"sort_by,omitempty"`     // 排序规则：DATE_ASC / DATE_DESC，默认为 DATE_DESC
	CountTotal string `json:"count_total,omitempty"` //是否查询数据的总数量0：不查询总数（默认）1：查询总数
}

// QueryOrdersRes 查询能量值/业务费购买结果列表接口返回值
type QueryOrdersRes struct {
	Data struct {
		PrevPageKey string `json:"prev_page_key"` //上一页数据的Key， Avata会根据该值进行上一页数据的查询
		NextPageKey string `json:"next_page_key"` //下一页数据的Key， Avata会根据该值进行下一页数据的查询
		Limit       int    `json:"limit"`         // 每页记录数
		TotalCount  int    `json:"total_count"`   // 总记录数
		OrderInfos  []struct {
			OperationId string `json:"operation_id"` //
			Status      string `json:"status"`       // 订单状态：success 充值成功 / failed 充值失败 / pending 正在充值
			Message     string `json:"message"`      // 订单失败的错误描述信息
			Account     string `json:"account"`      // 链账户地址 （调用「批量购买能量值」接口不展示此字段）
			Amount      string `json:"amount"`       // 充值金额，为整数元金额；单位：分 （调用「批量购买能量值」接口不展示此字段）
			Number      string `json:"number"`       // 充值的数量，充值 gas 该值单位为 ugas，充值业务费单位为分（调用「批量购买能量值」接口不展示此字段）
			CreateTime  string `json:"create_time"`  // 创建时间（UTC 时间）
			UpdateTime  string `json:"update_time"`  // 最后操作时间（UTC 时间）
			OrderType   int    `json:"order_type"`   // 订单类型，gas / business
		} `json:"order_infos"`
	} `json:"data"`
}

// QueryOrderRes 查询能量值/业务费购买结果接口返回值
type QueryOrderRes struct {
	Data struct {
		OperationID string `json:"operation_id"` // 订单流水号
		Status      int    `json:"status"`       // 订单状态：success 充值成功 / failed 充值失败 / pending 正在充值
		Message     string `json:"message"`      // 订单失败的错误描述信息
		Account     string `json:"account"`      // 链账户地址 （调用「批量购买能量值」接口不展示此字段）
		Amount      string `json:"amount"`       // 充值金额，为整数元金额；单位：分 （调用「批量购买能量值」接口不展示此字段）
		Number      string `json:"number"`       // 充值的数量，充值 gas 该值单位为 ugas，充值业务费单位为分（调用「批量购买能量值」接口不展示此字段）
		CreateTime  string `json:"create_time"`  // 创建时间（UTC 时间）
		UpdateTime  string `json:"update_time"`  // 最后操作时间（UTC 时间）
		OrderType   int    `json:"order_type"`   // 订单类型，gas / business
	} `json:"data"`
}

// BatchCreateOrderReq 批量购买能量值接口请求参数
type (
	BatchCreateOrderReq struct {
		List        []List `json:"list"`         // 充值信息
		OperationID string `json:"operation_id"` // 自定义订单流水号，必需且仅包含数字、下划线及英文字母大/小写
	}
	List struct {
		Account string `json:"account"` // 链账户地址
		Amount  int    `json:"amount"`  // 购买金额 ，只能购买整数元金额；单位：分
	}
)
