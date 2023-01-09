package services

import (
	"encoding/json"
	"net/http"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	*logrus.Logger    // 日志
	models.BaseParams // 域名和项目参数
}

// CreateAccount 创建链账户
func (a AccountService) CreateAccount(params *models.CreateAccountReq) (*models.CreateAccountRes, *models.Error) {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "CreateAccount",
		"Params":   params,
	})
	log.Info("CreateAccount start")

	nilRes := &models.CreateAccountRes{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}

	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.CreateAccount, a.BaseParams, bytesData, nil)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}

	log.Debugln("body: ", string(body))

	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}

	// 请求成功
	var result *models.CreateAccountRes
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}

	log.Info("CreateAccount end")
	return result, nil
}

// BatchCreateAccounts 批量创建链账户
func (a AccountService) BatchCreateAccounts(params *models.BatchCreateAccountsReq) (*models.BatchCreateAccountsRes, *models.Error) {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "BatchCreateAccounts",
		"Params":   params,
	})
	log.Info("BatchCreateAccounts start")

	nilRes := &models.BatchCreateAccountsRes{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}

	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.BatchCreateAccounts, a.BaseParams, bytesData, nil)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}

	log.Debugln("body: ", string(body))

	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}

	// 请求成功
	var result *models.BatchCreateAccountsRes
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}

	log.Info("BatchCreateAccounts end")
	return result, nil
}

// GetAccounts 查询链账户
func (a AccountService) GetAccounts(params *models.GetAccountsReq) (*models.GetAccountsRes, *models.Error) {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "GetAccounts",
		"Params":   params,
	})
	log.Info("GetAccounts start")

	nilRes := &models.GetAccountsRes{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}

	statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.GetAccounts, a.BaseParams, nil, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}

	log.Debugln("body: ", string(body))

	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}

	// 请求成功
	var result *models.GetAccountsRes
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}

	log.Info("GetAccounts end")
	return result, nil
}

// GetAccountsHistory 查询链账户操作记录
func (a AccountService) GetAccountsHistory(params *models.GetAccountsHistoryReq) (*models.GetAccountsHistoryRes, *models.Error) {
	log := a.Logger.WithFields(map[string]interface{}{
		"module":   "Account",
		"function": "GetAccountsHistory",
		"Params":   params,
	})
	log.Info("GetAccountsHistory start")

	nilRes := &models.GetAccountsHistoryRes{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}

	statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.GetAccountsHistory, a.BaseParams, nil, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}

	log.Debugln("body: ", string(body))

	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}

	// 请求成功
	var result *models.GetAccountsHistoryRes
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}

	log.Info("GetAccountsHistory end")
	return result, nil
}
