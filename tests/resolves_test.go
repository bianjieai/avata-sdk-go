package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

func TestResolves(t *testing.T) {
	owner := "0xF382c2b83aFdAe7c50dbea9e35ED462aa3cD5A08"
	name := "test123123.w"
	params := &models.SetResolvesReq{
		ResolveType: 2,
		Text: struct {
			Key       string "json:\"key,omitempty\""
			TextValue string "json:\"text_value\""
		}{
			"email", "120410123@qq.com",
		},
		OperationID: "Resolves",
	}

	result, err := client.Resolves.SetResolves(params, owner, name)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestQueryResolves(t *testing.T) {
	name := "test123123.w"
	params := &models.QueryResolvesReq{
		ResolveType: "2",
	}

	result, err := client.Resolves.QueryResolves(params, name)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestSetReverseResolves(t *testing.T) {
	owner := "0xF382c2b83aFdAe7c50dbea9e35ED462aa3cD5A08"
	params := &models.SetReverseResolvesReq{
		Name:        "test123123.w",
		OperationID: "SetReverseResolves",
	}

	result, err := client.Resolves.SetReverseResolves(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

func TestQueryReverseResolves(t *testing.T) {
	owner := "0xF382c2b83aFdAe7c50dbea9e35ED462aa3cD5A08"

	result, err := client.Resolves.QueryReverseResolves(owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
