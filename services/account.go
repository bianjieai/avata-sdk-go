package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// AccountService 链账户接口
type AccountService interface {
	CreateAccount(params *models.CreateAccountReq) *models.Response               // 创建链账户
	BatchCreateAccounts(params *models.BatchCreateAccountsReq) *models.Response   // 批量创建链账户
	QueryAccounts(params *models.QueryAccountsReq) *models.Response               // 查询链账户
	QueryAccountsHistory(params *models.QueryAccountsHistoryReq) *models.Response // 查询链账户操作记录
}

type accountService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewAccountService(log *logrus.Logger, httpClient utils.HttpClient) *accountService {
	return &accountService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateAccount 创建链账户
func (a accountService) CreateAccount(params *models.CreateAccountReq) *models.Response {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "CreateAccount",
		"params":   params,
	})
	log.Info("CreateAccount start")

	result := &models.Response{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "name")
		return result
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := a.HttpClient.DoHttpRequest(http.MethodPost, models.CreateAccount, bytesData, nil)

	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("CreateAccount end")
	return result
}

// BatchCreateAccounts 批量创建链账户
func (a accountService) BatchCreateAccounts(params *models.BatchCreateAccountsReq) *models.Response {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "BatchCreateAccounts",
		"params":   params,
	})
	log.Info("BatchCreateAccounts start")

	result := &models.Response{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := a.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateAccounts, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchCreateAccounts end")
	return result
}

// QueryAccounts 查询链账户
func (a accountService) QueryAccounts(params *models.QueryAccountsReq) *models.Response {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccounts",
		"params":   params,
	})
	log.Info("QueryAccounts start")

	result := &models.Response{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryAccounts, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryAccounts end")
	return result
}

// QueryAccountsHistory 查询链账户操作记录
func (a accountService) QueryAccountsHistory(params *models.QueryAccountsHistoryReq) *models.Response {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccountsHistory",
		"params":   params,
	})
	log.Info("QueryAccountsHistory start")

	result := &models.Response{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryAccountsHistory, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryAccountsHistory end")
	return result
}
