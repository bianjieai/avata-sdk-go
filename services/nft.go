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
	CreateNFTClass(params *models.CreateNFTClassReq) (*models.TxRes, models.Error)                                                         // 创建 NFT 类别
	CreateNativeNFTClass(params *models.CreateNativeNFTClassReq) (*models.TxRes, models.Error)                                             // 以原生方式创建 NFT 类别
	QueryNFTClasses(params *models.QueryNFTClassesReq) (*models.QueryNFTClassesRes, models.Error)                                          // 查询 NFT 类别
	QueryNativeNFTClasses(params *models.QueryNativeNFTClassesReq) (*models.QueryNativeNFTClassesRes, models.Error)                        // 以原生方式 查询 NFT 类别
	QueryNFTClass(id string) (*models.QueryNFTClassRes, models.Error)                                                                      // 查询 NFT 类别详情
	QueryNativeNFTClass(id string) (*models.QueryNativeNFTClassRes, models.Error)                                                          // 以原生方式 查询 NFT 类别详情
	TransferNFTClass(params *models.TransferNFClassReq, classID, owner string) (*models.TxRes, models.Error)                               // 转让 NFT 类别
	TransferNativeNFTClass(params *models.TransferNativeNFClassReq, classID, owner string) (*models.TxRes, models.Error)                   // 以原生方式 转让 NFT 类别
	MintNFT(params *models.MintNFTReq, classID string) (*models.TxRes, models.Error)                                                       // 发行 NFT
	MintNativeNFT(params *models.MintNativeNFTReq, classID string) (*models.TxRes, models.Error)                                           // 以原生方式 发行 NFT
	TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                                 // 转让 NFT
	TransferNativeNFT(params *models.TransferNativeNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                     // 以原生方式 转让 NFT
	EditNFT(params *models.EditNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                                         // 编辑 NFT
	EditNativeNFT(params *models.EditNativeNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                             // 以原生方式 编辑 NFT
	BurnNFT(params *models.BurnNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                                         // 销毁 NFT
	BurnNativeNFT(params *models.BurnNativeNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error)                             // 以原生方式 销毁 NFT
	QueryNFTs(params *models.QueryNFTsReq) (*models.QueryNFTsRes, models.Error)                                                            // 查询 NFT
	QueryNativeNFTs(params *models.QueryNativeNFTsReq) (*models.QueryNativeNFTsRes, models.Error)                                          // 以原生方式 查询 NFT
	QueryNFT(classID, nftID string) (*models.QueryNFTRes, models.Error)                                                                    // 查询 NFT 详情
	QueryNativeNFT(classID, nftID string) (*models.QueryNativeNFTRes, models.Error)                                                        // 以原生方式 查询 NFT 详情
	QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) (*models.QueryNFTHistoryRes, models.Error)                   // 查询 NFT 操作记录
	QueryNativeNFTHistory(params *models.QueryNativeNFTHistoryReq, classID, nftID string) (*models.QueryNativeNFTHistoryRes, models.Error) // 以原生方式 查询 NFT 操作记录
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
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	if params.Symbol == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "symbol"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "symbol"))
	}
	if params.Owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
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

// CreateNativeNFTClass 以原生方式创建 NFT 类别
func (n nftService) CreateNativeNFTClass(params *models.CreateNativeNFTClassReq) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNativeNFTClass",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("CreateNativeNFTClass start")

	nilRes := &models.TxRes{}

	if params == nil {
		log.Debugln(fmt.Sprintf(models.ErrParam, "params"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "params"))
	}
	if params.Name == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "name"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "name"))
	}
	if params.Owner == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "owner"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "owner"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("CreateNativeNFTClass Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNativeNFTClass, bytesData, nil)
	log.Debugf("CreateNativeNFTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("CreateNativeNFTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("CreateNativeNFTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("CreateNativeNFTClass end")
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

// QueryNativeNFTClasses 以原生方式查询 NFT 类别
func (n nftService) QueryNativeNFTClasses(params *models.QueryNativeNFTClassesReq) (*models.QueryNativeNFTClassesRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNativeNFTClasses",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryNativeNFTClasses start")

	nilRes := &models.QueryNativeNFTClassesRes{}
	if params.CountTotal != "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "CountTotal"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "CountTotal"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNativeNFTClasses Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNativeNFTClasses, nil, bytesData)
	log.Debugf("QueryNativeNFTClasses body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeNFTClasses DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeNFTClassesRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeNFTClasses Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeNFTClasses end")
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

// QueryNativeNFTClass 以原生方式查询 NFT 类别详情
func (n nftService) QueryNativeNFTClass(id string) (*models.QueryNativeNFTClassRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNativeNFTClass",
		"id":       id,
	})
	log.Info("QueryNativeNFTClass start")

	nilRes := &models.QueryNativeNFTClassRes{}

	// 校验必填参数
	if id == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "id"))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNativeNFTClass, id), nil, nil)
	log.Debugf("QueryNativeNFTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeNFTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeNFTClassRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeNFTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeNFTClass end")
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
	// todo check 改字段后非空的校验 （对对应的字段作非空校验）

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

