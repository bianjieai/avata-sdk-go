package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

func TestUseContract(t *testing.T) {
	data, err := utils.ABICoding("0x6b01389b141721da199f33c1bdacb1c8908f15d1612b77d605ed82730fb062dd", 1, "v2_contract")
	if err != nil {
		t.Log(err)
		return
	}
	params := &models.UseContractReq{
		To:          "0xc2B8C8849A02E8D98BE114c128aab536B4E98b62", //固定
		Data:        data,
		GasLimit:    500000,
		OperationID: "v2_TestUseContract2",
		From:        "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82",
	}
	result, err := client.Contract.UseContract(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestQueryContract(t *testing.T) {
	data, err := utils.QueryABICoding("0x6b01389b141721da199f33c1bdacb1c8908f15d1612b77d605ed82730fb062dd", 1)
	if err != nil {
		t.Log(err)
		return
	}
	params := &models.QueryContractReq{
		To:   "0xc2B8C8849A02E8D98BE114c128aab536B4E98b62", //固定
		Data: data,
	}
	result, err := client.Contract.QueryContract(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
