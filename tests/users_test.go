package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 创建钱包用户接口请求示例
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

// 更新钱包用户接口请求示例
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

// 查询钱包用户信息接口请求示例
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
