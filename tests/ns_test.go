package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

//{\"data\":{\"module\":2,\"operation\":1,\"tx_hash\":\"0xfae2835167aa9ec87f71a929f7bdec0875055b1b748531b97165e446ab71f2ef\",\"status\":1,\"message\":\"\",\"block_height\":17183281,\"timestamp\":\"2023-04-20 02:53:48 +0000 UTC\",\"nft\":null,\"record\":null,\"ns\":{\"name\":\"v2-ns-test.w\",\"owner\":\"0x1A6f8Ed0d40Fcb915e59444AA096241146108D82\",\"node\":\"0x6b01389b141721da199f33c1bdacb1c8908f15d1612b77d605ed82730fb062dd\",\"expires\":1713581624}}}
func TestRegisterDomain(t *testing.T) {
	params := &models.RegisterDomainReq{
		Name:        "v2-ns-test.w",
		Owner:       "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82",
		OperationID: "v2_TestRegisterDomain1",
		Duration:    1,
	}
	result, err := client.NS.RegisterDomain(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestQueryDomain(t *testing.T) {
	params := &models.QueryDomainReq{
		Name: "v2-ns-test",
		Tld:  "W",
	}
	result, err := client.NS.QueryDomain(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestQueryDomains(t *testing.T) {
	owner := "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82"
	params := &models.QueryDomainsReq{}
	result, err := client.NS.QueryDomains(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestTransferDomin(t *testing.T) {
	owner := "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82"
	name := "v2-ns-test.w"
	params := &models.TransferDominReq{
		Recipient:   "0x9c3d37463fCA8Cd2cec3548a63f1910ec2Cb0BCe",
		OperationID: "v2_TestTransferDomin",
	}
	result, err := client.NS.TransferDomin(params, owner, name)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
