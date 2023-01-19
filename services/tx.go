package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// TxService 交易结果查询接口
type TxService interface {
	QueryTxResult(operationID string) *models.Response                    // 上链交易结果查询
	QueryTxQueueInfo(params *models.QueryTxQueueInfoReq) *models.Response // 上链交易排队状态查询
}

type txService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewTxService(log *logrus.Logger, client utils.HttpClient) *txService {
	return &txService{
		Logger:     log,
		HttpClient: client,
	}
}

// QueryTxResult 上链交易结果查询
func (t txService) QueryTxResult(operationID string) *models.Response {
	log := t.Logger.WithFields(map[string]interface{}{
		"module":      "Tx",
		"function":    "QueryTxResult",
		"operationID": operationID,
	})
	log.Info("QueryTxResult start")

	result := &models.Response{}

	//校验必填参数
	if operationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	body, result := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryTxResult, operationID), nil, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}
	log.Info("QueryTxResult end")
	return result
}

// QueryTxQueueInfo 上链交易结果查询
func (t txService) QueryTxQueueInfo(params *models.QueryTxQueueInfoReq) *models.Response {
	log := t.Logger.WithFields(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryTxQueueInfo",
		"params":   params,
	})
	log.Info("QueryTxQueueInfo start")

	result := &models.Response{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := t.HttpClient.DoHttpRequest(http.MethodGet, models.QueryTxQueueInfo, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryTxQueueInfo end")
	return result
}
