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
	DoHttpRequest(method, path string, bodyParams, queryParams []byte) ([]byte, *models.Response)
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
func (h httpClient) DoHttpRequest(method, path string, bodyParams, queryParams []byte) ([]byte, *models.Response) {
	r, err := http.NewRequest(method, fmt.Sprintf("%s%s", h.baseParams.Domain, path), bytes.NewReader(bodyParams))
	if err != nil {
		return nil, &models.Response{
			Code:    models.CodeFailed,
			Http:    models.Http{},
			Message: err.Error(),
			Error:   models.Error{},
			Data:    nil,
		}
	}

	if method == http.MethodGet && queryParams != nil {
		queryParamsMap := make(map[string]string)
		if err = json.Unmarshal(queryParams, &queryParamsMap); err != nil {
			return nil, &models.Response{
				Code:    models.CodeFailed,
				Http:    models.Http{},
				Message: err.Error(),
				Error:   models.Error{},
				Data:    nil,
			}
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
		return nil, &models.Response{
			Code:    models.CodeFailed,
			Http:    models.Http{},
			Message: err.Error(),
			Error:   models.Error{},
			Data:    nil,
		}
	}
	defer res.Body.Close()

	result := &models.Response{
		Http: models.Http{
			Code:    res.StatusCode,
			Message: res.Status,
		},
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return nil, result
	}

	if err = json.Unmarshal(body, &result); err != nil {
		result.Code = models.CodeFailed
		result.Message = err.Error()
		return body, result
	}

	return body, result
}
