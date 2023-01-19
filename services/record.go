package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// RecordService 存证接口
type RecordService interface {
	CreateRecord(params *models.CreateRecordReq) *models.Response //数字作品存证接口
}

type recordService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewRecordService(log *logrus.Logger, client utils.HttpClient) *recordService {
	return &recordService{
		Logger:     log,
		HttpClient: client,
	}
}

// CreateRecord 数字作品存证接口
func (r recordService) CreateRecord(params *models.CreateRecordReq) *models.Response {
	log := r.Logger.WithFields(map[string]interface{}{
		"module":   "Record",
		"function": "CreateRecord",
		"params":   params,
	})
	log.Info("CreateRecord start")

	result := &models.Response{}

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

	body, result := r.HttpClient.DoHttpRequest(http.MethodPost, models.CreateRecord, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("CreateRecord end")
	return result
}
