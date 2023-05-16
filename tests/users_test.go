package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

//  user_id=72l390X471J0t0w3
func TestCreateUsers(t *testing.T) {
	params := &models.CreateUsersReq{
		UserType:        1,
		Name:            "nihao",
		Region:          2,
		CertificateType: 1,
		CertificateNum:  "341124199907064523",
		PhoneNum:        "18372232397",
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
		UserID:   "72l390X471J0t0w3",
		PhoneNum: "18372232396",
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
		UserType: "1",
		Code:     "341124199907064523",
	}

	result, err := client.Users.QueryUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
