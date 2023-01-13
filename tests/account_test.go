package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	client2 "github.com/bianjieai/avata-sdk-go"
	"github.com/bianjieai/avata-sdk-go/configs"
	"github.com/bianjieai/avata-sdk-go/models"
)

var OperationID = fmt.Sprintf("%s%d", "operationID", time.Now().Unix())

func GetClient() *client2.AvataClient {
	options := []configs.Options{
		configs.Level(logrus.DebugLevel),
		configs.HttpTimeout(15 * time.Second),
	}
	//client := client2.NewClient("域名", "项目参数 API KEY", "项目参数 API SECRET", options...)
	client := client2.NewClient("http://192.168.150.41:18081", "000001", "b2m2V1L1d1p8z0j3y5q4T5b4M4l0M45Y", options...)
	return client
}

// 创建链账户示例
func TestCreateAccount(t *testing.T) {
	client := GetClient()

	params := &models.CreateAccountReq{
		Name:        "链账户1",
		OperationID: OperationID,
	}

	result := client.Account.CreateAccount(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 批量创建链账户示例
func TestBatchCreateAccounts(t *testing.T) {
	client := GetClient()

	params := &models.BatchCreateAccountsReq{
		Count:       3,
		OperationID: OperationID,
	}

	result := client.Account.BatchCreateAccounts(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询链账户示例
func TestQueryAccounts(t *testing.T) {
	client := GetClient()

	params := &models.QueryAccountsReq{
		//Account: "iaa1tf7wa9vm9zvlhxcdnctcxd3mag99uyefs58vjl",
	}

	result := client.Account.QueryAccounts(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询链账户操作记录示例
func TestQueryAccountsHistory(t *testing.T) {
	client := GetClient()

	params := &models.QueryAccountsHistoryReq{
		TxHash: "83333FF1BB96F17EC5F8ADD1FAEAC6AC9C6B7D2E463E35F1E3DB035FF9188C9E",
	}

	result := client.Account.QueryAccountsHistory(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}
