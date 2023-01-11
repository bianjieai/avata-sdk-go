package services

import (
	"avata-sdk-go/models"
	"avata-sdk-go/utils"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type NftService struct {
	*logrus.Logger    // 日志
	models.BaseParams // 域名和项目参数
}

//创建nft类别
func (nft NftService) CreateNFTClass(params *models.CreateNFTClassReq) (*models.CreateNFTClassResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "CreateNFTClass",
		"params":   params,
	})
	log.Info("CreateNFTClass start")
	nilRes := &models.CreateNFTClassResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.CreateNFTClass, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.CreateNFTClassResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("CreateNFTClass end")
	return result, nil
}

//查询nft类别
func (nft NftService) QueryNFTClass(params *models.QueryNFTClassReq, path string) (*models.QueryNFTClassResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "QueryNFTClassReq",
		"params":   params,
	})
	log.Info("QueryNFTClass start")
	nilRes := &models.QueryNFTClassResp{}
	if path == "" {
		statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.QueryNFTClass, nft.BaseParams, nil)
		if err != nil {
			log.WithError(err).Errorln("DoHttpRequest")
			return nilRes, &models.Error{Exception: err}
		}
		log.Debugln("body: ", string(body))
		// 非 200 请求失败
		if statusCode != http.StatusOK {
			errorResponse := models.Response{}
			if err := json.Unmarshal(body, &errorResponse); err != nil {
				log.WithError(err).Errorln("Unmarshal body")
				return nilRes, &models.Error{Exception: err}
			}
			return nilRes, &models.Error{HttpResponse: models.HttpResponse{
				Status:     status,
				StatusCode: statusCode,
				Response:   errorResponse,
			}}
		}
		// 请求成功
		var result *models.QueryNFTClassResp
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		log.Info("QueryNFTClass end")
		return result, nil
	} else {
		statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, path, nft.BaseParams, nil)
		if err != nil {
			log.WithError(err).Errorln("DoHttpRequest")
			return nilRes, &models.Error{Exception: err}
		}
		log.Debugln("body: ", string(body))
		// 非 200 请求失败
		if statusCode != http.StatusOK {
			errorResponse := models.Response{}
			if err := json.Unmarshal(body, &errorResponse); err != nil {
				log.WithError(err).Errorln("Unmarshal body")
				return nilRes, &models.Error{Exception: err}
			}
			return nilRes, &models.Error{HttpResponse: models.HttpResponse{
				Status:     status,
				StatusCode: statusCode,
				Response:   errorResponse,
			}}
		}
		// 请求成功
		var result *models.QueryNFTClassResp
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		log.Info("QueryNFTClass end")
		return result, nil
	}

}

//查询nft类别详情
func (nft NftService) QueryNFTClassById(id string) (*models.QueryNFTClassByIdResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "QueryNFTClassByIdReq",
		"id":       id,
	})
	log.Info("QueryNFTClassByIdReq start")
	nilRes := &models.QueryNFTClassByIdResp{}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.CreateNFTClass+"/"+id, nft.BaseParams, nil)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.QueryNFTClassByIdResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("QueryNFTClassByIdReq end")
	return result, nil
}

//转让nft类别
func (nft NftService) TransfersNFClass(params *models.TransfersNFClassReq, class_id string, owner string) (*models.TransfersNFClassResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "TransfersNFClass",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
	})
	log.Info("TransfersNFClass start")
	nilRes := &models.TransfersNFClassResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.TransfersNFClass+"/"+class_id+"/"+owner, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.TransfersNFClassResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("TransfersNFClass end")
	return result, nil
}

//发行nft
func (nft NftService) CreateNFT(params *models.CreateNFTReq, class_id string) (*models.CreateNFTResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "CreateNFT",
		"params":   params,
		"class_id": class_id,
	})
	log.Info("CreateNFT start")
	nilRes := &models.CreateNFTResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.CreateNFT+"/"+class_id, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.CreateNFTResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("CreateNFT end")
	return result, nil
}

//转让nft
func (nft NftService) TransfersNFT(params *models.TransfersNFTReq, class_id string, owner string, nft_id string) (*models.TransfersNFTResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "TransfersNFT",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
		"nft_id":   nft_id,
	})
	log.Info("TransfersNFT start")
	nilRes := &models.TransfersNFTResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.TransfersNFT+"/"+class_id+"/"+owner+"/"+nft_id, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.TransfersNFTResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("TransfersNFT end")
	return result, nil
}

