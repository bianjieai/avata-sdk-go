package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siddontang/go-log/loggers"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// NFTService NFT 接口
type NFTService interface {
	CreateNFTClass(params *models.CreateNFTClassReq) (*models.TxRes, models.Error)                                       // 创建 NFT 类别
	QueryNFTClasses(params *models.QueryNFTClassesReq) (*models.QueryNFTClassesRes, models.Error)                        // 查询 NFT 类别
	QueryNFTClass(id string) (*models.QueryNFTClassRes, models.Error)                                                    // 查询 NFT 类别详情
	TransferNFTClass(params *models.TransferNFClassReq, classID, owner string) (*models.TxRes, models.Error)             // 转让 NFT 类别
	MintNFT(params *models.MintNFTReq, classID string) (*models.TxRes, models.Error)                                     // 发行 NFT
	TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)               // 转让 NFT
	EditNFT(params *models.EditNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                       // 编辑 NFT
	BurnNFT(params *models.BurnNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                       // 销毁 NFT
	BatchMintNFT(params *models.BatchMintNFTReq, classID string) (*models.TxRes, models.Error)                           // 批量发行 NFT
	BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) (*models.TxRes, models.Error)                     // 批量转让 NFT
	BatchEditNFT(params *models.BatchEditNFTReq, owner string) (*models.TxRes, models.Error)                             // 批量编辑 NFT
	BatchBurnNFT(params *models.BatchBurnNFTReq, owner string) (*models.TxRes, models.Error)                             // 批量销毁 NFT
	QueryNFTs(params *models.QueryNFTsReq) (*models.QueryNFTsRes, models.Error)                                          // 查询 NFT
	QueryNFT(classID, nftID string) (*models.QueryNFTRes, models.Error)                                                  // 查询 NFT 详情
	QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) (*models.QueryNFTHistoryRes, models.Error) // 查询 NFT 操作记录
}

type nftService struct {
	Logger loggers.Advanced // 日志
	utils.HttpClient
}

func NewNFTService(log loggers.Advanced, httpClient utils.HttpClient) *nftService {
	return &nftService{
		Logger:     log,
		HttpClient: httpClient,
	}
}

// CreateNFTClass 创建 NFT 类别
func (n nftService) CreateNFTClass(params *models.CreateNFTClassReq) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFTClass",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateNFTClass start")

	nilRes := &models.TxRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("CreateNFTClass Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFTClass, bytesData, nil)
	log.Debugf("CreateNFTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateNFTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateNFTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateNFTClass end")
	return result, nil
}

// QueryNFTClasses 查询 NFT 类别
func (n nftService) QueryNFTClasses(params *models.QueryNFTClassesReq) (*models.QueryNFTClassesRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClasses",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryNFTClasses start")

	nilRes := &models.QueryNFTClassesRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNFTClasses Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClasses, nil, bytesData)
	log.Debugf("QueryNFTClasses body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNFTClasses DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNFTClassesRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNFTClasses Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNFTClasses end")
	return result, nil
}

// QueryNFTClass 查询 NFT 类别详情
func (n nftService) QueryNFTClass(id string) (*models.QueryNFTClassRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClass",
		"id":       id,
	})
	log.Info("QueryNFTClass start")

	nilRes := &models.QueryNFTClassRes{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "id"))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFTClass, id), nil, nil)
	log.Debugf("QueryNFTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNFTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNFTClassRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNFTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNFTClass end")
	return result, nil
}

// TransferNFTClass 转让 NFT 类别
func (n nftService) TransferNFTClass(params *models.TransferNFClassReq, classID, owner string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNFTClass",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
	})
	log.Info("TransferNFTClass start")

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
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferNFTClass Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferNFTClass, classID, owner), bytesData, nil)
	log.Debugf("TransferNFTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferNFTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferNFTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferNFTClass end")
	return result, nil
}

