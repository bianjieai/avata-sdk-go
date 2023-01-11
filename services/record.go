package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
	"github.com/sirupsen/logrus"
)

type RecordService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}

// CreateRecord 数字作品存证接口
func (r RecordService) CreateRecord(params *models.CreateRecordReq) *models.TxRes {
	log := r.Logger.WithFields(map[string]interface{}{
		"module":   "Record",
		"function": "CreateRecord",
		"params":   params,
	})
	log.Info("CreateRecord start")

	result := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "name")
		return result
	}
	if params.Type == 0 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "type"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "type")
		return result
	}
	if params.Description == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "description"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "description")
		return result
	}
	if params.HashType == 0 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "hash_type"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "hash_type")
		return result
	}
	if params.Hash == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "hash"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "hash")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := r.HttpClient.DoHttpRequest(http.MethodPost, models.CreateRecord, bytesData, nil)
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

	log.Info("CreateRecord end")
	return result
}
