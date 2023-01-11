package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
)

type TxService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}

// GetTxResult 上链交易结果查询
func (t TxService) GetTxResult(operationID string) *models.GetTxResultRes {
	log := t.Logger.WithFields(map[string]interface{}{
		"module":      "Tx",
		"function":    "GetTxResult",
		"operationID": operationID,
	})
	log.Info("GetTxResult start")

	result := &models.GetTxResultRes{}

	//校验必填参数
	if operationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	body, baseRes := t.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.GetTxResult, operationID), nil, nil)
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

	log.Info("GetTxResult end")
	return result
}

// GetTxQueueInfo 上链交易结果查询
func (t TxService) GetTxQueueInfo(params *models.GetTxQueueInfoReq) *models.GetTxQueueInfoRes {
	log := t.Logger.WithFields(map[string]interface{}{
		"module":   "Tx",
		"function": "GetTxQueueInfo",
		"params":   params,
	})
	log.Info("GetTxQueueInfo start")

	result := &models.GetTxQueueInfoRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := t.HttpClient.DoHttpRequest(http.MethodGet, models.GetTxQueueInfo, nil, bytesData)
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

	log.Info("GetTxQueueInfo end")
	return result
}
