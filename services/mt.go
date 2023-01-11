package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
	"github.com/sirupsen/logrus"
)

type MtService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}

// CreateMTClass 创建 MT 类别
func (m MtService) CreateMTClass(params *models.CreateMTClassReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "CreateMTClass",
		"params":   params,
	})
	log.Info("CreateMTClass start")

	result := &models.TxRes{}

	// 校验必填参数
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "name")
		return result
	}
	if params.Owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodPost, models.CreateMTClass, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("CreateMTClass end")
	return result
}

// GetMTClasses 查询 MT 类别
func (m MtService) GetMTClasses(params *models.GetMTClassesReq) *models.GetMTClassesRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "GetMTClasses",
		"params":   params,
	})
	log.Info("GetMTClasses start")

	result := &models.GetMTClassesRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, models.GetMTClasses, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("GetMTClasses end")
	return result
}

// GetMTClass 查询 MT 类别详情
func (m MtService) GetMTClass(id string) *models.GetMTClassRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "GetMTClass",
		"id":       id,
	})
	log.Info("GetMTClass start")

	result := &models.GetMTClassRes{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "id")
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.GetMTClass, id), nil, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("GetMTClass end")
	return result
}

// TransferMTClass 转让 MT 类别
func (m MtService) TransferMTClass(classID, owner string, params *models.TransferMTClassReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "TransferMTClass",
		"classID":  classID,
		"owner":    owner,
		"params":   params,
	})
	log.Info("TransferMTClass start")

	result := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "recipient")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferMTClass, classID, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("TransferMTClass end")
	return result
}

// IssueMT 发行 MT
func (m MtService) IssueMT(classID string, params *models.IssueMTReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "IssueMT",
		"classID":  classID,
		"params":   params,
	})
	log.Info("IssueMT start")

	result := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.IssueMT, classID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("IssueMT end")
	return result
}

// MintMT 增发 MT
func (m MtService) MintMT(classID, mtID string, params *models.MintMTReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "MintMT",
		"classID":  classID,
		"mtID":     mtID,
		"params":   params,
	})
	log.Info("MintMT start")

	result := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.MintMT, classID, mtID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("MintMT end")
	return result
}

// TransferMT 转让 MT
func (m MtService) TransferMT(classID, owner, mtID string, params *models.TransferMTReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "TransferMT",
		"classID":  classID,
		"owner":    owner,
		"mtID":     mtID,
		"params":   params,
	})
	log.Info("TransferMT start")

	result := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "recipient")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferMT, classID, owner, mtID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("TransferMT end")
	return result
}

// EditMT 编辑 MT
func (m MtService) EditMT(classID, owner, mtID string, params *models.EditMTReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "EditMT",
		"classID":  classID,
		"owner":    owner,
		"mtID":     mtID,
		"params":   params,
	})
	log.Info("EditMT start")

	result := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}
	if params.Data == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "data"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "data")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.EditMT, classID, owner, mtID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("EditMT end")
	return result
}

// BurnMT 销毁 MT
func (m MtService) BurnMT(classID, owner, mtID string, params *models.BurnMTReq) *models.TxRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "BurnMT",
		"classID":  classID,
		"owner":    owner,
		"mtID":     mtID,
		"params":   params,
	})
	log.Info("BurnMT start")

	result := &models.TxRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BurnMT, classID, owner, mtID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("BurnMT end")
	return result
}

// GetMTs 查询 MT
func (m MtService) GetMTs(params *models.GetMTsReq) *models.GetMTsRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "GetMTs",
		"params":   params,
	})
	log.Info("GetMTs start")

	result := &models.GetMTsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, models.GetMTs, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("GetMTs end")
	return result
}

// GetMT 查询 MT 详情
func (m MtService) GetMT(classID, mtID string) *models.GetMTRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "GetMT",
		"classID":  classID,
		"mtID":     mtID,
	})
	log.Info("GetMT start")

	result := &models.GetMTRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.GetMT, classID, mtID), nil, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("GetMT end")
	return result
}

// GetMTHistory 查询 MT 操作记录
func (m MtService) GetMTHistory(classID, mtID string, params *models.GetAccountsHistoryReq) *models.GetMTHistoryRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "GetMTHistory",
		"classID":  classID,
		"mtID":     mtID,
		"params":   params,
	})
	log.Info("GetMTHistory start")

	result := &models.GetMTHistoryRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.GetMTHistory, classID, mtID), nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("GetMTHistory end")
	return result
}

// GetMTBalance 查询 MT 余额
func (m MtService) GetMTBalance(classID, account string, params *models.GetMTBalanceReq) *models.GetMTBalanceRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "GetMTBalance",
		"classID":  classID,
		"account":  account,
		"params":   params,
	})
	log.Info("GetMTBalance start")

	result := &models.GetMTBalanceRes{}

	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if account == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "account"))
		result.Code = -1
		result.Message = fmt.Sprintf(models.ErrParam, "account")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.GetMTBalance, classID, account), nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()

	result.BaseRes = baseRes

	// 记录错误日志
	if baseRes.Code == -1 {
		log.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}

	log.Info("GetMTBalance end")
	return result
}
