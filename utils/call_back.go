/*
 * @description:
 * @param:
 * @return:
 */
package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

//v1版本签名回调验签
func CallBackV1(r *http.Request, apiSecret string) string {
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
		return "FAILED"
	}
	return "SUCCESS"
}

//v2及以上版本请使用以下签名、验签
func CallBackV2(r *http.Request, path, apiSecret string) string {
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
		return "FAILED"
	}
	return "SUCCESS"
}

/**
 * @description: 接收推送消息
 * @param {*} version :版本
 * @param {*} apiSecret
 * @param {string} path ：路由地址(回调地址去掉域名)
 * @param {*http.Request} r
 * @return {*}
 */
func ReceiveMessages(version, apiSecret, path string, r *http.Request, businessFunction business) string {
	if version == "v1" {
		a := CallBackV1(r, apiSecret)
		return judge(a, businessFunction)
	}
	if version == "v2" || version == "v3" {
		a := CallBackV2(r, path, apiSecret)
		return judge(a, businessFunction)
	} else {
		return "版本不符合，要求请重新填写"
	}

}

func judge(result string, businessFunction business) string {
	if result == "SUCCESS" {
		//该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
		//TODO
		return ""
	} else {
		return "验签不通过"
	}
}

type business interface {
}
