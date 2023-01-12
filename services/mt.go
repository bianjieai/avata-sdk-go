package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"
)

// MTService MT 接口
type MTService interface {
	CreateMTClass(params *models.CreateMTClassReq) *models.TxRes                                           // 创建 MT 类别
	QueryMTClasses(params *models.QueryMTClassesReq) *models.QueryMTClassesRes                             // 查询 MT 类别
	QueryMTClass(id string) *models.QueryMTClassRes                                                        // 查询 MT 类别详情
	TransferMTClass(classID, owner string, params *models.TransferMTClassReq) *models.TxRes                // 转让 MT 类别
	IssueMT(classID string, params *models.IssueMTReq) *models.TxRes                                       // 发行 MT
	MintMT(classID, mtID string, params *models.MintMTReq) *models.TxRes                                   // 增发 MT
	TransferMT(classID, owner, mtID string, params *models.TransferMTReq) *models.TxRes                    // 转让 MT
	EditMT(classID, owner, mtID string, params *models.EditMTReq) *models.TxRes                            // 编辑 MT
	BurnMT(classID, owner, mtID string, params *models.BurnMTReq) *models.TxRes                            // 销毁 MT
	QueryMTs(params *models.QueryMTsReq) *models.QueryMTsRes                                               // 查询 MT
	QueryMT(classID, mtID string) *models.QueryMTRes                                                       // 查询 MT 详情
	QueryMTHistory(classID, mtID string, params *models.QueryAccountsHistoryReq) *models.QueryMTHistoryRes // 查询 MT 操作记录
	QueryMTBalance(classID, account string, params *models.QueryMTBalanceReq) *models.QueryMTBalanceRes    // 查询 MT 余额
}

type mtService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewMTService(log *logrus.Logger, httpClient utils.HttpClient) *mtService {
	return &mtService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateMTClass 创建 MT 类别
func (m mtService) CreateMTClass(params *models.CreateMTClassReq) *models.TxRes {
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
	if params.Owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("CreateMTClass end")
	return result
}

// QueryMTClasses 查询 MT 类别
func (m mtService) QueryMTClasses(params *models.QueryMTClassesReq) *models.QueryMTClassesRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTClasses",
		"params":   params,
	})
	log.Info("QueryMTClasses start")

	result := &models.QueryMTClassesRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, models.QueryMTClasses, nil, bytesData)
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

	log.Info("QueryMTClasses end")
	return result
}

// QueryMTClass 查询 MT 类别详情
func (m mtService) QueryMTClass(id string) *models.QueryMTClassRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTClass",
		"id":       id,
	})
	log.Info("QueryMTClass start")

	result := &models.QueryMTClassRes{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "id")
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMTClass, id), nil, nil)
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

	log.Info("QueryMTClass end")
	return result
}

// TransferMTClass 转让 MT 类别
func (m mtService) TransferMTClass(classID, owner string, params *models.TransferMTClassReq) *models.TxRes {
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
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "recipient")
		return result
	}
	if params.OperationId == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "operation_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("TransferMTClass end")
	return result
}

// IssueMT 发行 MT
func (m mtService) IssueMT(classID string, params *models.IssueMTReq) *models.TxRes {
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
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
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

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("IssueMT end")
	return result
}

// MintMT 增发 MT
func (m mtService) MintMT(classID, mtID string, params *models.MintMTReq) *models.TxRes {
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
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
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

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("MintMT end")
	return result
}

// TransferMT 转让 MT
func (m mtService) TransferMT(classID, owner, mtID string, params *models.TransferMTReq) *models.TxRes {
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
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
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
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "recipient")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("TransferMT end")
	return result
}

// EditMT 编辑 MT
func (m mtService) EditMT(classID, owner, mtID string, params *models.EditMTReq) *models.TxRes {
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
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
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
	if params.Data == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "data"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "data")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("EditMT end")
	return result
}

// BurnMT 销毁 MT
func (m mtService) BurnMT(classID, owner, mtID string, params *models.BurnMTReq) *models.TxRes {
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
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "owner")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}
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

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
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

	log.Info("BurnMT end")
	return result
}

// QueryMTs 查询 MT
func (m mtService) QueryMTs(params *models.QueryMTsReq) *models.QueryMTsRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTs",
		"params":   params,
	})
	log.Info("QueryMTs start")

	result := &models.QueryMTsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, models.QueryMTs, nil, bytesData)
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

	log.Info("QueryMTs end")
	return result
}

// QueryMT 查询 MT 详情
func (m mtService) QueryMT(classID, mtID string) *models.QueryMTRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMT",
		"classID":  classID,
		"mtID":     mtID,
	})
	log.Info("QueryMT start")

	result := &models.QueryMTRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMT, classID, mtID), nil, nil)
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

	log.Info("QueryMT end")
	return result
}

// QueryMTHistory 查询 MT 操作记录
func (m mtService) QueryMTHistory(classID, mtID string, params *models.QueryAccountsHistoryReq) *models.QueryMTHistoryRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTHistory",
		"classID":  classID,
		"mtID":     mtID,
		"params":   params,
	})
	log.Info("QueryMTHistory start")

	result := &models.QueryMTHistoryRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if mtID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "mt_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "mt_id")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMTHistory, classID, mtID), nil, bytesData)
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

	log.Info("QueryMTHistory end")
	return result
}

// QueryMTBalance 查询 MT 余额
func (m mtService) QueryMTBalance(classID, account string, params *models.QueryMTBalanceReq) *models.QueryMTBalanceRes {
	log := m.Logger.WithFields(map[string]interface{}{
		"module":   "MT",
		"function": "QueryMTBalance",
		"classID":  classID,
		"account":  account,
		"params":   params,
	})
	log.Info("QueryMTBalance start")

	result := &models.QueryMTBalanceRes{}

	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if account == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "account"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "account")
		return result
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, baseRes := m.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryMTBalance, classID, account), nil, bytesData)
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

	log.Info("QueryMTBalance end")
	return result
}
