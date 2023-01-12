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
	CreateNFTClass(params *models.CreateNFTClassReq) *models.CreateNFTResp
	QueryNFTClass(params *models.QueryNFTClassReq) *models.QueryNFTClassResp
	QueryNFTClassById(id string) *models.QueryNFTClassByIdResp
	TransfersNFClass(params *models.TransfersNFClassReq, classID string, owner string) *models.TransfersNFClassResp
	CreateNFT(params *models.CreateNFTReq, classID string) *models.CreateNFTResp
	TransfersNFT(params *models.TransfersNFTReq, classID string, owner string, nftID string) *models.TransfersNFTResp
	EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.EditNFTResp
	DeleteNFT(params *models.DeleteNFTReq, classID, owner, nftID string) *models.DeleteNFTResp
	BatchCreateNFT(params *models.BatchCreateNFTReq, classID string) *models.BatchCreateNFTResp
	BatchTransfersNFT(params *models.BatchTransfersNFTReq, owner string) *models.BatchTransfersNFTResp
	BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.BatchEditNFTResp
	BatchDeleteNFT(params *models.BatchDeleteNFTReq, owner string) *models.BatchDeleteNFTResp
	QueryNFT(params *models.QueryNFTReq) *models.QueryNFTResp
	QueryNFTById(classID string, nftID string) *models.QueryNFTByIdResp
	QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.QueryNFTHistoryResp
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
func (nft nftService) CreateNFTClass(params *models.CreateNFTClassReq) *models.CreateNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFTClass",
		"params":   params,
	})
	result := &models.CreateNFTResp{}
	log.Info("CreateNFTClass start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFTClass, bytesData, nil)
	log.WithFields(map[string]interface{}{
		"body":    string(body),
		"baseRes": baseRes,
	}).Debug()
	result.BaseRes = baseRes
	if baseRes.Code == -1 {
		nft.Logger.WithField("error", baseRes.Message).Errorln("DoHttpRequest")
		return result
	}
	// 请求成功
	if baseRes.Http.Code == http.StatusOK {
		if err := json.Unmarshal(body, &result); err != nil {
			nft.Logger.WithError(err).Errorln("Unmarshal body")
			result.Code = -1
			result.Message = err.Error()
			return result
		}
	}
	nft.Logger.Info("CreateNFTClass end")
	return result
}

// QueryNFTClass 查询 NFT 类别
func (nft nftService) QueryNFTClass(params *models.QueryNFTClassReq) *models.QueryNFTClassResp {
	log := nft.Logger.WithFields(map[string]interface{}{
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClass, nil, bytesData)
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
func (nft nftService) QueryNFTClassById(id string) *models.QueryNFTClassByIdResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "QueryNFTClassByIdReq",
		"id":       id,
	})
	result := &models.QueryNFTClassByIdResp{}
	log.Info("QueryNFTClassByIdReq start")
	//错误结果集合
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.CreateNFTClass+"/"+id, nil, nil)
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
func (nft nftService) TransfersNFClass(params *models.TransfersNFClassReq, classID string, owner string) *models.TransfersNFClassResp {
	log := nft.Logger.WithFields(map[string]interface{}{
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.TransfersNFClass+"/"+classID+"/"+owner, bytesData, nil)
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
func (nft nftService) CreateNFT(params *models.CreateNFTReq, classID string) *models.CreateNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "CreateNFT",
		"params":   params,
		"class_id": classID,
	})
	result := &models.CreateNFTResp{}
	log.Info("CreateNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFT+"/"+classID, bytesData, nil)
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

// TransfersNFT 转让 NFT
func (nft nftService) TransfersNFT(params *models.TransfersNFTReq, classID string, owner string, nftID string) *models.TransfersNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "TransfersNFT",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
		"nft_id":   nftID,
	})
	result := &models.TransfersNFTResp{}
	log.Info("TransfersNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.TransfersNFT+"/"+classID+"/"+owner+"/"+nftID, bytesData, nil)
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
	log.Info("TransfersNFT end")
	return result
}

// EditNFT 编辑 NFT
func (nft nftService) EditNFT(params *models.EditNFTReq, classID, owner, nftID string) *models.EditNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "EditNFT",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
		"nft_id":   nftID,
	})
	result := &models.EditNFTResp{}
	log.Info("EditNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPatch, models.EditNFT+"/"+classID+"/"+owner+"/"+nftID, bytesData, nil)
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
func (nft nftService) DeleteNFT(params *models.DeleteNFTReq, classID, owner, nftID string) *models.DeleteNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"params":   params,
		"class_id": classID,
		"owner":    owner,
		"nft_id":   nftID,
	})
	result := &models.DeleteNFTResp{}
	log.Info("DeleteNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodDelete, models.DeleteNFT+"/"+classID+"/"+owner+"/"+nftID, bytesData, nil)
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
func (nft nftService) BatchCreateNFT(params *models.BatchCreateNFTReq, classID string) *models.BatchCreateNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "BatchCreateNFT",
		"params":   params,
		"class_id": classID,
	})
	result := &models.BatchCreateNFTResp{}
	log.Info("BatchCreateNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.BatchCreateNFT+"/"+classID, bytesData, nil)
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

// BatchTransfersNFT 批量转让 NFT
func (nft nftService) BatchTransfersNFT(params *models.BatchTransfersNFTReq, owner string) *models.BatchTransfersNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.BatchTransfersNFTResp{}
	log.Info("BatchTransfersNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.BatchTransfersNFT+"/"+owner, bytesData, nil)
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
	log.Info("BatchTransfersNFT end")
	return result
}

// BatchEditNFT 批量编辑 NFT
func (nft nftService) BatchEditNFT(params *models.BatchEditNFTReq, owner string) *models.BatchEditNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.BatchEditNFTResp{}
	log.Info("BatchEditNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPatch, models.BatchEditNFT+"/"+owner, bytesData, nil)
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
func (nft nftService) BatchDeleteNFT(params *models.BatchDeleteNFTReq, owner string) *models.BatchDeleteNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.BatchDeleteNFTResp{}
	log.Info("BatchDeleteNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodDelete, models.BatchDeleteNFT+"/"+owner, bytesData, nil)
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
func (nft nftService) QueryNFT(params *models.QueryNFTReq) *models.QueryNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFT, nil, bytesData)
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
func (nft nftService) QueryNFTById(classID string, nftID string) *models.QueryNFTByIdResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
		"class_id": classID,
		"nft_id":   nftID,
	})
	result := &models.QueryNFTByIdResp{}
	log.Info("QueryNFTById start")
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTById+"/"+classID+"/"+nftID, nil, nil)
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
		if err := json.Unmarshal(body, &result.Data); err != nil {
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
func (nft nftService) QueryNFTHistory(params *models.QueryNFTHistoryReq, classID, nftID string) *models.QueryNFTHistoryResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"module":   "NFT",
		"function": "DeleteNFT",
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTById+"/"+classID+"/"+nftID+"/history", nil, bytesData)
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
