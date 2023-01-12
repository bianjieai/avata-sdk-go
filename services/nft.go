package services

import (
	"encoding/json"
	"net/http"

	"avata-sdk-go/models"
	"github.com/sirupsen/logrus"

	"avata-sdk-go/utils"
)

// NFTService NFT 接口
type NFTService interface {
	CreateNFTClass(params *models.CreateNFTClassReq) *models.TxRes                                           // 创建 NFT 类别
	QueryNFTClass(params *models.QueryNFTClassReq) *models.QueryNFTClassResp                                 // 查询 NFT 类别
	QueryNFTClassById(id string) *models.QueryNFTClassByIdResp                                               // 查询 NFT 类别详情
	TransfersNFClass(params *models.TransfersNFClassReq, classID, owner string) *models.TransfersNFClassResp // 转让 NFT 类别
	CreateNFT(params *models.CreateNFTReq, classID string) *models.TxRes                                     // 发行 NFT
	TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) *models.TxRes                   // 转让 NFT
	EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.TxRes                           // 编辑 NFT
	DeleteNFT(params *models.DeleteNFTReq, classID, owner, nftID string) *models.TxRes                       // 销毁 NFT
	BatchCreateNFT(params *models.BatchCreateNFTReq, classID string) *models.TxRes                           // 批量发行 NFT
	BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) *models.TxRes                         // 批量转让 NFT
	BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.TxRes                                 // 批量编辑 NFT
	BatchDeleteNFT(params *models.BatchDeleteNFTReq, owner string) *models.TxRes                             // 批量销毁 NFT
	QueryNFT(params *models.QueryNFTReq) *models.QueryNFTResp                                                // 查询 NFT
	QueryNFTById(classID, nftID string) *models.QueryNFTByIdResp                                             // 查询 NFT 详情
	QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.QueryNFTHistoryResp    // 查询 NFT 操作记录
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
func (n nftService) CreateNFTClass(params *models.CreateNFTClassReq) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFTClass",
		"params":   params,
	})
	result := &models.TxRes{}
	log.Info("CreateNFTClass start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFTClass, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()
	result.BaseRes = baseRes
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
	log.Info("CreateNFTClass end")
	return result
}

// QueryNFTClass 查询 NFT 类别
func (n nftService) QueryNFTClass(params *models.QueryNFTClassReq) *models.QueryNFTClassResp {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClassReq",
		"params":   params,
	})
	result := &models.QueryNFTClassResp{}
	log.Info("QueryNFTClass start")
	//参数集合
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClass, nil, bytesData)
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
	log.Info("QueryNFTClass end")
	return result
}

// QueryNFTClassById 查询 NFT 类别详情
func (n nftService) QueryNFTClassById(id string) *models.QueryNFTClassByIdResp {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClassByIdReq",
		"id":       id,
	})
	result := &models.QueryNFTClassByIdResp{}
	log.Info("QueryNFTClassByIdReq start")
	//错误结果集合
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClassById+"/"+id, nil, nil)
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
	log.Info("QueryNFTClassByIdReq end")
	return result
}

// TransfersNFClass 转让 NFT 类别
func (n nftService) TransfersNFClass(params *models.TransfersNFClassReq, classID, owner string) *models.TransfersNFClassResp {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "TransfersNFClass",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
	})
	result := &models.TransfersNFClassResp{}
	log.Info("TransfersNFClass start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.TransfersNFClass+"/"+classID+"/"+owner, bytesData, nil)
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
	log.Info("TransfersNFClass end")
	return result
}

// CreateNFT 发行 NFT
func (n nftService) CreateNFT(params *models.CreateNFTReq, classID string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFT",
		"params":   params,
		"class_id": classID,
	})
	result := &models.TxRes{}
	log.Info("CreateNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFT+"/"+classID, bytesData, nil)
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
	log.Info("CreateNFT end")
	return result
}

// TransferNFT 转让 NFT
func (n nftService) TransferNFT(params *models.TransferNFTReq, classID, owner, nftID string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "TransferNFT",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
		"nft_id":   nftID,
	})
	result := &models.TxRes{}
	log.Info("TransferNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.TransferNFT+"/"+classID+"/"+owner+"/"+nftID, bytesData, nil)
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
	log.Info("TransferNFT end")
	return result
}

// EditNFT 编辑 NFT
func (n nftService) EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "EditNFT",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
		"nft_id":   nftID,
	})
	result := &models.TxRes{}
	log.Info("EditNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPatch, models.EditNFT+"/"+classID+"/"+owner+"/"+nftID, bytesData, nil)
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
	log.Info("EditNFT end")
	return result
}

// DeleteNFT 销毁 NFT
func (n nftService) DeleteNFT(params *models.DeleteNFTReq, classID, owner, nftID string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
		"nft_id":   nftID,
	})
	result := &models.TxRes{}
	log.Info("DeleteNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodDelete, models.DeleteNFT+"/"+classID+"/"+owner+"/"+nftID, bytesData, nil)
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
	log.Info("DeleteNFT end")
	return result
}

// BatchCreateNFT 批量发行 NFT
func (n nftService) BatchCreateNFT(params *models.BatchCreateNFTReq, classID string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchCreateNFT",
		"params":   params,
		"class_id": classID,
	})
	result := &models.TxRes{}
	log.Info("BatchCreateNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateNFT+"/"+classID, bytesData, nil)
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
	log.Info("BatchCreateNFT end")
	return result
}

// BatchTransferNFT 批量转让 NFT
func (n nftService) BatchTransferNFT(params *models.BatchTransferNFTReq, owner string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchTransferNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.TxRes{}
	log.Info("BatchTransferNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPost, models.BatchTransferNFT+"/"+owner, bytesData, nil)
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
	log.Info("BatchTransferNFT end")
	return result
}

// BatchEditNFT 批量编辑 NFT
func (n nftService) BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchEditNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.TxRes{}
	log.Info("BatchEditNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodPatch, models.BatchEditNFT+"/"+owner, bytesData, nil)
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
	log.Info("BatchEditNFT end")
	return result
}

// BatchDeleteNFT 批量销毁 NFT
func (n nftService) BatchDeleteNFT(params *models.BatchDeleteNFTReq, owner string) *models.TxRes {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchDeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.TxRes{}
	log.Info("BatchDeleteNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodDelete, models.BatchDeleteNFT+"/"+owner, bytesData, nil)
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
	log.Info("BatchDeleteNFT end")
	return result
}

// QueryNFT 查询 NFT
func (n nftService) QueryNFT(params *models.QueryNFTReq) *models.QueryNFTResp {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFT",
		"params":   params,
	})
	result := &models.QueryNFTResp{}
	log.Info("QueryNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFT, nil, bytesData)
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
	log.Info("QueryNFT end")
	return result
}

// QueryNFTById 查询 NFT 详情
func (n nftService) QueryNFTById(classID, nftID string) *models.QueryNFTByIdResp {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTById",
		"class_id": classID,
		"nft_id":   nftID,
	})
	result := &models.QueryNFTByIdResp{}
	log.Info("QueryNFTById start")
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTById+"/"+classID+"/"+nftID, nil, nil)
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
	log.Info("QueryNFTById end")
	return result
}

// QueryNFTHistory 查询 NFT 操作记录
func (n nftService) QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.QueryNFTHistoryResp {
	log := n.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTHistory",
		"class_id": classID,
		"nft_id":   nftID,
	})
	result := &models.QueryNFTHistoryResp{}
	log.Info("QueryNFTHistory start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := n.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTHistory+"/"+classID+"/"+nftID+"/history", nil, bytesData)
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
	log.Info("TransfersNFClass end")
	return result
}
