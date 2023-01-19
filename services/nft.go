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
	CreateNFTClass(params *models.CreateNFTClassReq) *models.Response                           // 创建 NFT 类别
	QueryNFTClasses(params *models.QueryNFTClassesReq) *models.Response                         // 查询 NFT 类别
	QueryNFTClass(id string) *models.Response                                                   // 查询 NFT 类别详情
	TransferNFTClass(params *models.TransferNFClassReq, classID, owner string) *models.Response // 转让 NFT 类别
	MintNFT(params *models.MintNFTReq, classID string) *models.Response                         // 发行 NFT
	TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) *models.Response   // 转让 NFT
	EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.Response           // 编辑 NFT
	BurnNFT(params *models.BurnNFTReq, classID, owner, nftID string) *models.Response           // 销毁 NFT
	BatchMintNFT(params *models.BatchMintNFTReq, classID string) *models.Response               // 批量发行 NFT
	BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) *models.Response         // 批量转让 NFT
	BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.Response                 // 批量编辑 NFT
	BatchBurnNFT(params *models.BatchBurnNFTReq, owner string) *models.Response                 // 批量销毁 NFT
	QueryNFTs(params *models.QueryNFTsReq) *models.Response                                     // 查询 NFT
	QueryNFT(classID, nftID string) *models.Response                                            // 查询 NFT 详情
	QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.Response  // 查询 NFT 操作记录
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

// QueryNFTClasses 查询 NFT 类别
func (n nftService) QueryNFTClasses(params *models.QueryNFTClassesReq) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClasses",
		"params":   params,
	})
	log.Info("QueryNFTClasses start")

	result := &models.Response{}

	//参数集合
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClasses, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryNFTClasses end")
	return result
}

// QueryNFTClass 查询 NFT 类别详情
func (n nftService) QueryNFTClass(id string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClass",
		"id":       id,
	})
	log.Info("QueryNFTClass start")

	result := &models.Response{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		result.Code = models.CodeFailed
		result.Message = fmt.Sprintf(models.ErrParam, "id")
		return result
	}

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFTClass, id), nil, nil)
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

// TransferNFTClass 转让 NFT 类别
func (n nftService) TransferNFTClass(params *models.TransferNFClassReq, classID, owner string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNFTClass",
		"params":   params,
		"classID":  classID,
		"owner":    owner,
	})
	log.Info("TransferNFTClass start")

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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferNFTClass, classID, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("TransferNFTClass end")
	return result
}

// MintNFT 发行 NFT
func (n nftService) MintNFT(params *models.MintNFTReq, classID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "MintNFT",
		"params":   params,
		"classID":  classID,
	})
	log.Info("MintNFT start")

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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.MintNFT, classID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("MintNFT end")
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

// BurnNFT 销毁 NFT
func (n nftService) BurnNFT(params *models.BurnNFTReq, classID, owner, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BurnNFT",
		"params":   params,
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("BurnNFT start")

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

	body, result := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BurnNFT, classID, owner, nftID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BurnNFT end")
	return result
}

// BatchMintNFT 批量发行 NFT
func (n nftService) BatchMintNFT(params *models.BatchMintNFTReq, classID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchMintNFT",
		"params":   params,
		"classID":  classID,
	})
	log.Info("BatchMintNFT start")

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

	body, result := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.BatchMintNFT, classID), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchMintNFT end")
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

// BatchBurnNFT 批量销毁 NFT
func (n nftService) BatchBurnNFT(params *models.BatchBurnNFTReq, owner string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchBurnNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("BatchBurnNFT start")

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

	body, result := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BatchBurnNFT, owner), bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("BatchBurnNFT end")
	return result
}

// QueryNFTs 查询 NFT
func (n nftService) QueryNFTs(params *models.QueryNFTsReq) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTs",
		"params":   params,
	})
	log.Info("QueryNFTs start")

	result := &models.Response{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return result
	}
	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTs, nil, bytesData)
	log.WithFields(map[string]interface{}{
		"body":   string(body),
		"result": result,
	}).Debug()

	// 记录错误日志
	if result.Code == models.CodeFailed {
		log.WithField("error", result.Message).Errorln("DoHttpRequest")
		return result
	}

	log.Info("QueryNFTs end")
	return result
}

// QueryNFT 查询 NFT 详情
func (n nftService) QueryNFT(classID, nftID string) *models.Response {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFT",
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNFT start")

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

	body, result := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFT, classID, nftID), nil, nil)
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

	log.Info("QueryNFTHistory end")
	return result
}