//编辑nft
func (nft NftService) EditorNFT(params *models.EditorNFTReq, class_id string, owner string, nft_id string) (*models.EditorNFTResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "EditorNFT",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
		"nft_id":   nft_id,
	})
	log.Info("EditorNFT start")
	nilRes := &models.EditorNFTResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPatch, models.EditorNFT+"/"+class_id+"/"+owner+"/"+nft_id, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.EditorNFTResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("EditorNFT end")
	return result, nil
}

//销毁nft
func (nft NftService) DeleteNFT(params *models.DeleteNFTReq, class_id string, owner string, nft_id string) (*models.DeleteNFTResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"class_id": class_id,
		"owner":    owner,
		"nft_id":   nft_id,
	})
	log.Info("DeleteNFT start")
	nilRes := &models.DeleteNFTResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodDelete, models.DeleteNFT+"/"+class_id+"/"+owner+"/"+nft_id, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.DeleteNFTResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("DeleteNFT end")
	return result, nil
}

//批量发行nft
func (nft NftService) CreateNFTBatch(params *models.CreateNFTBatchReq, class_id string) (*models.CreateNFTBatchResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "CreateNFTBatch",
		"params":   params,
		"class_id": class_id,
	})
	log.Info("CreateNFTBatch start")
	nilRes := &models.CreateNFTBatchResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.CreateNFTBatch+"/"+class_id, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.CreateNFTBatchResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("CreateNFTBatch end")
	return result, nil
}

//批量转让nft
func (nft NftService) TransfersNFTBatch(params *models.TransfersNFTBatchReq, owner string) (*models.TransfersNFTBatchResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("TransfersNFTBatch start")
	nilRes := &models.TransfersNFTBatchResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPost, models.TransfersNFTBatch+"/"+owner, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.TransfersNFTBatchResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("TransfersNFTBatch end")
	return result, nil
}

//批量编辑nft
func (nft NftService) EditorNFTBatch(params *models.EditorNFTBatchReq, owner string) (*models.EditorNFTBatchResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("EditorNFTBatch start")
	nilRes := &models.EditorNFTBatchResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodPatch, models.EditorNFTBatch+"/"+owner, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.EditorNFTBatchResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("EditorNFTBatch end")
	return result, nil
}

//批量销毁nft
func (nft NftService) DeleteNFTBatch(params *models.DeleteNFTBatchReq, owner string) (*models.DeleteNFTBatchResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
		"owner":    owner,
	})
	log.Info("DeleteNFTBatch start")
	nilRes := &models.DeleteNFTBatchResp{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Errorln("Marshal Params")
		return nilRes, &models.Error{Exception: err}
	}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodDelete, models.DeleteNFTBatch+"/"+owner, nft.BaseParams, bytesData)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.DeleteNFTBatchResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("DeleteNFTBatch end")
	return result, nil
}

//查询nft
func (nft NftService) QueryNFT(params *models.QueryNFTReq, path string) (*models.QueryNFTResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"params":   params,
	})
	log.Info("QueryNFT start")
	nilRes := &models.QueryNFTResp{}
	if path == "" {
		statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.QueryNFT, nft.BaseParams, nil)
		if err != nil {
			log.WithError(err).Errorln("DoHttpRequest")
			return nilRes, &models.Error{Exception: err}
		}
		log.Debugln("body: ", string(body))
		// 非 200 请求失败
		if statusCode != http.StatusOK {
			errorResponse := models.Response{}
			if err := json.Unmarshal(body, &errorResponse); err != nil {
				log.WithError(err).Errorln("Unmarshal body")
				return nilRes, &models.Error{Exception: err}
			}
			return nilRes, &models.Error{HttpResponse: models.HttpResponse{
				Status:     status,
				StatusCode: statusCode,
				Response:   errorResponse,
			}}
		}
		// 请求成功
		var result *models.QueryNFTResp
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		log.Info("QueryNFT end")
		return result, nil
	} else {
		statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, path, nft.BaseParams, nil)
		if err != nil {
			log.WithError(err).Errorln("DoHttpRequest")
			return nilRes, &models.Error{Exception: err}
		}
		log.Debugln("body: ", string(body))
		// 非 200 请求失败
		if statusCode != http.StatusOK {
			errorResponse := models.Response{}
			if err := json.Unmarshal(body, &errorResponse); err != nil {
				log.WithError(err).Errorln("Unmarshal body")
				return nilRes, &models.Error{Exception: err}
			}
			return nilRes, &models.Error{HttpResponse: models.HttpResponse{
				Status:     status,
				StatusCode: statusCode,
				Response:   errorResponse,
			}}
		}
		// 请求成功
		var result *models.QueryNFTResp
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		log.Info("QueryNFT end")
		return result, nil
	}
}

