package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/v3/models"
	"github.com/bianjieai/avata-sdk-go/v3/utils"
	"github.com/siddontang/go-log/loggers"
)

type ResolvesService interface {
	SetResolves(params *models.SetResolvesReq, owner, name string) (*models.TxRes, models.Error)         // 设置域名解析
	QueryResolves(params *models.QueryResolvesReq, name string) (*models.QueryResolvesRes, models.Error) // 查询域名解析
	SetReverseResolves(params *models.SetReverseResolvesReq, owner string) (*models.TxRes, models.Error) // 设置域名反向解析
	QueryReverseResolves(owner string) (*models.QueryReverseResolvesRes, models.Error)                   // 查询域名反向解析
}

type resolvesService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewResolvesService(log loggers.Advanced, client utils.HttpClient) *resolvesService {
	return &resolvesService{
		Logger:     log,
		HttpClient: client,
	}
}

/**
 * @description: 设置域名解析
 * @param {*models.SetResolvesReq} params
 * @param {*} owner ：域名拥有者（链账户）
 * @param {string} name ：域名
 * @return {*}
 */
func (r resolvesService) SetResolves(params *models.SetResolvesReq, owner, name string) (*models.TxRes, models.Error) {
	log := r.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "SetResolves",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("SetResolves start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.ResolveType != 1 && params.ResolveType != 2 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "resolve_type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "resolve_type"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("SetResolves Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := r.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.SetResolves, owner, name), bytesData, nil)
	log.Debugf("SetResolves body: %s", string(body))
	if errorRes != nil {
		log.Errorf("SetResolves DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}
	result := &models.TxRes{}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("SetResolves Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("SetResolves end")
	return result, nil
}

/**
 * @description: 查询域名解析
 * @param {*models.QueryResolvesReq} params
 * @param {string} name ：域名
 * @return {*}
 */
func (r resolvesService) QueryResolves(params *models.QueryResolvesReq, name string) (*models.QueryResolvesRes, models.Error) {
	log := r.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "QueryResolves",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryResolves start")
	nilRes := &models.QueryResolvesRes{}
	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryResolves Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := r.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryResolves, name), nil, bytesData)
	log.Debugf("QueryResolves body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryResolves DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}
	result := &models.QueryResolvesRes{}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryResolves Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryResolves end")
	return result, nil
}

/**
 * @description: 设置域名反解析
 * @param {*models.SetReverseResolvesReq} params
 * @param {string} owner ：域名拥有者（链账户）
 * @return {*}
 */
func (r resolvesService) SetReverseResolves(params *models.SetReverseResolvesReq, owner string) (*models.TxRes, models.Error) {
	log := r.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "SetReverseResolves",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("SetReverseResolves start")
	nilRes := &models.TxRes{}
	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("SetReverseResolves Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := r.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.SetReverseResolves, owner), bytesData, nil)
	log.Debugf("SetReverseResolves body: %s", string(body))
	if errorRes != nil {
		log.Errorf("SetReverseResolves DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}
	result := &models.TxRes{}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("SetReverseResolves Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("SetReverseResolves end")
	return result, nil
}

/**
 * @description: 查询域名反解析
 * @param {string} owner ：域名拥有者（链账户）
 * @return {*}
 */
func (r resolvesService) QueryReverseResolves(owner string) (*models.QueryReverseResolvesRes, models.Error) {
	log := r.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "QueryReverseResolves",
	})
	log.Info("QueryReverseResolves start")
	nilRes := &models.QueryReverseResolvesRes{}

	body, errorRes := r.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryReverseResolves, owner), nil, nil)
	log.Debugf("QueryReverseResolves body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryReverseResolves DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}
	result := &models.QueryReverseResolvesRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryReverseResolves Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}
	log.Info("QueryReverseResolves end")
	return result, nil
}
