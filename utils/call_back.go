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

//v2版本签名回调验签
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
