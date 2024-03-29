package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bianjieai/avata-sdk-go/v3/models"
)

type HttpClient interface {
	DoHttpRequest(method, path string, bodyParams, queryParams []byte) ([]byte, models.Error)
}

type httpClient struct {
	client     *http.Client
	baseParams models.BaseParams
}

func NewHttpClient(httpTimeout time.Duration, baseParams models.BaseParams) *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: httpTimeout,
		},
		baseParams: baseParams,
	}
}

// DoHttpRequest http 请求
func (h httpClient) DoHttpRequest(method, path string, bodyParams, queryParams []byte) ([]byte, models.Error) {
	r, err := http.NewRequest(method, fmt.Sprintf("%s%s", h.baseParams.Domain, path), bytes.NewReader(bodyParams))
	if err != nil {
		return nil, models.NewSDKError(err.Error())
	}

	if method == http.MethodGet && queryParams != nil {
		queryParamsMap := make(map[string]string)
		if err = json.Unmarshal(queryParams, &queryParamsMap); err != nil {
			return nil, models.NewSDKError(err.Error())
		}
		q := r.URL.Query() // 检查格式错误的键值对
		for k, v := range queryParamsMap {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode() // 编码成 url 形式
	}

	SignRequest(r, h.baseParams.APIKey, h.baseParams.APISecret)

	res, err := h.client.Do(r)
	if err != nil {
		return nil, models.NewSDKError(err.Error())
	}
	defer res.Body.Close()

	body := new(bytes.Buffer) // 写入器是一个 buffer 也会自动扩容，详见 buffer.Write
	_, err = io.Copy(body, res.Body)
	if err != nil {
		return nil, models.NewSDKError(err.Error())
	}
	bodyBytes := body.Bytes()

	if res.StatusCode != http.StatusOK {
		var response *models.Response
		if err = json.Unmarshal(bodyBytes, &response); err != nil {
			return bodyBytes, models.NewSDKError(err.Error())
		}
		return bodyBytes, models.NewAvataError(response.AvataError)
	}
	return bodyBytes, nil
}
