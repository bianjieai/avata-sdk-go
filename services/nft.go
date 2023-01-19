package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/sirupsen/logrus"

	"github.com/bianjieai/avata-sdk-go/utils"
)

// NFTService NFT 接口
type NFTService interface {
	CreateNFTClass(params *models.CreateNFTClassReq) *models.Response                            // 创建 NFT 类别
	QueryNFTClass(params *models.QueryNFTClassReq) *models.Response                              // 查询 NFT 类别
	QueryNFTClassById(id string) *models.Response                                                // 查询 NFT 类别详情
	TransfersNFClass(params *models.TransfersNFClassReq, classID, owner string) *models.Response // 转让 NFT 类别
	CreateNFT(params *models.CreateNFTReq, classID string) *models.Response                      // 发行 NFT
	TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) *models.Response    // 转让 NFT
	EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.Response            // 编辑 NFT
	DeleteNFT(params *models.DeleteNFTReq, classID, owner, nftID string) *models.Response        // 销毁 NFT
	BatchCreateNFT(params *models.BatchCreateNFTReq, classID string) *models.Response            // 批量发行 NFT
	BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) *models.Response          // 批量转让 NFT
	BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.Response                  // 批量编辑 NFT
	BatchDeleteNFT(params *models.BatchDeleteNFTReq, owner string) *models.Response              // 批量销毁 NFT
	QueryNFT(params *models.QueryNFTReq) *models.Response                                        // 查询 NFT
	QueryNFTById(classID, nftID string) *models.Response                                         // 查询 NFT 详情
	QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.Response   // 查询 NFT 操作记录
}

type nftService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewNFTService(log *logrus.Logger, httpClient utils.HttpClient) *nftService {
	return &nftService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateNFTClass 创建 NFT 类别
func (n nftService) CreateNFTClass(params *models.CreateNFTClassReq) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFTClass",
		"params":   params,
	})
	log.Info("CreateNFTClass start")

	result := &models.Response{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationID == "" {
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
	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFTClass, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("CreateNFTClass end")
	return result
}

// QueryNFTClass 查询 NFT 类别
func (n nftService) QueryNFTClass(params *models.QueryNFTClassReq) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClassReq",
		"params":   params,
	})
	log.Info("QueryNFTClass start")

	result := &models.Response{}

	//参数集合
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClass, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryNFTClass end")
	return result
}

// QueryNFTClassById 查询 NFT 类别详情
func (n nftService) QueryNFTClassById(id string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClassByIdReq",
		"id":       id,
	})
	log.Info("QueryNFTClassByIdReq start")

	result := &models.Response{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "id")
		return result
	}

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFTClassById, id), nil, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryNFTClassByIdReq end")
	return result
}

// TransfersNFClass 转让 NFT 类别
func (n nftService) TransfersNFClass(params *models.TransfersNFClassReq, classID, owner string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "TransfersNFClass",
		"params":   params,
		"classID":  classID,
		"owner":    owner,
	})
	log.Info("TransfersNFClass start")

	result := &models.Response{}

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
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransfersNFClass, classID, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("TransfersNFClass end")
	return result
}

// CreateNFT 发行 NFT
func (n nftService) CreateNFT(params *models.CreateNFTReq, classID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFT",
		"params":   params,
		"classID":  classID,
	})
	log.Info("CreateNFT start")

	result := &models.Response{}

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
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.CreateNFT, classID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("CreateNFT end")
	return result
}

// TransferNFT 转让 NFT
func (n nftService) TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNFT",
		"params":   params,
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("TransferNFT start")

	result := &models.Response{}

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
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "nft_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferNFT, classID, owner, nftID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("TransferNFT end")
	return result
}

// EditNFT 编辑 NFT
func (n nftService) EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "EditNFT",
		"params":   params,
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("EditNFT start")

	result := &models.Response{}

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
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "nft_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.EditNFT, classID, owner, nftID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("EditNFT end")
	return result
}

// DeleteNFT 销毁 NFT
func (n nftService) DeleteNFT(params *models.DeleteNFTReq, classID, owner, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"params":   params,
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("DeleteNFT start")

	result := &models.Response{}

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
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "nft_id")
		return result
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "params")
		return result
	}
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.DeleteNFT, classID, owner, nftID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("DeleteNFT end")
	return result
}

// BatchCreateNFT 批量发行 NFT
func (n nftService) BatchCreateNFT(params *models.BatchCreateNFTReq, classID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchCreateNFT",
		"params":   params,
		"classID":  classID,
	})
	log.Info("BatchCreateNFT start")

	result := &models.Response{}

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
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.BatchCreateNFT, classID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchCreateNFT end")
	return result
}

// BatchTransferNFT 批量转让 NFT
func (n nftService) BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchTransferNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("BatchTransferNFT start")

	result := &models.Response{}

	// 校验必填参数
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
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.BatchTransferNFT, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchTransferNFT end")
	return result
}

// BatchEditNFT 批量编辑 NFT
func (n nftService) BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchEditNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("BatchEditNFT start")

	result := &models.Response{}

	// 校验必填参数
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
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.BatchEditNFT, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchEditNFT end")
	return result
}

// BatchDeleteNFT 批量销毁 NFT
func (n nftService) BatchDeleteNFT(params *models.BatchDeleteNFTReq, owner string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchDeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("BatchDeleteNFT start")

	result := &models.Response{}

	// 校验必填参数
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
	if params.OperationID == "" {
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

	body, result := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BatchDeleteNFT, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchDeleteNFT end")
	return result
}

// QueryNFT 查询 NFT
func (n nftService) QueryNFT(params *models.QueryNFTReq) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFT",
		"params":   params,
	})
	log.Info("QueryNFT start")

	result := &models.Response{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}
	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFT, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryNFT end")
	return result
}

// QueryNFTById 查询 NFT 详情
func (n nftService) QueryNFTById(classID, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTById",
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNFTById start")

	result := &models.Response{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "class_id")
		return result
	}
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "nft_id")
		return result
	}

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFTById, classID, nftID), nil, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryNFTById end")
	return result
}

// QueryNFTHistory 查询 NFT 操作记录
func (n nftService) QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTHistory",
		"params":   params,
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNFTHistory start")

	result := &models.Response{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFTHistory, classID, nftID), nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("TransfersNFClass end")
	return result
}
