package models

const (
	CreateOrder      = "/v1beta1/orders"       // 购买能量值/业务费接口
	GetOrder         = "/v1beta1/orders/%s"    // 查询能量值/业务费购买结果接口
	GetOrders        = "/v1beta1/orders"       // 查询能量值/业务费购买结果列表接口
	CreateBatchOrder = "/v1beta1/orders/batch" // 批量购买能量值接口
)

// CreateOrderReq 购买能量值/业务费接口请求参数
type CreateOrderReq struct {
	Account   string `json:"account"`
	Amount    int    `json:"amount"`
	OrderType string `json:"order_type"`
	OrderId   string `json:"order_id"`
}

// OrderRes 购买能量值/业务费接口返回值/批量购买能量值接口返回值
type OrderRes struct {
	BaseRes
	Data struct {
		OrderId string `json:"order_id"`
	} `json:"data"`
}

// GetOrdersReq 查询能量值/业务费购买结果列表接口请求参数
type GetOrdersReq struct {
	Offset    string `json:"offset,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Status    string `json:"status,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	SortBy    string `json:"sort_by,omitempty"`
}

// GetOrdersRes 查询能量值/业务费购买结果列表接口返回值
type GetOrdersRes struct {
	BaseRes
	Data struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		TotalCount int `json:"total_count"`
		OrderInfos []struct {
			OrderId    string `json:"order_id"`
			Status     string `json:"status"`
			Message    string `json:"message"`
			Account    string `json:"account"`
			Amount     string `json:"amount"`
			Number     string `json:"number"`
			CreateTime string `json:"create_time"`
			UpdateTime string `json:"update_time"`
			OrderType  string `json:"order_type"`
		} `json:"order_infos"`
	} `json:"data"`
}

// GetOrderRes 查询能量值/业务费购买结果接口返回值
type GetOrderRes struct {
	BaseRes
	Data struct {
		OrderId    string `json:"order_id"`
		Status     string `json:"status"`
		Message    string `json:"message"`
		Account    string `json:"account"`
		Amount     string `json:"amount"`
		Number     string `json:"number"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
		OrderType  string `json:"order_type"`
	} `json:"data"`
}

// CreateBatchOrderReq 批量购买能量值接口请求参数
type (
	CreateBatchOrderReq struct {
		List    []List `json:"list"`
		OrderId string `json:"order_id"`
	}
	List struct {
		Account string `json:"account"`
		Amount  int    `json:"amount"`
	}
)
