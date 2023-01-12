package tests

import (
	"testing"

	"avata-sdk-go/models"
)

// 创建 MT 类别接口示例
func TestCreateMTClass(t *testing.T) {
	client := GetClient()

	tag := make(map[string]string)
	tag["123wwww"] = "werfdwerf"

	params := &models.CreateMTClassReq{
		Name:        "类别1",
		Owner:       "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e",
		OperationId: OperationID,
		Tag:         tag,
	}

	result := client.MT.CreateMTClass(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询 MT 类别接口示例
func TestGetMTClasses(t *testing.T) {
	client := GetClient()

	result := client.MT.GetMTClasses(nil)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询 MT 类别详情接口示例
func TestGetMTClass(t *testing.T) {
	client := GetClient()

	result := client.MT.GetMTClass("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda")
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 转让 MT 类别接口示例
func TestTransferMtClass(t *testing.T) {
	client := GetClient()

	params := &models.TransferMTClassReq{
		Recipient:   "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd",
		OperationId: OperationID,
	}

	result := client.MT.TransferMTClass("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e", params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 发行 MT 接口示例
func TestIssueMT(t *testing.T) {
	client := GetClient()

	params := &models.IssueMTReq{
		OperationId: OperationID,
	}

	result := client.MT.IssueMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 增发 MT 接口示例
func TestMintMT(t *testing.T) {
	client := GetClient()

	params := &models.MintMTReq{
		OperationId: OperationID,
	}

	result := client.MT.MintMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 转让 MT 接口示例
func TestTransferMT(t *testing.T) {
	client := GetClient()

	params := &models.TransferMTReq{
		Recipient:   "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e",
		OperationId: OperationID,
	}
	result := client.MT.TransferMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 编辑 MT 接口示例
func TestEditMT(t *testing.T) {
	client := GetClient()

	params := &models.EditMTReq{
		Data:        "秋海棠",
		OperationId: OperationID,
	}

	result := client.MT.EditMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 销毁 MT 接口示例
func TestBurnMT(t *testing.T) {
	client := GetClient()

	params := &models.BurnMTReq{
		OperationId: OperationID,
	}

	result := client.MT.BurnMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询 MT 接口示例
func TestGetMTs(t *testing.T) {
	client := GetClient()

	result := client.MT.GetMTs(nil)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询 MT 详情接口示例
func TestGetMT(t *testing.T) {
	client := GetClient()

	result := client.MT.GetMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69")
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询 MT 操作记录接口示例
func TestGetMTHistory(t *testing.T) {
	client := GetClient()

	result := client.MT.GetMTHistory("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", nil)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询 MT 余额接口示例
func TestGetMTBalance(t *testing.T) {
	client := GetClient()

	result := client.MT.GetMTBalance("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", nil)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}
