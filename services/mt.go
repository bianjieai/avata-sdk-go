package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// MTService MT 接口
type MTService interface {
	CreateMTClass(params *models.CreateMTClassReq) (*models.TxRes, models.Error)                                        // 创建 MT 类别
	QueryMTClasses(params *models.QueryMTClassesReq) (*models.QueryMTClassesRes, models.Error)                          // 查询 MT 类别
	QueryMTClass(id string) (*models.QueryMTClassRes, models.Error)                                                     // 查询 MT 类别详情
	TransferMTClass(classID, owner string, params *models.TransferMTClassReq) (*models.TxRes, models.Error)             // 转让 MT 类别
	IssueMT(classID string, params *models.IssueMTReq) (*models.TxRes, models.Error)                                    // 发行 MT
	MintMT(classID, mtID string, params *models.MintMTReq) (*models.TxRes, models.Error)                                // 增发 MT
	TransferMT(classID, owner, mtID string, params *models.TransferMTReq) (*models.TxRes, models.Error)                 // 转让 MT
	EditMT(classID, owner, mtID string, params *models.EditMTReq) (*models.TxRes, models.Error)                         // 编辑 MT
	BurnMT(classID, owner, mtID string, params *models.BurnMTReq) (*models.TxRes, models.Error)                         // 销毁 MT
	QueryMTs(params *models.QueryMTsReq) (*models.QueryMTsRes, models.Error)                                            // 查询 MT
	QueryMT(classID, mtID string) (*models.QueryMTRes, models.Error)                                                    // 查询 MT 详情
	QueryMTHistory(classID, mtID string, params *models.QueryMTHistoryReq) (*models.QueryMTHistoryRes, models.Error)    // 查询 MT 操作记录
	QueryMTBalance(classID, account string, params *models.QueryMTBalanceReq) (*models.QueryMTBalanceRes, models.Error) // 查询 MT 余额
}

type mtService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewMTService(log loggers.Advanced, httpClient utils.HttpClient) *mtService {
	return &mtService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateMTClass 创建 MT 类别
func (m mtService) CreateMTClass(params *models.CreateMTClassReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "CreateMTClass",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateMTClass start")

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
	if params.Owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("CreateMTClass Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodPost, models.CreateMTClass, bytesData, nil)
	log.Debugf("CreateMTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateMTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateMTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateMTClass end")
	return result, nil
}

// QueryMTClasses 查询 MT 类别
func (m mtService) QueryMTClasses(params *models.QueryMTClassesReq) (*models.QueryMTClassesRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTClasses",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryMTClasses start")

	nilRes := &models.QueryMTClassesRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryMTClasses Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodGet, models.QueryMTClasses, nil, bytesData)
	log.Debugf("QueryMTClasses body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryMTClasses DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryMTClassesRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryMTClasses Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryMTClasses end")
	return result, nil
}

// QueryMTClass 查询 MT 类别详情
func (m mtService) QueryMTClass(id string) (*models.QueryMTClassRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTClass",
		"id":       id,
	})
	log.Info("QueryMTClass start")

	nilRes := &models.QueryMTClassRes{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "id"))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMTClass, id), nil, nil)
	log.Debugf("QueryMTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryMTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryMTClassRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryMTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryMTClass end")
	return result, nil
}

