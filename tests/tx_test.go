package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 上链交易结果查询示例
func TestQueryTxResult(t *testing.T) {
	client := GetClient()

	result := client.Tx.QueryTxResult("operationID1673512500")
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 上链交易排队状态查询示例
func TestQueryTxQueueInfo(t *testing.T) {
	client := GetClient()

	params := &models.QueryTxQueueInfoReq{OperationID: OperationID}
	result := client.Tx.QueryTxQueueInfo(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}
