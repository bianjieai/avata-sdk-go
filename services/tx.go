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
	QueryTxResult(operationID string) (*models.QueryTxResultRes, models.Error)                       // 上链交易结果查询
	QueryTxQueueInfo(params *models.QueryTxQueueInfoReq) (*models.QueryTxQueueInfoRes, models.Error) // 上链交易排队状态查询
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

// QueryTxQueueInfo 上链交易结果查询
func (t txService) QueryTxQueueInfo(params *models.QueryTxQueueInfoReq) (*models.QueryTxQueueInfoRes, models.Error) {
	log := t.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryTxQueueInfo",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryTxQueueInfo start")

	nilRes := &models.QueryTxQueueInfoRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryTxQueueInfo Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := t.HttpClient.DoHttpRequest(http.MethodGet, models.QueryTxQueueInfo, nil, bytesData)
	log.Debugf("QueryTxQueueInfo body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryTxQueueInfo DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryTxQueueInfoRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryTxQueueInfo Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryTxQueueInfo end")
	return result, nil
}
