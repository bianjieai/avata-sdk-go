package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 创建 MT 类别接口示例
func TestCreateMTClass(t *testing.T) {
	tag := make(map[string]string)
	tag["123wwww"] = "werfdwerf"

	params := &models.CreateMTClassReq{
		Name:        "类别1",
		Owner:       "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e",
		Data:        "创建类别",
		OperationId: OperationID,
		Tag:         tag,
	}

	result := client.MT.CreateMTClass(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 查询 MT 类别接口示例
func TestQueryMTClasses(t *testing.T) {
	params := &models.QueryMTClassesReq{
		Limit: "1",
	}
	result := client.MT.QueryMTClasses(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var mtClassesRes models.QueryMTClassesRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &mtClassesRes)

	t.Logf("%+v \n", mtClassesRes)
}

// 查询 MT 类别详情接口示例
func TestQueryMTClass(t *testing.T) {
	result := client.MT.QueryMTClass("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda")
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var mtClassRes models.QueryMTClassRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &mtClassRes)

	t.Logf("%+v \n", mtClassRes)
}

// 转让 MT 类别接口示例
func TestTransferMtClass(t *testing.T) {
	tag := make(map[string]string)
	tag["20230112"] = "20230112"
	params := &models.TransferMTClassReq{
		Recipient:   "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd",
		OperationId: OperationID,
		Tag:         tag,
	}

	result := client.MT.TransferMTClass("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e", params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 发行 MT 接口示例
func TestIssueMT(t *testing.T) {
	params := &models.IssueMTReq{
		OperationId: OperationID,
	}

	result := client.MT.IssueMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 增发 MT 接口示例
func TestMintMT(t *testing.T) {
	params := &models.MintMTReq{
		OperationId: OperationID,
	}

	result := client.MT.MintMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 转让 MT 接口示例
func TestTransferMT(t *testing.T) {
	params := &models.TransferMTReq{
		Recipient:   "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e",
		OperationId: OperationID,
	}
	result := client.MT.TransferMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 编辑 MT 接口示例
func TestEditMT(t *testing.T) {
	params := &models.EditMTReq{
		Data:        "秋海棠",
		OperationId: OperationID,
	}

	result := client.MT.EditMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 销毁 MT 接口示例
func TestBurnMT(t *testing.T) {
	params := &models.BurnMTReq{
		OperationId: OperationID,
	}

	result := client.MT.BurnMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 查询 MT 接口示例
func TestQueryMTs(t *testing.T) {
	params := &models.QueryMTsReq{
		Limit: "1",
	}
	result := client.MT.QueryMTs(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var mtsRes models.QueryMTsRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &mtsRes)

	t.Logf("%+v \n", mtsRes)
}

// 查询 MT 详情接口示例
func TestQueryMT(t *testing.T) {
	result := client.MT.QueryMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69")
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var mtRes models.QueryMTRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &mtRes)

	t.Logf("%+v \n", mtRes)
}

// 查询 MT 操作记录接口示例
func TestQueryMTHistory(t *testing.T) {
	result := client.MT.QueryMTHistory("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", nil)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var mtHistory models.QueryMTHistoryRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &mtHistory)

	t.Logf("%+v \n", mtHistory)
}

// 查询 MT 余额接口示例
func TestQueryMTBalance(t *testing.T) {
	result := client.MT.QueryMTBalance("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", nil)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var mtBalanceRes models.QueryMTBalanceRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &mtBalanceRes)

	t.Logf("%+v \n", mtBalanceRes)
}