//查询nft详情
func (nft NftService) QueryNFTById(class_id string, nft_id string) (*models.QueryNFTByIdResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"class_id": class_id,
		"nft_id":   nft_id,
	})
	log.Info("QueryNFTById start")
	nilRes := &models.QueryNFTByIdResp{}
	statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.QueryNFTById+"/"+class_id+"/"+nft_id, nft.BaseParams, nil)
	if err != nil {
		log.WithError(err).Errorln("DoHttpRequest")
		return nilRes, &models.Error{Exception: err}
	}
	log.Debugln("body: ", string(body))
	// 非 200 请求失败
	if statusCode != http.StatusOK {
		errorResponse := models.Response{}
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		return nilRes, &models.Error{HttpResponse: models.HttpResponse{
			Status:     status,
			StatusCode: statusCode,
			Response:   errorResponse,
		}}
	}
	// 请求成功
	var result *models.QueryNFTByIdResp
	if err := json.Unmarshal(body, &result); err != nil {
		log.WithError(err).Errorln("Unmarshal body")
		return nilRes, &models.Error{Exception: err}
	}
	log.Info("QueryNFTById end")
	return result, nil
}

//查询nft操作记录
func (nft NftService) QueryNFTHistory(params *models.QueryNFTHistoryReq, path, class_id string, nft_id string) (*models.QueryNFTHistoryResp, *models.Error) {
	log := nft.Logger.WithFields(map[string]interface{}{
		"model":    "NFT",
		"functoin": "DeleteNFT",
		"class_id": class_id,
		"nft_id":   nft_id,
	})
	log.Info("QueryNFTHistory start")
	nilRes := &models.QueryNFTHistoryResp{}
	if path == "" {
		statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, models.QueryNFTById+"/"+class_id+"/"+nft_id+"/history", nft.BaseParams, nil)
		if err != nil {
			log.WithError(err).Errorln("DoHttpRequest")
			return nilRes, &models.Error{Exception: err}
		}
		log.Debugln("body: ", string(body))
		// 非 200 请求失败
		if statusCode != http.StatusOK {
			errorResponse := models.Response{}
			if err := json.Unmarshal(body, &errorResponse); err != nil {
				log.WithError(err).Errorln("Unmarshal body")
				return nilRes, &models.Error{Exception: err}
			}
			return nilRes, &models.Error{HttpResponse: models.HttpResponse{
				Status:     status,
				StatusCode: statusCode,
				Response:   errorResponse,
			}}
		}
		// 请求成功
		var result *models.QueryNFTHistoryResp
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		log.Info("QueryNFTHistory end")
		return result, nil
	} else {
		statusCode, status, body, err := utils.DoHttpRequest(http.MethodGet, path, nft.BaseParams, nil)
		if err != nil {
			log.WithError(err).Errorln("DoHttpRequest")
			return nilRes, &models.Error{Exception: err}
		}
		log.Debugln("body: ", string(body))
		// 非 200 请求失败
		if statusCode != http.StatusOK {
			errorResponse := models.Response{}
			if err := json.Unmarshal(body, &errorResponse); err != nil {
				log.WithError(err).Errorln("Unmarshal body")
				return nilRes, &models.Error{Exception: err}
			}
			return nilRes, &models.Error{HttpResponse: models.HttpResponse{
				Status:     status,
				StatusCode: statusCode,
				Response:   errorResponse,
			}}
		}
		// 请求成功
		var result *models.QueryNFTHistoryResp
		if err := json.Unmarshal(body, &result); err != nil {
			log.WithError(err).Errorln("Unmarshal body")
			return nilRes, &models.Error{Exception: err}
		}
		log.Info("QueryNFTHistory end")
		return result, nil
	}

}
