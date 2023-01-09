package tests

import (
	"fmt"
	"testing"
	"time"

	client2 "avata-sdk-go"
	"avata-sdk-go/models"
	log "github.com/sirupsen/logrus"
)

var OperationID = fmt.Sprintf("%s%d", "operationID", time.Now().Unix())

func GetClient() *client2.AvataClient {
	client := client2.NewClient("域名", "项目参数 API KEY", "项目参数 API SECRET", log.ErrorLevel)
	return client
}

func TestCreateAccount(t *testing.T) {
	client := GetClient()

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
