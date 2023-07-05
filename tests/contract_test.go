package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/utils"
)

// 调用合约测试示例
func TestUseContract(t *testing.T) {
	data, err := utils.ABICoding("0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9", 1, "v3_contract_test")
	if err != nil {
		t.Log(err)
		return
	}
	params := &models.UseContractReq{
		To:          "0xc2B8C8849A02E8D98BE114c128aab536B4E98b62", //固定
		Data:        data,
		GasLimit:    500000,
		OperationID: "v3_TestUseContract34411",
		From:        "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
	}
	result, err := client.Contract.UseContract(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询合约 测试示例
func TestQueryContract(t *testing.T) {
	data, err := utils.QueryABICoding("0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9", 1)
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
