package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
)

// OrderService 充值接口
type OrderService interface {
	CreateOrder(params *models.CreateOrderReq) *models.OrderRes           // 购买能量值/业务费
	QueryOrders(params *models.QueryOrdersReq) *models.QueryOrdersRes     // 查询能量值/业务费购买结果列表
	QueryOrder(orderID string) *models.QueryOrderRes                      // 查询能量值/业务费购买结果
	BatchCreateOrder(params *models.BatchCreateOrderReq) *models.OrderRes // 批量购买能量值
}

type orderService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewOrderService(log *logrus.Logger, httpClient utils.HttpClient) *orderService {
	return &orderService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateOrder 购买能量值/业务费接口
func (o orderService) CreateOrder(params *models.CreateOrderReq) *models.OrderRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "CreateOrder",
		"params":   params,
	})
	log.Info("CreateOrder start")

	result := &models.OrderRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OrderId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "order_id")
		return result
	}
	if params.OrderType == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_type"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "order_type")
		return result
	}
	if params.Account == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "account"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "account")
		return result
	}
	if params.Amount < 100 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "amount"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "amount")
		return result
	}
	if params.Amount%100 != 0 {
		log.Debugln(models.ErrAmount)
		result.Code = models.CodeFailed
		result.Message = models.ErrAmount
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodPost, models.CreateOrder, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == models.CodeFailed {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = models.CodeFailed
			result.Message = err.Error()
			return result
		}
	}

	log.Info("CreateOrder end")
	return result
}

// QueryOrders 查询能量值/业务费购买结果列表接口
func (o orderService) QueryOrders(params *models.QueryOrdersReq) *models.QueryOrdersRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "QueryOrders",
		"params":   params,
	})
	log.Info("QueryOrders start")

	result := &models.QueryOrdersRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodGet, models.QueryOrders, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == models.CodeFailed {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = models.CodeFailed
			result.Message = err.Error()
			return result
		}
	}

	log.Info("QueryOrders end")
	return result
}

// QueryOrder 查询能量值/业务费购买结果接口
func (o orderService) QueryOrder(orderID string) *models.QueryOrderRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "QueryOrder",
		"orderID":  orderID,
	})
	log.Info("QueryOrder start")

	result := &models.QueryOrderRes{}

	// 校验必填参数
	if orderID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "order_id")
		return result
	}

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryOrder, orderID), nil, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == models.CodeFailed {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = models.CodeFailed
			result.Message = err.Error()
			return result
		}
	}

	log.Info("QueryOrder end")
	return result
}

// BatchCreateOrder 批量购买能量值接口
func (o orderService) BatchCreateOrder(params *models.BatchCreateOrderReq) *models.OrderRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "BatchCreateOrder",
		"params":   params,
	})
	log.Info("BatchCreateOrder start")

	result := &models.OrderRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.List == nil || len(params.List) < 1 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "list"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "list")
		return result
	}
	if params.OrderId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "order_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateOrder, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == models.CodeFailed {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = models.CodeFailed
			result.Message = err.Error()
			return result
		}
	}

	log.Info("BatchCreateOrder end")
	return result
}
