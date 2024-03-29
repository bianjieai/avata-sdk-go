package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/v3/models"
	"github.com/bianjieai/avata-sdk-go/v3/utils"
)

// TxService 交易结果查询接口
type TxService interface {
	QueryTxResult(operationID string) (*models.QueryTxResultRes, models.Error)             // 上链交易结果查询
	QueryNativeTxResult(operationID string) (*models.QueryNativeTxResultRes, models.Error) // 上链交易结果查询（原生模块）
	QueryTxTypes() (*models.QueryTxTypesRes, models.Error)                                 // 查询枚举值列表
	QueryNativeTxTypes() (*models.QueryNativeTxTypesRes, models.Error)                     // 查询枚举值列表（原生模块）
}

type txService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewTxService(log loggers.Advanced, client utils.HttpClient) *txService {
	return &txService{
		Logger:     log,
		HttpClient: client,
	}
}

/**
 * @description: 上链交易结果查询（原生模块）
 * @param {string} operationID ：操作 ID
 * @return {*}
 */
func (t txService) QueryNativeTxResult(operationID string) (*models.QueryNativeTxResultRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":      "Tx",
		"function":    "QueryNativeTxResult",
		"operationID": operationID,
	})

	log.Info("QueryNativeTxResult start")

	nilRes := &models.QueryNativeTxResultRes{}

	//校验必填参数
	if operationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNativeTxResult, operationID), nil, nil)
	log.Debugf("QueryNativeTxResult body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeTxResult DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeTxResultRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeTxResult Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeTxResult end")
	return result, nil
}

/**
 * @description: 上链交易结果查询
 * @param {string} operationID ：操作 ID
 * @return {*}
 */
func (t txService) QueryTxResult(operationID string) (*models.QueryTxResultRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":      "Tx",
		"function":    "QueryTxResult",
		"operationID": operationID,
	})

	log.Info("QueryTxResult start")

	nilRes := &models.QueryTxResultRes{}

	//校验必填参数
	if operationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryTxResult, operationID), nil, nil)
	log.Debugf("QueryTxResult body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryTxResult DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryTxResultRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryTxResult Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryTxResult end")
	return result, nil
}

/**
 * @description: 枚举值列表查询
 * @return {*}
 */
func (t txService) QueryTxTypes() (*models.QueryTxTypesRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryTxTypes",
	})

	log.Info("QueryTxType start")

	nilRes := &models.QueryTxTypesRes{}
	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryTxTypes), nil, nil)
	log.Debugf("QueryTxTypes body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryTxTypes DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryTxTypesRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryTxTypes Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryTxTypes end")
	return result, nil
}

/**
 * @description: 枚举值列表查询（原生模块）
 * @return {*}
 */
func (t txService) QueryNativeTxTypes() (*models.QueryNativeTxTypesRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryNativeTxTypes",
	})

	log.Info("QueryNativeTxTypes start")

	nilRes := &models.QueryNativeTxTypesRes{}
	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNativeTxTypes), nil, nil)
	log.Debugf("QueryNativeTxTypes body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeTxTypes DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeTxTypesRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeTxTypes Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeTxTypes end")
	return result, nil
}
