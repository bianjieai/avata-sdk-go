package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// 调用合约接口请求示例
// node="0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9"
func TestUseContract(t *testing.T) {
	data, err := utils.ABICoding("0x9ebafa531cb3627702260ea6660c6d6bec9b61e825477c1d4b304bac13723193", "v3_contract_test", "setABI", 1)
	if err != nil {
		t.Log(err)
		return
	}
	params := &models.UseContractReq{
		To: "0x88218318DF4C7b0077bA95C4C4F523968396e450",
		//To:   "0xc2B8C8849A02E8D98BE114c128aab536B4E98b62", //固定
		Data:        data,
		GasLimit:    600000,
		OperationID: "abc071100",
		//From:        "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
		From: "0x5461f3c0EE080460F1CC150d2a7D783869d0da0d",
	}
	result, err := client.Contract.UseContract(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询合约接口请求示例
func TestQueryContract(t *testing.T) {
	data, err := utils.QueryABICoding("0x9ebafa531cb3627702260ea6660c6d6bec9b61e825477c1d4b304bac13723193", "ABI", 1)
	if err != nil {
		t.Log(err)
		return
	}
	params := &models.QueryContractReq{
		To: "0x88218318DF4C7b0077bA95C4C4F523968396e450",
		//To:   "0xc2B8C8849A02E8D98BE114c128aab536B4E98b62", //固定
		Data: data,
	}
	result, err := client.Contract.QueryContract(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 解析合约示例
func TestResolve(t *testing.T) {
	result, err := utils.ABIResolver("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001076335f636f6e74726163745f7465737400000000000000000000000000000000", "ABI")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
