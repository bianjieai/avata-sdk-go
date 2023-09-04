/*
 * @description:
 * @param:
 * @return:
 */
package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 验证签名时传入的 API 版本号
const (
	APIVersionV1     = "v1" // v1 版本 AVATA Open API
	APIVersionsOther = ""   // 其它版本 AVATA Open API
)

// CallBackV1 v1 版本签名回调验签
func CallBackV1(r *http.Request, apiSecret string) bool {
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

// CallBack v2 及以上版本请使用以下签名、验签
func CallBack(r *http.Request, path, apiSecret string) bool {
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

/**
 * @description: 接收推送消息
 * @param {*} ctx 上下文
 * @param {*} version :版本 ：v1需要传 APIVersionV1 , v2或者v3版本传 APIVersionsOther
 * @param {*} apiSecret
 * @param {string} path ：路由地址(回调地址去掉域名)
 * @param {*http.Request} r ：该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
 * @param {http.ResponseWriter} w ：返回给消息推送端
 * @param app ：自己的业务逻辑代码（验证签名通过才会执行）
 * @return {*}
 */
func OnCallBack(ctx context.Context, version, apiSecret, path string, r *http.Request, w http.ResponseWriter, app App) error {
	var result bool

	// 验证签名
	switch version {
	case APIVersionV1:
		result = CallBackV1(r, apiSecret)
	case APIVersionsOther:
		result = CallBack(r, path, apiSecret)
	default:
		return models.NewSDKError("version verification failed")
	}
	if !result {
		return models.NewSDKError("signature verification failed")
	}

	// 该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
	app(ctx, r)

	// 返回给消息推送端
	w.Write([]byte("SUCCESS"))

	return nil
}

// App 回调函数的定义
type App func(ctx context.Context, r *http.Request)
