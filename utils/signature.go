package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SignRequest 生成签名
func SignRequest(r *http.Request, apiKey, apiSecret string) *http.Request {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	// 获取 path params
	params := map[string]interface{}{}
	params["path_url"] = r.URL.Path

	// 获取 query params
	for k, v := range r.URL.Query() {
		k = "query_" + k
		params[k] = v[0]
	}

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

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(params)
	hexHash := hash(strings.TrimRight(bf.String(), "\n") + timestamp + apiSecret)

	r.Header.Set("X-Api-Key", apiKey)
	r.Header.Set("X-Signature", hexHash)
	r.Header.Set("X-Timestamp", timestamp)
	return r
}

func hash(oriText string) string {
	oriTextHashBytes := sha256.Sum256([]byte(oriText))
	return hex.EncodeToString(oriTextHashBytes[:])
}
