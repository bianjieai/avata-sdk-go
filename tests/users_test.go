package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/v3/models"
)

// 创建钱包用户接口请求示例  p2s3Y1a0H3w130Z8
func TestCreateUsers(t *testing.T) {
	params := &models.CreateUsersReq{
		PhoneNum: "",
	}

	result, err := client.Users.CreateUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 认证钱包用户接口请求示例
func TestKycUsers(t *testing.T) {
	params := &models.KycUsersReq{
		UserType:       1,
		UserID:         "",
		Name:           "",
		CertificateNum: "",
	}
	result, err := client.Users.KycUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 更新钱包用户接口请求示例
func TestEditUsers(t *testing.T) {
	params := &models.EditUsersReq{
		UserID:   "",
		PhoneNum: "",
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
		PhoneNum: "",
	}

	result, err := client.Users.QueryUsers(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
