package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 上链交易结果查询示例
func TestQueryTxResult(t *testing.T) {
	result := client.Tx.QueryTxResult("operationID1673512500")
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txResult models.QueryTxResultRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txResult)

	t.Logf("%+v \n", txResult)
}

// 上链交易排队状态查询示例
func TestQueryTxQueueInfo(t *testing.T) {
	params := &models.QueryTxQueueInfoReq{OperationID: OperationID}
	result := client.Tx.QueryTxQueueInfo(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txQueueInfo models.QueryTxQueueInfoRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txQueueInfo)

	t.Logf("%+v \n", txQueueInfo)
}
