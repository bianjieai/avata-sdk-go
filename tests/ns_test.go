package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// {\"data\":{\"module\":2,\"operation\":1,\"tx_hash\":\"0xfae2835167aa9ec87f71a929f7bdec0875055b1b748531b97165e446ab71f2ef\",\"status\":1,\"message\":\"\",\"block_height\":17183281,\"timestamp\":\"2023-04-20 02:53:48 +0000 UTC\",\"nft\":null,\"record\":null,\"ns\":{\"name\":\"v2-ns-test.w\",\"owner\":\"0x1A6f8Ed0d40Fcb915e59444AA096241146108D82\",\"node\":\"0x6b01389b141721da199f33c1bdacb1c8908f15d1612b77d605ed82730fb062dd\",\"expires\":1713581624}}}
// 注册域名示例
func TestRegisterDomain(t *testing.T) {
	params := &models.RegisterDomainReq{
		Name:        "v3-ns-test-sh-xs-s.w",
		Owner:       "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
		OperationID: "v3_TestRegisterDomain11ss221",
		Duration:    1,
	}
	result, err := client.NS.RegisterDomain(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询域名示例
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

// 查询用户域名示例
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

// 转让域名实例
func TestTransferDomin(t *testing.T) {
	owner := "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9"
	name := "v3-ns-test-sh-x-s.w"
	params := &models.TransferDominReq{
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