// TransferNativeNFTClass  以原生方式转让 NFT 类别
func (n nftService) TransferNativeNFTClass(params *models.TransferNativeNFClassReq, classID, owner string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNativeNFTClass",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
	})
	log.Info("TransferNativeNFTClass start")

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
	//todo: 将operationid 和recipient 删除后对新增的字段做非空校验

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferNFTClass Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferNativeNFTClass, classID, owner), bytesData, nil)
	log.Debugf("TransferNativeNFTClass body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferNativeNFTClass DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferNativeNFTClass Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferNativeNFTClass end")
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
	if params.Uri == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "uri"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "uri"))
	}
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "recipient"))
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

// MintNativeNFT 以原生方式发行 NFT
func (n nftService) MintNativeNFT(params *models.MintNativeNFTReq, classID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "MintNativeNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
	})
	log.Info("MintNativeNFT start")

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
		log.Errorf("MintNativeNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.MintNativeNFT, classID), bytesData, nil)
	log.Debugf("MintNativeNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("MintNativeNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("MintNativeNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("MintNativeNFT end")
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
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "recipient"))
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

// TransferNativeNFT  以原生方式转让 NFT
func (n nftService) TransferNativeNFT(params *models.TransferNativeNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNativeNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("TransferNativeNFT start")

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
	if params.Recipient == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "recipient"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "recipient"))
	}
	if params.OperationID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "operation_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "operation_id"))
	}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("TransferNativeNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPost, fmt.Sprintf(models.TransferNativeNFT, classID, owner, nftID), bytesData, nil)
	log.Debugf("TransferNativeNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("TransferNativeNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("TransferNativeNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("TransferNativeNFT end")
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
	if params.Uri == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "uri"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "uri"))
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

// EditNativeNFT 以原生方式 编辑 NFT
func (n nftService) EditNativeNFT(params *models.EditNativeNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "EditNativeNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("EditNativeNFT start")

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

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodPatch, fmt.Sprintf(models.EditNativeNFT, classID, owner, nftID), bytesData, nil)
	log.Debugf("EditNativeNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("EditNativeNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("EditNativeNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("EditNativeNFT end")
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

// BurnNativeNFT 以原生方式 销毁 NFT
func (n nftService) BurnNativeNFT(params *models.BurnNativeNFTReq, classID, owner, nftID string) (*models.TxRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "BurnNativeNFT",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"owner":    owner,
		"nftID":    nftID,
	})
	log.Info("BurnNativeNFT start")

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
		log.Errorf("BurnNativeNFT Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodDelete, fmt.Sprintf(models.BurnNativeNFT, classID, owner, nftID), bytesData, nil)
	log.Debugf("BurnNativeNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("BurnNativeNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.TxRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("BurnNativeNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("BurnNativeNFT end")
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

// QueryNativeNFTs 以原生方式查询 NFT
func (n nftService) QueryNativeNFTs(params *models.QueryNativeNFTsReq) (*models.QueryNativeNFTsRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNativeNFTs",
		"params":   fmt.Sprintf("%v", params),
	})
	log.Info("QueryNativeNFTs start")

	nilRes := &models.QueryNativeNFTsRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNativeNFTs Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNativeNFTs, nil, bytesData)
	log.Debugf("QueryNativeNFTs body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeNFTs DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeNFTsRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeNFTs Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeNFTs end")
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

// QueryNativeNFT 以原生方式查询 NFT 详情
func (n nftService) QueryNativeNFT(classID, nftID string) (*models.QueryNativeNFTRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNativeNFT",
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNativeNFT start")

	nilRes := &models.QueryNativeNFTRes{}

	// 校验必填参数
	if classID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "class_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "class_id"))
	}
	if nftID == "" {
		log.Debugln(fmt.Sprintf(models.ErrParam, "nft_id"))
		return nilRes, models.InvalidParam(fmt.Sprintf(models.ErrParam, "nft_id"))
	}
	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNativeNFT, classID, nftID), nil, nil)
	log.Debugf("QueryNFT body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNFT DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeNFTRes{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNFT Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeNFT end")
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

// QueryNativeNFTHistory 以原生方式查询 NFT 操作记录
func (n nftService) QueryNativeNFTHistory(params *models.QueryNativeNFTHistoryReq, classID, nftID string) (*models.QueryNativeNFTHistoryRes, models.Error) {
	log := n.Logger
	log.Debugln(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNativeNFTHistory",
		"params":   fmt.Sprintf("%v", params),
		"classID":  classID,
		"nftID":    nftID,
	})
	log.Info("QueryNativeNFTHistory start")

	nilRes := &models.QueryNativeNFTHistoryRes{}

	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Errorf("QueryNativeNFTHistory Marshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Marshal Params: %s", err.Error()))
	}

	body, errorRes := n.HttpClient.DoHttpRequest(http.MethodGet, fmt.Sprintf(models.QueryNativeNFTHistory, classID, nftID), nil, bytesData)
	log.Debugf("QueryNativeNFTHistory body: %s", string(body))
	if errorRes != nil {
		log.Errorf("QueryNativeNFTHistory DoHttpRequest error: %s", errorRes.Error())
		return nilRes, errorRes
	}

	result := &models.QueryNativeNFTHistoryRes{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Errorf("QueryNativeNFTHistory Unmarshal Params: %s", err.Error())
		return nilRes, models.NewSDKError(fmt.Sprintf("Unmarshal Params: %s", err.Error()))
	}

	log.Info("QueryNativeNFTHistory end")
	return result, nil
}
