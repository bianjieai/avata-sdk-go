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

	nilRes := &models.CreateAccountRes{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err)
		return nilRes, &models.Error{Exception: err}
	}

	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.CreateAccount, a.BaseParams, bytesData)
	if err != nil {
		log.WithError(err)
		return nilRes, &models.Error{Exception: err}
	}

	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err)
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}

	var result *models.CreateAccountRes
	// 请求成功
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err)
		return nilRes, &models.Error{Exception: err}
	}
	return result, nil
}