// TransferMTClass 转让 MT 类别
func (m mtService) TransferMTClass(classID, owner string, params *models.TransferMTClassReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "TransferMTClass",
		"classID":  classID,
		"owner":    owner,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("TransferMTClass start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "recipient"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferMTClass Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferMTClass, classID, owner), bytesData, nil)
	log.Debugf("TransferMTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferMTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferMTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferMTClass end")
	return result, nil
}

// IssueMT 发行 MT
func (m mtService) IssueMT(classID string, params *models.IssueMTReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "IssueMT",
		"classID":  classID,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("IssueMT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("IssueMT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.IssueMT, classID), bytesData, nil)
	log.Debugf("IssueMT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("IssueMT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("IssueMT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("IssueMT end")
	return result, nil
}

// MintMT 增发 MT
func (m mtService) MintMT(classID, mtID string, params *models.MintMTReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "MintMT",
		"classID":  classID,
		"mtID":     mtID,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("MintMT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "mt_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("MintMT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.MintMT, classID, mtID), bytesData, nil)
	log.Debugf("MintMT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("MintMT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("MintMT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("MintMT end")
	return result, nil
}

// TransferMT 转让 MT
func (m mtService) TransferMT(classID, owner, mtID string, params *models.TransferMTReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "TransferMT",
		"classID":  classID,
		"owner":    owner,
		"mtID":     mtID,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("TransferMT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "mt_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "recipient"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferMT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferMT, classID, owner, mtID), bytesData, nil)
	log.Debugf("TransferMT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferMT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferMT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferMT end")
	return result, nil
}

// EditMT 编辑 MT
func (m mtService) EditMT(classID, owner, mtID string, params *models.EditMTReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "EditMT",
		"classID":  classID,
		"owner":    owner,
		"mtID":     mtID,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("EditMT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "mt_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}

	if params.Data == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "data"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "data"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("EditMT Marshal Params: %s", err.Error())
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.EditMT, classID, owner, mtID), bytesData, nil)
	log.Debugf("EditMT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("EditMT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("EditMT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("EditMT end")
	return result, nil
}

// BurnMT 销毁 MT
func (m mtService) BurnMT(classID, owner, mtID string, params *models.BurnMTReq) (*models.TxRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "BurnMT",
		"classID":  classID,
		"owner":    owner,
		"mtID":     mtID,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("BurnMT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "mt_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BurnMT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BurnMT, classID, owner, mtID), bytesData, nil)
	log.Debugf("BurnMT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BurnMT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BurnMT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BurnMT end")
	return result, nil
}

// QueryMTs 查询 MT
func (m mtService) QueryMTs(params *models.QueryMTsReq) (*models.QueryMTsRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTs",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryMTs start")

	nilRes := &models.QueryMTsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryMTs Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodGet, models.QueryMTs, nil, bytesData)
	log.Debugf("QueryMTs body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryMTs DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryMTsRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryMTs Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryMTs end")
	return result, nil
}

// QueryMT 查询 MT 详情
func (m mtService) QueryMT(classID, mtID string) (*models.QueryMTRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMT",
		"classID":  classID,
		"mtID":     mtID,
	})
	log.Info("QueryMT start")

	nilRes := &models.QueryMTRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))

	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "mt_id"))

	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMT, classID, mtID), nil, nil)
	log.Debugf("QueryMT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryMT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryMTRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryMT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryMT end")
	return result, nil
}

// QueryMTHistory 查询 MT 操作记录
func (m mtService) QueryMTHistory(classID, mtID string, params *models.QueryMTHistoryReq) (*models.QueryMTHistoryRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTHistory",
		"classID":  classID,
		"mtID":     mtID,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryMTHistory start")

	nilRes := &models.QueryMTHistoryRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))

	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "mt_id"))

	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryMTHistory Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMTHistory, classID, mtID), nil, bytesData)
	log.Debugf("QueryMTHistory body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryMTHistory DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryMTHistoryRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryMTHistory Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryMTHistory end")
	return result, nil
}

// QueryMTBalance 查询 MT 余额
func (m mtService) QueryMTBalance(classID, account string, params *models.QueryMTBalanceReq) (*models.QueryMTBalanceRes, models.Error) {
	log := m.Logger
	log.Debugln(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTBalance",
		"classID":  classID,
		"account":  account,
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryMTBalance start")

	nilRes := &models.QueryMTBalanceRes{}

	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if account == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "account"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "account"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryMTBalance Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMTBalance, classID, account), nil, bytesData)
	log.Debugf("QueryMTBalance body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryMTBalance DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryMTBalanceRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryMTBalance Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryMTBalance end")
	return result, nil
}
