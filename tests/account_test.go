package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/bianjieai/avata-sdk-go/models"
)

var OperationID = fmt.Sprintf("%s%d", "operationID", time.Now().Unix())

// 创建链账户示例
func TestCreateAccount(t *testing.T) {
	params := &models.CreateAccountReq{
		Name:        "链账户1",
		OperationID: OperationID,
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
		Count:       3,
		OperationID: OperationID,
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
		Account: "iaa1tf7wa9vm9zvlhxcdnctcxd3mag99uyefs58vjl",
	}

	result, err := client.Account.QueryAccounts(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询链账户操作记录示例
func TestQueryAccountsHistory(t *testing.T) {
	params := &models.QueryAccountsHistoryReq{
		//TxHash: "83333FF1BB96F17EC5F8ADD1FAEAC6AC9C6B7D2E463E35F1E3DB035FF9188C9E",
	}

	result, err := client.Account.QueryAccountsHistory(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
