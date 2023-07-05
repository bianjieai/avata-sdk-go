package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
	"github.com/siddontang/go-log/loggers"
)

type NSService interface {
	RegisterDomain(params *models.RegisterDomainReq) (*models.TxRes, models.Error)                     //注册域名
	QueryDomain(params *models.QueryDomainReq) (*models.QueryDomainRes, models.Error)                  //查询域名
	TransferDomin(params *models.TransferDominReq, owner, name string) (*models.TxRes, models.Error)   //转让域名
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

// 注册域名
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

// 查询域名
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

// 转让域名
func (n nsService) TransferDomin(params *models.TransferDominReq, owner, name string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "ns",
		"function": "TransferDomin",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("TransferDominReq start")

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
		log.Errorf("TransferDomin Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferDomin, owner, name), bytesData, nil)
	log.Debugf("TransferDomin body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferDomin DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferDomin Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferDomin end")
	return result, nil
}

// 查询用户域名
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
