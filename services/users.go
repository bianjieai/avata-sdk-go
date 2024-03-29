package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/v3/models"
	"github.com/bianjieai/avata-sdk-go/v3/utils"
	"github.com/siddontang/go-log/loggers"
)

type UsersService interface {
	CreateUsers(params *models.CreateUsersReq) (*models.CreateUsersRes, models.Error) // 创建钱包用户
	KycUsers(params *models.KycUsersReq) (*models.TxRes, models.Error)                //认证钱包用户
	EditUsers(params *models.EditUsersReq) (*models.TxRes, models.Error)              // 更新钱包用户
	QueryUsers(params *models.QueryUsersReq) (*models.QueryUsersRes, models.Error)    // 查询钱包用户信息
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

/**
 * @description: 创建钱包用户
 * @param {*models.CreateUsersReq} params
 * @return {*}
 */
func (a usersService) CreateUsers(params *models.CreateUsersReq) (*models.CreateUsersRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Users",
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
	if params.PhoneNum == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "phone_num"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "phone_num"))
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

/**
 * @description: 认证钱包用户
 * @param {*models.EditUsersReq} params
 * @return {*}
 */
func (a usersService) KycUsers(params *models.KycUsersReq) (*models.TxRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Users",
		"function": "KycUsers",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("KycUsers start")
	nilRes := &models.TxRes{}
	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.UserType != 1 && params.UserType != 2 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "User_type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "User_type"))
	}
	if params.UserID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "user_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "user_id"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("EditUsers Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := a.HttpClient.DoHttpRequest(http.MethodPost, models.KycUsers, bytesData, nil)
	log.Debugf("KycUsers body: %s", string(body))
	if errorRes != nil {
		log.Errorf("KycUsers DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}
	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("KycUsers Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("KycUsers end")
	return result, nil
}

/**
 * @description: 更新钱包用户
 * @param {*models.EditUsersReq} params
 * @return {*}
 */
func (a usersService) EditUsers(params *models.EditUsersReq) (*models.TxRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Users",
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

/**
 * @description: 查询钱包用户信息
 * @param {*models.QueryUsersReq} params
 * @return {*}
 */
func (a usersService) QueryUsers(params *models.QueryUsersReq) (*models.QueryUsersRes, models.Error) {
	log := a.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Users",
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
	if params.PhoneNum == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "phone_num"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "phone_num"))
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
