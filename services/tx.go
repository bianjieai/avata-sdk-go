package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// TxService 交易结果查询接口
type TxService interface {
	QueryTxResult(operationID string) (*models.QueryTxResultRes, models.Error)             // 上链交易结果查询
	QueryNativeTxResult(operationID string) (*models.QueryNativeTxResultRes, models.Error) // 以原生方式查询上链交易结果
	QueryTxTypes() (*models.QueryTxTypesRes, models.Error)                                 // 枚举值列表查询
	QueryNativeTxTypes() (*models.QueryNativeTxTypesRes, models.Error)                     // 以原生方式查询枚举值列表
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

// QueryNativeTxResult 上链交易结果查询
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
	log.Debugf("QueryTxResult body: %s", string(body))
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

// QueryTxResult 上链交易结果查询
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

// QueryTxTypes 枚举值列表查询
func (t txService) QueryTxTypes() (*models.QueryTxTypesRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryTxResult",
	})

	log.Info("QueryTxType start")

	nilRes := &models.QueryTxTypesRes{}
	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryTxTypes), nil, nil)
	log.Debugf("QueryTxTypesRes body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryTxTypesRes DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryTxTypesRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryTxTypesRes Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryTxTypesRes end")
	return result, nil

}

// QueryNativeTxTypes 枚举值列表
func (t txService) QueryNativeTxTypes() (*models.QueryNativeTxTypesRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryNativeTxResult",
	})

	log.Info("QueryNativeTxType start")

	nilRes := &models.QueryNativeTxTypesRes{}
	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNativeTxTypes), nil, nil)
	log.Debugf("QueryNativeTxTypesRes body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeTxTypesRes DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeTxTypesRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeTxTypesRes Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeTxTypesRes end")
	return result, nil

}
