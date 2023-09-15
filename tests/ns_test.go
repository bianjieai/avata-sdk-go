package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 注册域名接口请求示例
func TestRegisterDomain(t *testing.T) {
	params := &models.RegisterDomainReq{
		Name:        "v3-ns-test-sh-xss-s.w",
		Owner:       "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
		OperationID: "v3_TestRegisterDomain11ss2121",
		Duration:    1,
	}
	result, err := client.NS.RegisterDomain(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询域名接口请求示例
func TestQueryDomain(t *testing.T) {
	params := &models.QueryDomainReq{
		Name: "v3-ns-test",
	}
	result, err := client.NS.QueryDomain(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询用户域名接口请求示例
func TestQueryDomains(t *testing.T) {
	owner := "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9"
	params := &models.QueryDomainsReq{}
	result, err := client.NS.QueryDomains(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 转让域名接口请求示例
func TestTransferDomin(t *testing.T) {
	owner := "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9"
	name := "v3-ns-test-sh-x-s.w"
	params := &models.TransferDomainReq{
		OperationID: "v3_TestTransferDomin",
		Recipient:   "0x849Ed8726B6D755721b0EeAd2Ea5BA371f9E3c22",
	}
	result, err := client.NS.TransferDomin(params, owner, name)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
