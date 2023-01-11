package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
)

type OrderService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}

// CreateOrder 购买能量值/业务费接口
func (o OrderService) CreateOrder(params *models.CreateOrderReq) *models.OrderRes {
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
	if params.Amount <= 100 || params.Amount%100 != 0 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "amount"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "amount")
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

// GetOrders 查询能量值/业务费购买结果列表接口
func (o OrderService) GetOrders(params *models.GetOrdersReq) *models.GetOrdersRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "GetOrders",
		"params":   params,
	})
	log.Info("GetOrders start")

	result := &models.GetOrdersRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodGet, models.GetOrders, nil, bytesData)
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

	log.Info("GetOrders end")
	return result
}

// GetOrder 查询能量值/业务费购买结果接口
func (o OrderService) GetOrder(orderID string) *models.GetOrderRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "GetOrder",
		"orderID":  orderID,
	})
	log.Info("GetOrder start")

	result := &models.GetOrderRes{}

	// 校验必填参数
	if orderID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "order_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "order_id")
		return result
	}

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.GetOrder, orderID), nil, nil)
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

	log.Info("GetOrder end")
	return result
}

// CreateBatchOrder 批量购买能量值接口
func (o OrderService) CreateBatchOrder(params *models.CreateBatchOrderReq) *models.OrderRes {
	log := o.Logger.WithFields(map[string]interface{}{
		"module":   "Order",
		"function": "CreateBatchOrder",
		"params":   params,
	})
	log.Info("CreateBatchOrder start")

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

	body, baseRes := o.HttpClient.DoHttpRequest(http.MethodPost, models.CreateBatchOrder, bytesData, nil)
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

	log.Info("CreateBatchOrder end")
	return result
}
