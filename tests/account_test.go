package tests

import (
	"fmt"
	"testing"
	"time"

	client2 "avata-sdk-go"
	"avata-sdk-go/models"
	"avata-sdk-go/pkgs/configs"
	"github.com/sirupsen/logrus"
)

var OperationID = fmt.Sprintf("%s%d", "operationID", time.Now().Unix())

func GetClient() *client2.AvataClient {
	options := []configs.Options{
		configs.Level(logrus.DebugLevel),
		configs.HttpTimeout(1),
	}
	//client := client2.NewClient("域名", "项目参数 API KEY", "项目参数 API SECRET", log.ErrorLevel,options...)
	client := client2.NewClient("http://192.168.150.41:18081", "000001", "b2m2V1L1d1p8z0j3y5q4T5b4M4l0M45Y", options...)
	return client
}

// 创建链账户示例
func TestCreateAccount(t *testing.T) {
	client := GetClient()

	params := &models.CreateAccountReq{
		//Name:        "链账户1",
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
	client := GetClient()

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
func TestGetAccounts(t *testing.T) {
	client := GetClient()

	params := &models.GetAccountsReq{
		Account: "iaa1tf7wa9vm9zvlhxcdnctcxd3mag99uyefs58vjl",
	}

	result, err := client.Account.GetAccounts(params)
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询链账户操作记录示例
func TestGetAccountsHistory(t *testing.T) {
	client := GetClient()

	params := &models.GetAccountsHistoryReq{
		TxHash: "83333FF1BB96F17EC5F8ADD1FAEAC6AC9C6B7D2E463E35F1E3DB035FF9188C9E",
	}

	result, err := client.Account.GetAccountsHistory(params)
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("%+v \n", result)
}
