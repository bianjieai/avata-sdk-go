package tests

import (
	"testing"

	"avata-sdk-go/models"
)

// 数字作品存证示例
func TestCreateRecord(t *testing.T) {
	client := GetClient()

	params := &models.CreateRecordReq{
		//IdentityType: 0,
		//IdentityName: "",
		//IdentityNum:  "",
		Type:        1,
		Name:        "血色海棠",
		Description: "问海棠花，谁留恋、未教飘坠。真个好，一般标格，聘梅双李。怯冷拟将苏幕护，怕惊莫把金铃缀。望铜梁、玉垒正春深，花空美。",
		Hash:        "234ertyujnbvfrukmnertyhgfetyjgfhgfryuhgf",
		HashType:    3,
		OperationId: OperationID,
	}

	result := client.Record.CreateRecord(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}