// MintNFT 发行 NFT
func (n nftService) MintNFT(params *models.MintNFTReq, classID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "MintNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
	})
	log.Info("MintNFT start")

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
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("MintNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.MintNFT, classID), bytesData, nil)
	log.Debugf("MintNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("MintNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("MintNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("MintNFT end")
	return result, nil
}

// TransferNFT 转让 NFT
func (n nftService) TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("TransferNFT start")

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
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "nft_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferNFT, classID, owner, nftID), bytesData, nil)
	log.Debugf("TransferNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferNFT end")
	return result, nil
}

// EditNFT 编辑 NFT
func (n nftService) EditNFT(params *models.EditNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "EditNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("EditNFT start")

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
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "nft_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("EditNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.EditNFT, classID, owner, nftID), bytesData, nil)
	log.Debugf("EditNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("EditNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("EditNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("EditNFT end")
	return result, nil
}

// BurnNFT 销毁 NFT
func (n nftService) BurnNFT(params *models.BurnNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "BurnNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("BurnNFT start")

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
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "nft_id"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BurnNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BurnNFT, classID, owner, nftID), bytesData, nil)
	log.Debugf("BurnNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BurnNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BurnNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BurnNFT end")
	return result, nil
}

// BatchMintNFT 批量发行 NFT
func (n nftService) BatchMintNFT(params *models.BatchMintNFTReq, classID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchMintNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
	})
	log.Info("BatchMintNFT start")

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
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BatchMintNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.BatchMintNFT, classID), bytesData, nil)
	log.Debugf("BatchMintNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BatchMintNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BatchMintNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BatchMintNFT end")
	return result, nil
}

// BatchTransferNFT 批量转让 NFT
func (n nftService) BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchTransferNFT",
		"params":   fmt.Sprintf("%v", params),
		"owner":    owner,
	})
	log.Info("BatchTransferNFT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BatchTransferNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.BatchTransferNFT, owner), bytesData, nil)
	log.Debugf("BatchTransferNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BatchTransferNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BatchTransferNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BatchTransferNFT end")
	return result, nil
}

// BatchEditNFT 批量编辑 NFT
func (n nftService) BatchEditNFT(params *models.BatchEditNFTReq, owner string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchEditNFT",
		"params":   fmt.Sprintf("%v", params),
		"owner":    owner,
	})
	log.Info("BatchEditNFT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BatchEditNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.BatchEditNFT, owner), bytesData, nil)
	log.Debugf("BatchEditNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BatchEditNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BatchEditNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BatchEditNFT end")
	return result, nil
}

// BatchBurnNFT 批量销毁 NFT
func (n nftService) BatchBurnNFT(params *models.BatchBurnNFTReq, owner string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchBurnNFT",
		"params":   fmt.Sprintf("%v", params),
		"owner":    owner,
	})
	log.Info("BatchBurnNFT start")

	nilRes := &models.TxRes{}

	// 校验必填参数
	if owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("BatchBurnNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BatchBurnNFT, owner), bytesData, nil)
	log.Debugf("BatchBurnNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BatchBurnNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BatchBurnNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BatchBurnNFT end")
	return result, nil
}

// QueryNFTs 查询 NFT
func (n nftService) QueryNFTs(params *models.QueryNFTsReq) (*models.QueryNFTsRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTs",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryNFTs start")

	nilRes := &models.QueryNFTsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNFTs Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTs, nil, bytesData)
	log.Debugf("QueryNFTs body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNFTs DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNFTsRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNFTs Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNFTs end")
	return result, nil
}

// QueryNFT 查询 NFT 详情
func (n nftService) QueryNFT(classID, nftID string) (*models.QueryNFTRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFT",
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNFT start")

	nilRes := &models.QueryNFTRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "nft_id"))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFT, classID, nftID), nil, nil)
	log.Debugf("QueryNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNFTRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNFT end")
	return result, nil
}

// QueryNFTHistory 查询 NFT 操作记录
func (n nftService) QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) (*models.QueryNFTHistoryRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTHistory",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNFTHistory start")

	nilRes := &models.QueryNFTHistoryRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNFTHistory Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNFTHistory, classID, nftID), nil, bytesData)
	log.Debugf("QueryNFTHistory body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNFTHistory DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNFTHistoryRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNFTHistory Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNFTHistory end")
	return result, nil
}
