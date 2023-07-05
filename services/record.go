package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// RecordService 存证接口
type RecordService interface {
	CreateRecord(params *models.CreateRecordReq) (*models.TxRes, models.Error) //数字作品存证接口
}

type recordService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewRecordService(log loggers.Advanced, client utils.HttpClient) *recordService {
	return &recordService{
		Logger:     log,
		HttpClient: client,
	}
}

// CreateRecord 数字作品存证接口
func (r recordService) CreateRecord(params *models.CreateRecordReq) (*models.TxRes, models.Error) {
	log := r.Logger
	log.Debugln(map[string]interface{}{
		"module":   "Record",
		"function": "CreateRecord",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateRecord start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	if params.Type == 0 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "type"))
	}
	if params.Description == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "description"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "description"))
	}
	if params.HashType == 0 {
		log.Debugln(fmt.Sprintf(models.ErrParam, "hash_type"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "hash_type"))
	}
	if params.Hash == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "hash"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "hash"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("CreateRecord Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := r.HttpClient.DoHttpRequest(http.MethodPost, models.CreateRecord, bytesData, nil)
	log.Debugf("CreateRecord body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateRecord DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}
	result := &models.TxRes{}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateRecord Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateRecord end")
	return result, nil
}
