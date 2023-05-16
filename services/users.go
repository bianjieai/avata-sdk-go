package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
	"github.com/siddontang/go-log/loggers"
)

type UsersService interface {
	CreateUsers(params *models.CreateUsersReq) (*models.CreateUsersRes, models.Error)
	EditUsers(params *models.EditUsersReq) (*models.TxRes, models.Error)
	QueryUsers(params *models.QueryUsersReq) (*models.QueryUsersRes, models.Error)
}
type usersService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewUsersService(log loggers.Advanced, httpClient utils.HttpClient) *usersService {
	return &usersService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

//创建用户
func (a usersService) CreateUsers(params *models.CreateUsersReq) (*models.CreateUsersRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "user",
		"function": "CreateUsers",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateUsers start")

	nilRes := &models.CreateUsersRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.UserType != 1 && params.UserType != 2 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "user_type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "user_type"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	if params.PhoneNum == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "phone_num"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "phone_num"))
	}
	if params.UserType == 1 {
		if params.CertificateNum == "" {
			log.Debugln(fmt.Sprintf(models.ErrParam, "certificate_num"))
			return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "certificate_num"))
		}
	}
	if params.UserType == 2 {
		if params.RegistrationNum == "" {
			log.Debugln(fmt.Sprintf(models.ErrParam, "registration_num"))
			return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "registration_num"))
		}
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("CreateUsers Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodPost, models.CreateUsers, bytesData, nil)
	log.Debugf("CreateUsers body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateUsers DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.CreateUsersRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateUsers Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateUsers end")
	return result, nil
}

//更新用户
func (a usersService) EditUsers(params *models.EditUsersReq) (*models.TxRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "user",
		"function": "EditUsers",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("EditUsers start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.UserID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "user_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "user_id"))
	}
	if params.PhoneNum == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "phone_num"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "phone_num"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("EditUsers Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodPatch, models.EditUsers, bytesData, nil)
	log.Debugf("EditUsers body: %s", string(body))
	if errorRes != nil {
		log.Errorf("EditUsers DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("EditUsers Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("EditUsers end")
	return result, nil
}

//查询用户信息
func (a usersService) QueryUsers(params *models.QueryUsersReq) (*models.QueryUsersRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "user",
		"function": "QueryUsers",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryUsers start")

	nilRes := &models.QueryUsersRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.UserType == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "user_type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "user_type"))
	}
	if params.Code == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "code"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "code"))
	}
	bytesData, err := json.Marshal(params)

	if err != nil {
		log.Errorf("QueryUsers Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodGet, models.QueryUsers, nil, bytesData)
	log.Debugf("QueryUsers body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryUsers DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryUsersRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryUsers Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryUsers end")
	return result, nil
}
