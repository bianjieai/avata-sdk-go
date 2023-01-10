package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"avata-sdk-go/models"
)

// DoHttpRequest http 请求
func DoHttpRequest(method, path string, httpTimeout int, baseParams models.BaseParams, bodyParams, queryParams []byte) (int, string, []byte, error) {
	client := &http.Client{Timeout: time.Duration(httpTimeout) * time.Second}
	r, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseParams.Domain, path), bytes.NewReader(bodyParams))
	if err != nil {
		return 0, "", nil, err
	}

	if method == http.MethodGet {
		queryParamsMap := make(map[string]string)
		if err = json.Unmarshal(queryParams, &queryParamsMap); err != nil {
			return 0, "", nil, err
		}
		q := r.URL.Query()
		for k, v := range queryParamsMap {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode()
	}

	SignRequest(r, baseParams.APIKey, baseParams.APISecret)

	res, err := client.Do(r)
	if err != nil {
		return 0, "", nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", nil, err
	}

	return res.StatusCode, res.Status, body, nil
}
