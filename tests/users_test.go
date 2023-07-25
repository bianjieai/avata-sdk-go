package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

func TestCreateUsers(t *testing.T) {
	params := &models.CreateUsersReq{
		UserType:        1,
		Name:            "创建用户",
		PhoneNum:        "13875512645",
		CertificateType: 1,
		CertificateNum:  "341127199307262359",
	}

	result, err := client.Users.CreateUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestEditUsers(t *testing.T) {
	params := &models.EditUsersReq{
		UserID:   "31",
		PhoneNum: "1231321",
	}

	result, err := client.Users.EditUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
func TestQueryUsers(t *testing.T) {
	params := &models.QueryUsersReq{
		UserType: "",
		Code:     "",
	}

	result, err := client.Users.QueryUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
