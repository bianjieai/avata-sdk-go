package tests

import (
	"fmt"
	"testing"
	"time"

	client2 "avata-sdk-go"
	"avata-sdk-go/models"
)

var OperationID = fmt.Sprintf("%s%d", "operationID", time.Now().Unix())

func GetClient() *client2.AvataClient {
	client := client2.NewClient("https://stage.apis.avata.bianjie.ai", "dAuB90yF9Wi6PY59nr4R7mvY2ngiWBC", "AKIDz8krbsJ5yKBZQpn74WFkmLPx3gn")
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
