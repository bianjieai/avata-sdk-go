package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// OrderService 充值接口
type OrderService interface {
	CreateOrder(params *models.CreateOrderReq) (*models.TxRes, models.Error)           // 购买能量值/业务费
	QueryOrders(params *models.QueryOrdersReq) (*models.QueryOrdersRes, models.Error)  // 查询能量值/业务费购买结果列表
	QueryOrder(orderID string) (*models.QueryOrderRes, models.Error)                   // 查询能量值/业务费购买结果
	BatchCreateOrder(params *models.BatchCreateOrderReq) (*models.TxRes, models.Error) // 批量购买能量值
}

type orderService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewOrderService(log loggers.Advanced, httpClient utils.HttpClient) *orderService {
	return &orderService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateOrder 购买能量值/业务费接口
func (o orderService) CreateOrder(params *models.CreateOrderReq) (*models.TxRes, models.Error) {
	log := o.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Order",
		"function": "CreateOrder",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateOrder start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OrderId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "order_id"))
	}
	if params.OrderType == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "order_type"))
	}
	if params.Account == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "account"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "account"))
	}
	if params.Amount < 100 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "amount"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "amount"))
	}
	if params.Amount%100 != 0 {
		log.Debugln(models.ErrAmount)
		return nilRes, models.InvalidParam(models.ErrAmount)
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("CreateOrder Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := o.HttpClient.DoHttpRequest(http.MethodPost, models.CreateOrder, bytesData, nil)
	log.Debugf("CreateOrder body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateOrder DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateOrder Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateOrder end")
	return result, nil
}

// QueryOrders 查询能量值/业务费购买结果列表接口
func (o orderService) QueryOrders(params *models.QueryOrdersReq) (*models.QueryOrdersRes, models.Error) {
	log := o.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Order",
		"function": "QueryOrders",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryOrders start")

	nilRes := &models.QueryOrdersRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryOrders Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := o.HttpClient.DoHttpRequest(http.MethodGet, models.QueryOrders, nil, bytesData)
	log.Debugf("QueryOrders body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryOrders DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryOrdersRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryOrders Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryOrders end")
	return result, nil
}

// QueryOrder 查询能量值/业务费购买结果接口
func (o orderService) QueryOrder(orderID string) (*models.QueryOrderRes, models.Error) {
	log := o.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Order",
		"function": "QueryOrder",
		"orderID":  orderID,
	})
	log.Info("QueryOrder start")

	nilRes := &models.QueryOrderRes{}

	// 校验必填参数
	if orderID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "order_id"))
	}

	body, errorRes := o.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryOrder, orderID), nil, nil)
	log.Debugf("QueryOrder body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryOrder DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryOrderRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryOrder Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryOrder end")
	return result, nil
}

// BatchCreateOrder 批量购买能量值接口
func (o orderService) BatchCreateOrder(params *models.BatchCreateOrderReq) (*models.TxRes, models.Error) {
	log := o.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Order",
		"function": "BatchCreateOrder",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("BatchCreateOrder start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.List == nil || len(params.List) < 1 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "list"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "list"))
	}
	if params.OrderId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "order_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BatchCreateOrder Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := o.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateOrder, bytesData, nil)
	log.Debugf("BatchCreateOrder body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BatchCreateOrder DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BatchCreateOrder Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BatchCreateOrder end")
	return result, nil
}
