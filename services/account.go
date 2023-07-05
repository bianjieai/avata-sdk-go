package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// AccountService 链账户接口
type AccountService interface {
	CreateAccount(params *models.CreateAccountReq) (*models.CreateAccountRes, models.Error)                      // 创建链账户
	BatchCreateAccounts(params *models.BatchCreateAccountsReq) (*models.BatchCreateAccountsRes, models.Error)    // 批量创建链账户
	QueryAccounts(params *models.QueryAccountsReq) (*models.QueryAccountsRes, models.Error)                      // 查询链账户
	QueryAccountsHistory(params *models.QueryAccountsHistoryReq) (*models.QueryAccountsHistoryRes, models.Error) // 查询链账户操作记录
	QueryNativeAccountsHistory(params *models.QueryNativeAccountsHistoryReq) (*models.QueryNativeAccountsHistoryRes, models.Error)
}

type accountService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewAccountService(log loggers.Advanced, httpClient utils.HttpClient) *accountService {
	return &accountService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateAccount 创建链账户
func (a accountService) CreateAccount(params *models.CreateAccountReq) (*models.CreateAccountRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Account",
		"function": "CreateAccount",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateAccount start")

	nilRes := &models.CreateAccountRes{}

	// 校验必填参数
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
		log.Errorf("CreateAccount Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodPost, models.CreateAccount, bytesData, nil)
	log.Debugf("CreateAccount body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateAccount DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.CreateAccountRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateAccount Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateAccount end")
	return result, nil
}

// BatchCreateAccounts 批量创建链账户
func (a accountService) BatchCreateAccounts(params *models.BatchCreateAccountsReq) (*models.BatchCreateAccountsRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Account",
		"function": "BatchCreateAccounts",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("BatchCreateAccounts start")

	nilRes := &models.BatchCreateAccountsRes{}

	// 校验必填参数
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
		log.Errorf("BatchCreateAccounts Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateAccounts, bytesData, nil)
	log.Debugf("BatchCreateAccounts body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BatchCreateAccounts DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.BatchCreateAccountsRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BatchCreateAccounts Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BatchCreateAccounts end")
	return result, nil
}

// QueryAccounts 查询链账户
func (a accountService) QueryAccounts(params *models.QueryAccountsReq) (*models.QueryAccountsRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccounts",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryAccounts start")

	nilRes := &models.QueryAccountsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryAccounts Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryAccounts, nil, bytesData)
	log.Debugf("QueryAccounts body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryAccounts DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryAccountsRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryAccounts Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryAccounts end")
	return result, nil
}

// QueryAccountsHistory 查询链账户操作记录
func (a accountService) QueryAccountsHistory(params *models.QueryAccountsHistoryReq) (*models.QueryAccountsHistoryRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccountsHistory",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryAccountsHistory start")

	nilRes := &models.QueryAccountsHistoryRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryAccountsHistory Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryAccountsHistory, nil, bytesData)
	log.Debugf("QueryAccountsHistory body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryAccountsHistory DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryAccountsHistoryRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryAccountsHistory Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryAccountsHistory end")
	return result, nil
}

// QueryNativeAccountsHistory 以原生方式查询链账户操作记录
func (a accountService) QueryNativeAccountsHistory(params *models.QueryNativeAccountsHistoryReq) (*models.QueryNativeAccountsHistoryRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccountsHistory",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryAccountsHistory start")

	nilRes := &models.QueryNativeAccountsHistoryRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNativeAccountsHistory Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNativeAccountsHistory, nil, bytesData)
	log.Debugf("QueryNativeAccountsHistory body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeAccountsHistory DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeAccountsHistoryRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeAccountsHistory Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeAccountsHistory end")
	return result, nil
}
