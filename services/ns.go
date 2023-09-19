package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/v3/models"
	"github.com/bianjieai/avata-sdk-go/v3/utils"
	"github.com/siddontang/go-log/loggers"
)

type NSService interface {
	RegisterDomain(params *models.RegisterDomainReq) (*models.TxRes, models.Error)                     //注册域名
	QueryDomain(params *models.QueryDomainReq) (*models.QueryDomainRes, models.Error)                  //查询域名
	TransferDomain(params *models.TransferDomainReq, owner, name string) (*models.TxRes, models.Error) //转让域名
	QueryDomains(params *models.QueryDomainsReq, owner string) (*models.QueryDomainsRes, models.Error) //查询用户域名
}

type nsService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewNSService(log loggers.Advanced, httpClient utils.HttpClient) *nsService {
	return &nsService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

/**
 * @description: 注册域名
 * @param {*models.RegisterDomainReq} params
 * @return {*}
 */
func (n nsService) RegisterDomain(params *models.RegisterDomainReq) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "RegisterDomain",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("RegisterDomain start")

	nilRes := &models.TxRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	if params.Owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("RegisterDomain Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.RegisterDomain, bytesData, nil)
	log.Debugf("RegisterDomain body: %s", string(body))
	if errorRes != nil {
		log.Errorf("RegisterDomain DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("RegisterDomain Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("RegisterDomain end")
	return result, nil
}

/**
 * @description: 查询域名
 * @param {*models.QueryDomainReq} params
 * @return {*}
 */
func (n nsService) QueryDomain(params *models.QueryDomainReq) (*models.QueryDomainRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "QueryDomain",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryDomain start")

	nilRes := &models.QueryDomainRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryDomain Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryDomain, nil, bytesData)
	log.Debugf("QueryDomain body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryDomain DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryDomainRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryDomain Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryDomain end")
	return result, nil
}

/**
 * @description: 转让域名
 * @return {*}
 */
func (n nsService) TransferDomain(params *models.TransferDomainReq, owner, name string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "TransferDomain",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("TransferDomain start")

	nilRes := &models.TxRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferDomain Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferDomain, owner, name), bytesData, nil)
	log.Debugf("TransferDomain body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferDomain DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferDomain Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferDomain end")
	return result, nil
}

/**
 * @description: 查询用户域名
 * @return {*}
 */
func (n nsService) QueryDomains(params *models.QueryDomainsReq, owner string) (*models.QueryDomainsRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "QueryDomains",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryDomains start")

	nilRes := &models.QueryDomainsRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryDomains Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryDomains, owner), nil, bytesData)
	log.Debugf("QueryDomains body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryDomains DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryDomainsRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryDomains Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryDomains end")
	return result, nil
}
