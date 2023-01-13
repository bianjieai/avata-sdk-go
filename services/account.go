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
	CreateAccount(params *models.CreateAccountReq) *models.CreateAccountRes                      // 创建链账户
	BatchCreateAccounts(params *models.BatchCreateAccountsReq) *models.BatchCreateAccountsRes    // 批量创建链账户
	QueryAccounts(params *models.QueryAccountsReq) *models.QueryAccountsRes                      // 查询链账户
	QueryAccountsHistory(params *models.QueryAccountsHistoryReq) *models.QueryAccountsHistoryRes // 查询链账户操作记录
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
func (a accountService) CreateAccount(params *models.CreateAccountReq) *models.CreateAccountRes {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "CreateAccount",
		"params":   params,
	})
	log.Info("CreateAccount start")

	result := &models.CreateAccountRes{}

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

	body, baseRes := a.HttpClient.DoHttpRequest(http.MethodPost, models.CreateAccount, bytesData, nil)

	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.Code = baseRes.Code
	result.Error = baseRes.Error
	result.Message = baseRes.Message
	result.Http = baseRes.Http

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

	log.Info("CreateAccount end")
	return result
}

// BatchCreateAccounts 批量创建链账户
func (a accountService) BatchCreateAccounts(params *models.BatchCreateAccountsReq) *models.BatchCreateAccountsRes {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "BatchCreateAccounts",
		"params":   params,
	})
	log.Info("BatchCreateAccounts start")

	result := &models.BatchCreateAccountsRes{}

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

	body, baseRes := a.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateAccounts, bytesData, nil)
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

	log.Info("BatchCreateAccounts end")
	return result
}

// QueryAccounts 查询链账户
func (a accountService) QueryAccounts(params *models.QueryAccountsReq) *models.QueryAccountsRes {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccounts",
		"params":   params,
	})
	log.Info("QueryAccounts start")

	result := &models.QueryAccountsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryAccounts, nil, bytesData)
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

	log.Info("QueryAccounts end")
	return result
}

// QueryAccountsHistory 查询链账户操作记录
func (a accountService) QueryAccountsHistory(params *models.QueryAccountsHistoryReq) *models.QueryAccountsHistoryRes {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "QueryAccountsHistory",
		"params":   params,
	})
	log.Info("QueryAccountsHistory start")

	result := &models.QueryAccountsHistoryRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryAccountsHistory, nil, bytesData)
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

	log.Info("QueryAccountsHistory end")
	return result
}
