package tests

import (
	"testing"
)

// 上链交易结果查询示例
func TestGetTxResult(t *testing.T) {
	client := GetClient()

	result := client.Tx.GetTxResult("operationID1673409104")
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 上链交易排队状态查询示例
func TestGetTxQueueInfo(t *testing.T) {
	client := GetClient()

	result := client.Tx.GetTxQueueInfo(nil)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}
