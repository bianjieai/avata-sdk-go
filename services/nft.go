package services

import (
	"encoding/json"
	"net/http"

	"avata-sdk-go/models"
	"avata-sdk-go/utils"

	"github.com/sirupsen/logrus"
)

type NftService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}

//创建nft类别
func (nft NftService) CreateNFTClass(params *models.CreateNFTClassReq) *models.CreateNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "CreateNFTClass",
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

//查询nft类别
func (nft NftService) QueryNFTClass(params *models.QueryNFTClassReq) *models.QueryNFTClassResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "QueryNFTClassReq",
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

//查询nft类别详情
func (nft NftService) QueryNFTClassById(id string) *models.QueryNFTClassByIdResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "QueryNFTClassByIdReq",
		"id":       id,
	})
	result := &models.QueryNFTClassByIdResp{}
	log.Info("QueryNFTClassByIdReq start")
	//错误结果集合
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTClassById+"/"+id, nil, nil)
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

//转让nft类别
func (nft NftService) TransfersNFClass(params *models.TransfersNFClassReq, class_id string, owner string) *models.TransfersNFClassResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "TransfersNFClass",
		"params":   params,
		"class_id": class_id,
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.TransfersNFClass+"/"+class_id+"/"+owner, bytesData, nil)
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

//发行nft
func (nft NftService) CreateNFT(params *models.CreateNFTReq, class_id string) *models.CreateNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "CreateNFT",
		"params":   params,
		"class_id": class_id,
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFT+"/"+class_id, bytesData, nil)
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

//转让nft
func (nft NftService) TransfersNFT(params *models.TransfersNFTReq, class_id string, owner string, nft_id string) *models.TransfersNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "TransfersNFT",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
		"nft_id":   nft_id,
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.TransfersNFT+"/"+class_id+"/"+owner+"/"+nft_id, bytesData, nil)
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

//编辑nft
func (nft NftService) EditorNFT(params *models.EditorNFTReq, class_id string, owner string, nft_id string) *models.EditorNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "EditorNFT",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
		"nft_id":   nft_id,
	})
	result := &models.EditorNFTResp{}
	log.Info("EditorNFT start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPatch, models.EditorNFT+"/"+class_id+"/"+owner+"/"+nft_id, bytesData, nil)
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
	log.Info("EditorNFT end")
	return result
}

//销毁nft
func (nft NftService) DeleteNFT(params *models.DeleteNFTReq, class_id string, owner string, nft_id string) *models.DeleteNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
		"nft_id":   nft_id,
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodDelete, models.DeleteNFT+"/"+class_id+"/"+owner+"/"+nft_id, bytesData, nil)
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

//批量发行nft
func (nft NftService) CreateNFTBatch(params *models.CreateNFTBatchReq, class_id string) *models.CreateNFTBatchResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "CreateNFTBatch",
		"params":   params,
		"class_id": class_id,
	})
	result := &models.CreateNFTBatchResp{}
	log.Info("CreateNFTBatch start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.CreateNFTBatch+"/"+class_id, bytesData, nil)
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
	log.Info("CreateNFTBatch end")
	return result
}

//批量转让nft
func (nft NftService) TransfersNFTBatch(params *models.TransfersNFTBatchReq, owner string) *models.TransfersNFTBatchResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.TransfersNFTBatchResp{}
	log.Info("TransfersNFTBatch start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPost, models.TransfersNFTBatch+"/"+owner, bytesData, nil)
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
	log.Info("TransfersNFTBatch end")
	return result
}

//批量编辑nft
func (nft NftService) EditorNFTBatch(params *models.EditorNFTBatchReq, owner string) *models.EditorNFTBatchResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.EditorNFTBatchResp{}
	log.Info("EditorNFTBatch start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodPatch, models.EditorNFTBatch+"/"+owner, bytesData, nil)
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
	log.Info("EditorNFTBatch end")
	return result
}

//批量销毁nft
func (nft NftService) DeleteNFTBatch(params *models.DeleteNFTBatchReq, owner string) *models.DeleteNFTBatchResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	result := &models.DeleteNFTBatchResp{}
	log.Info("DeleteNFTBatch start")
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		result.Code = -1
		result.Message = err.Error()
		return result
	}
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodDelete, models.DeleteNFTBatch+"/"+owner, bytesData, nil)
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
	log.Info("DeleteNFTBatch end")
	return result
}

//查询nft
func (nft NftService) QueryNFT(params *models.QueryNFTReq) *models.QueryNFTResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
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

//查询nft详情
func (nft NftService) QueryNFTById(class_id string, nft_id string) *models.QueryNFTByIdResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"class_id": class_id,
		"nft_id":   nft_id,
	})
	result := &models.QueryNFTByIdResp{}
	log.Info("QueryNFTById start")
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTById+"/"+class_id+"/"+nft_id, nil, nil)
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

//查询nft操作记录
func (nft NftService) QueryNFTHistory(params *models.QueryNFTHistoryReq, class_id string, nft_id string) *models.QueryNFTHistoryResp {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"class_id": class_id,
		"nft_id":   nft_id,
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
	body, baseRes := nft.HttpClient.DoHttpRequest(http.MethodGet, models.QueryNFTHistory+"/"+class_id+"/"+nft_id+"/history", nil, bytesData)
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
