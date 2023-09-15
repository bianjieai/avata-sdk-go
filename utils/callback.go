package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bianjieai/avata-sdk-go/models"
)

// CallbackV1 V1 版本签名回调验签
func CallbackV1(r *http.Request, apiSecret string) bool {
	var bodyBytes []byte
	params := map[string]interface{}{}
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	// 把刚刚读出来的再写进去
	if bodyBytes != nil {
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	} //保持请求体不变
	paramsBody := map[string]interface{}{}
	_ = json.Unmarshal(bodyBytes, &paramsBody)

	for k, v := range paramsBody {
		k = "body_" + k
		params[k] = v
	}

	signRequest := r.Header.Get("X-Signature")

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(params)
	hexHash := hash(strings.TrimRight(bf.String(), "\n") + apiSecret)
	if hexHash != signRequest {
		return false
	}
	return true
}

// Callback V2 及以上版本请使用以下签名、验签
func Callback(r *http.Request, path, apiSecret string) bool {
	// 获取 path params
	params := map[string]interface{}{}
	params["path_url"] = path

	// 获取 body params
	// 把request的内容读取出来
	// go 1.16 之前的版本可以使用标准库 io/ioutil
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	// 把刚刚读出来的再写进去
	if bodyBytes != nil {
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	} //保持请求体不变
	paramsBody := map[string]interface{}{}
	_ = json.Unmarshal(bodyBytes, &paramsBody)

	for k, v := range paramsBody {
		k = "body_" + k
		params[k] = v
	}
	signRequest := r.Header.Get("X-Signature")
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(params)
	hexHash := hash(strings.TrimRight(bf.String(), "\n") + r.Header.Get("X-Timestamp") + apiSecret)
	if hexHash != signRequest {
		return false
	}
	return true
}

// OnCallback
/**
 * @description: 接收推送消息
 * @param {*} ctx 上下文
 * @param {*} version 版本，V1 需要传 APIVersionV1 , V2 或者 V3 版本传 APIVersionsOther
 * @param {*} apiSecret 项目密钥
 * @param {string} path 路由地址（回调地址去掉域名）
 * @param {*http.Request} r 该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
 * @param {http.ResponseWriter} w 返回给消息推送端
 * @param app 自己的业务逻辑代码（验证签名通过才会执行）
 * @return {*}
 */
func OnCallback(ctx context.Context, version, apiSecret, path string, r *http.Request, w http.ResponseWriter, app App) error {
	var onCallbackRes interface{}
	var kindRes *models.KindRes
	var kind string

	switch version {
	case models.APIVersionV1:
		// 验证签名
		result := CallbackV1(r, apiSecret)
		if !result {
			return models.NewSDKError("signature verification failed")
		}

		// 解析回调结果
		body := new(bytes.Buffer)
		if _, err := io.Copy(body, r.Body); err != nil {
			return models.NewSDKError(err.Error())
		}
		onCallbackResV1 := &models.OnCallbackResV1{}
		if err := json.Unmarshal(body.Bytes(), &onCallbackResV1); err != nil {
			return models.NewSDKError(fmt.Sprintf("unmarshal body failed: %s", err.Error()))
		}
		onCallbackRes = onCallbackResV1
	case models.APIVersionsOther:
		// 验证签名
		result := Callback(r, path, apiSecret)
		if !result {
			return models.NewSDKError("signature verification failed")
		}
		body := new(bytes.Buffer)
		if _, err := io.Copy(body, r.Body); err != nil {
			return models.NewSDKError(err.Error())
		}

		// 根据不同的服务模块（native/evm）解析回调结果
		if err := json.Unmarshal(body.Bytes(), &kindRes); err != nil {
			return models.NewSDKError(fmt.Sprintf("unmarshal body to kind failed: %s", err.Error()))
		}
		kind = kindRes.Kind
		switch kind {
		case models.Native:
			onCallbackResNative := &models.OnCallbackResNative{}
			if err := json.Unmarshal(body.Bytes(), &onCallbackResNative); err != nil {
				return models.NewSDKError(fmt.Sprintf("unmarshal body to onCallbackResNative  failed: %s", err.Error()))
			}
			onCallbackRes = onCallbackResNative
		case models.EVM:
			onCallbackResEVM := &models.OnCallbackResEVM{}
			if err := json.Unmarshal(body.Bytes(), &onCallbackResEVM); err != nil {
				return models.NewSDKError(fmt.Sprintf("unmarshal body onCallbackResEVM failed: %s", err.Error()))
			}
			onCallbackRes = onCallbackResEVM
		}
	default:
		return models.NewSDKError("version verification failed")
	}

	// 该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
	if err := app(ctx, version, kind, onCallbackRes); err != nil {
		return err
	}

	// 返回给消息推送端
	w.Write([]byte("SUCCESS"))

	return nil
}

// App
/**
 * @description: 回调函数的定义
 * @param {*} ctx 上下文
 * @param {*} version 版本，V1 版本为 APIVersionV1，V2 或者 V3 版本为 APIVersionsOther
 * @param {*} kind 区分服务，native/evm
 * @param onCallbackRes 回调服务推送参数
 * @return {*}
 */
type App func(ctx context.Context, version, kind string, onCallbackRes interface{}) error
