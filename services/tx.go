package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
)

// TxService 交易结果查询接口
type TxService interface {
	QueryTxResult(operationID string) *models.QueryTxResultRes                       // 上链交易结果查询
	QueryTxQueueInfo(params *models.QueryTxQueueInfoReq) *models.QueryTxQueueInfoRes // 上链交易排队状态查询
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
func (t txService) QueryTxResult(operationID string) *models.QueryTxResultRes {
	log := t.Logger.WithFields(map[string]interface{}{
		"module":      "Tx",
		"function":    "QueryTxResult",
		"operationID": operationID,
	})
	log.Info("QueryTxResult start")

	result := &models.QueryTxResultRes{}

	//校验必填参数
	if operationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	body, baseRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryTxResult, operationID), nil, nil)
	if baseRes.Message != "" {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Debugln("body: ", string(body))

	result.BaseRes = baseRes
	// 请求成功
	if body != nil {
		if err := json.Unmarshal(body, &result.Data); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return result
		}
	}

	log.Info("QueryTxResult end")
	return result
}

// QueryTxQueueInfo 上链交易结果查询
func (t txService) QueryTxQueueInfo(params *models.QueryTxQueueInfoReq) *models.QueryTxQueueInfoRes {
	log := t.Logger.WithFields(map[string]interface{}{
		"module":   "Tx",
		"function": "QueryTxQueueInfo",
		"params":   params,
	})
	log.Info("QueryTxQueueInfo start")

	result := &models.QueryTxQueueInfoRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := t.HttpClient.DoHttpRequest(http.MethodGet, models.QueryTxQueueInfo, nil, bytesData)
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

	log.Info("QueryTxQueueInfo end")
	return result
}
