package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/v3/models"
)

// 数字作品存证接口请求示例
func TestCreateRecord(t *testing.T) {
	params := &models.CreateRecordReq{
		IdentityType: 1,
		IdentityNum:  "",
		IdentityName: "测试数据123",
		Type:         1,
		Name:         "trstttt123",
		Description:  "recordtest",
		Hash:         "ac7f0f712ab9d13fdbaed27a82bf1e62",
		HashType:     1,
		OperationId:  "abcabc111",
	}

	result, err := client.Record.CreateRecord(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
