package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 设置域名解析接口请求示例
func TestResolves(t *testing.T) {
	owner := "0xF382c2b83aFdAe7c50dbea9e35ED462aa3cD5A08"
	name := "test123123.w"
	params := &models.SetResolvesReq{
		ResolveType: 2,
		Text: struct {
			Key       string "json:\"key,omitempty\""
			TextValue string "json:\"text_value\""
		}{
			"email", "120410123@qq.com",
		},
		OperationID: "Resolves",
	}

	result, err := client.Resolves.SetResolves(params, owner, name)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询域名解析接口请求示例
func TestQueryResolves(t *testing.T) {
	name := "test123123.w"
	params := &models.QueryResolvesReq{
		ResolveType: "2",
	}

	result, err := client.Resolves.QueryResolves(params, name)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 设置域名反向解析接口请求示例
func TestSetReverseResolves(t *testing.T) {
	owner := "0xF382c2b83aFdAe7c50dbea9e35ED462aa3cD5A08"
	params := &models.SetReverseResolvesReq{
		Name:        "test123123.w",
		OperationID: "SetReverseResolves",
	}

	result, err := client.Resolves.SetReverseResolves(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询域名反向解析接口请求示例
func TestQueryReverseResolves(t *testing.T) {
	owner := "0xF382c2b83aFdAe7c50dbea9e35ED462aa3cD5A08"

	result, err := client.Resolves.QueryReverseResolves(owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
