package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"avata-sdk-go/models"
)

// DoHttpRequest http 请求
func DoHttpRequest(method, path string, baseParams models.BaseParams, params []byte) (int, string, []byte, error) {
	r, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseParams.Domain, path), bytes.NewReader(params))
	if err != nil {
		return 0, "", nil, err
	}

	SignRequest(r, baseParams.APIKey, baseParams.APISecret)

	res, err := http.DefaultClient.Do(r)
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
