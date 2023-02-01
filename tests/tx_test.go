package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 上链交易结果查询示例
func TestQueryTxResult(t *testing.T) {
	result, err := client.Tx.QueryTxResult("operationID1673512500")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 上链交易排队状态查询示例
func TestQueryTxQueueInfo(t *testing.T) {
	params := &models.QueryTxQueueInfoReq{OperationID: OperationID}
	result, err := client.Tx.QueryTxQueueInfo(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
