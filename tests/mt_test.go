package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 创建 MT 类别接口示例
func TestCreateMTClass(t *testing.T) {
	params := &models.CreateMTClassReq{
		Name:        "类别1",
		Owner:       "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e",
		Data:        "创建类别",
		OperationId: OperationID,
	}

	result, err := client.MT.CreateMTClass(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 类别接口示例
func TestQueryMTClasses(t *testing.T) {
	params := &models.QueryMTClassesReq{
		Limit: "1",
	}
	result, err := client.MT.QueryMTClasses(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 类别详情接口示例
func TestQueryMTClass(t *testing.T) {
	result, err := client.MT.QueryMTClass("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 转让 MT 类别接口示例
func TestTransferMtClass(t *testing.T) {
	params := &models.TransferMTClassReq{
		Recipient:   "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd",
		OperationId: OperationID,
	}

	result, err := client.MT.TransferMTClass("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 发行 MT 接口示例
func TestIssueMT(t *testing.T) {
	params := &models.IssueMTReq{
		OperationId: OperationID,
	}

	result, err := client.MT.IssueMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 增发 MT 接口示例
func TestMintMT(t *testing.T) {
	params := &models.MintMTReq{
		OperationId: OperationID,
	}

	result, err := client.MT.MintMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 转让 MT 接口示例
func TestTransferMT(t *testing.T) {
	params := &models.TransferMTReq{
		Recipient:   "iaa1k3lq9vxtvf8erkqm49zrqwqz2lv4u9sq4wku5e",
		OperationId: OperationID,
	}
	result, err := client.MT.TransferMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 编辑 MT 接口示例
func TestEditMT(t *testing.T) {
	params := &models.EditMTReq{
		Data:        "秋海棠",
		OperationId: OperationID,
	}

	result, err := client.MT.EditMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 销毁 MT 接口示例
func TestBurnMT(t *testing.T) {
	params := &models.BurnMTReq{
		OperationId: OperationID,
	}

	result, err := client.MT.BurnMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 接口示例
func TestQueryMTs(t *testing.T) {
	params := &models.QueryMTsReq{
		Limit: "1",
	}
	result, err := client.MT.QueryMTs(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 详情接口示例
func TestQueryMT(t *testing.T) {
	result, err := client.MT.QueryMT("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 操作记录接口示例
func TestQueryMTHistory(t *testing.T) {
	result, err := client.MT.QueryMTHistory("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "12d129912d58426891a8549c6ba87e96deca33224acd7fedf64da70b36f90a69", nil)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 余额接口示例
func TestQueryMTBalance(t *testing.T) {
	result, err := client.MT.QueryMTBalance("b68fe234f258a95855db3f8b2d37e291a874df65a6ac7a66c4fc3780b1ab0bda", "iaa1qtag7eh9z7l94am0fcn3te5s8wx5j8cggkkrjd", nil)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
