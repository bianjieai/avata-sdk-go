package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bianjieai/avata-sdk-go/models"
)

type HttpClient interface {
	DoHttpRequest(method, path string, bodyParams, queryParams []byte) ([]byte, models.BaseRes)
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
func (h httpClient) DoHttpRequest(method, path string, bodyParams, queryParams []byte) ([]byte, models.BaseRes) {
	baseRes := models.BaseRes{}

	r, err := http.NewRequest(method, fmt.Sprintf("%s%s", h.baseParams.Domain, path), bytes.NewReader(bodyParams))
	if err != nil {
		baseRes.Code = -1
		baseRes.Message = err.Error()
		return nil, baseRes
	}

	if method == http.MethodGet && queryParams != nil {
		queryParamsMap := make(map[string]string)
		if err = json.Unmarshal(queryParams, &queryParamsMap); err != nil {
			baseRes.Code = -1
			baseRes.Message = err.Error()
			return nil, baseRes
		}
		q := r.URL.Query()
		for k, v := range queryParamsMap {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode()
	}

	SignRequest(r, h.baseParams.APIKey, h.baseParams.APISecret)

	res, err := h.client.Do(r)
	if err != nil {
		baseRes.Code = -1
		baseRes.Message = err.Error()
		return nil, baseRes
	}
	defer res.Body.Close()

	baseRes.Http.Code = res.StatusCode
	baseRes.Http.Message = res.Status

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		baseRes.Code = -1
		baseRes.Message = err.Error()
		return nil, baseRes
	}

	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(body, &baseRes); err != nil {
			baseRes.Code = -1
			baseRes.Message = err.Error()
			return body, baseRes
		}
		return body, baseRes
	}

	return body, baseRes
}
