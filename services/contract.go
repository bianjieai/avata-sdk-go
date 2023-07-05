package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
	"github.com/siddontang/go-log/loggers"
)

type ContractService interface {
	UseContract(params *models.UseContractReq) (*models.TxRes, models.Error)
	QueryContract(params *models.QueryContractReq) (*models.QueryContractRes, models.Error)
}
type contractService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewContractService(log loggers.Advanced, httpClient utils.HttpClient) *contractService {
	return &contractService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// 调用合约示例
func (c contractService) UseContract(params *models.UseContractReq) (*models.TxRes, models.Error) {
	log := c.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "UseContract",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("UseContract start")

	nilRes := &models.TxRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.To == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "to"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "to"))
	}
	if params.Data == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "data"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "data"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}
	if params.From == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "from"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "from"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("UseContract Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := c.HttpClient.DoHttpRequest(http.MethodPost, models.UseContract, bytesData, nil)
	log.Debugf("UseContract body: %s", string(body))
	if errorRes != nil {
		log.Errorf("UseContract DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("UseContract Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("UseContract end")
	return result, nil
}

// 查询合约示例
func (c contractService) QueryContract(params *models.QueryContractReq) (*models.QueryContractRes, models.Error) {
	log := c.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryContract",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryContract start")

	nilRes := &models.QueryContractRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryContract Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := c.HttpClient.DoHttpRequest(http.MethodGet, models.QueryContract, nil, bytesData)
	log.Debugf("QueryContract body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryContract DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryContractRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryContract Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryContract end")
	return result, nil
}
