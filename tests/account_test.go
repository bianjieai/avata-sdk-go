package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/bianjieai/avata-sdk-go/v3/models"
)

var OperationID = fmt.Sprintf("%s%d", "operationID", time.Now().Unix())

// 创建链账户示例
func TestCreateAccount(t *testing.T) {
	params := &models.CreateAccountReq{
		Name:        "链账户12138",
		OperationID: operationID,
	}

	result, err := client.Account.CreateAccount(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 批量创建链账户示例
func TestBatchCreateAccounts(t *testing.T) {
	params := &models.BatchCreateAccountsReq{
		Count:       2,
		OperationID: operationID,
	}

	result, err := client.Account.BatchCreateAccounts(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询链账户示例
func TestQueryAccounts(t *testing.T) {
	params := &models.QueryAccountsReq{
		PageKey: "",
	}
	result, err := client.Account.QueryAccounts(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 EVM 模块链账户操作记录示例
func TestQueryAccountsHistory(t *testing.T) {
	params := &models.QueryAccountsHistoryReq{
		Account: "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
	}

	result, err := client.Account.QueryAccountsHistory(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询原生模块链账户操作记录示例
func TestNativeQueryAccountsHistory(t *testing.T) {
	params := &models.QueryNativeAccountsHistoryReq{
		Account: "iaa1jjmwg5ah27aynuwt2phwa8sfvzh4lvvlelddxm",
	}

	result, err := client.Account.QueryNativeAccountsHistory(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
