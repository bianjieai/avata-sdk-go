package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 创建 MT 类别接口示例
func TestCreateMTClass(t *testing.T) {
	params := &models.CreateMTClassReq{
		Name:        "testmt类别1",
		Owner:       "iaa1g4cuujvvf2q2du9gqstq8rjrf9u3uxu3sfay0d",
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
	result, err := client.MT.QueryMTClass("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 转让 MT 类别接口示例
func TestTransferMtClass(t *testing.T) {
	params := &models.TransferMTClassReq{
		Recipient:   "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz",
		OperationID: "test10",
	}

	result, err := client.MT.TransferMTClass("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "iaa1g4cuujvvf2q2du9gqstq8rjrf9u3uxu3sfay0d", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 发行 MT 接口示例
func TestIssueMT(t *testing.T) {
	params := &models.IssueMTReq{
		OperationID: "test_issueMt",
	}

	result, err := client.MT.IssueMT("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 增发 MT 接口示例
func TestMintMT(t *testing.T) {
	params := &models.MintMTReq{
		OperationID: "zengfamt11",
		Recipient:   "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz",
	}

	result, err := client.MT.MintMT("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "5bc9da432cb9771a70060f0e0258f61b8efd75a8a5328e7a12261474d0ec19bc", params)
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
		OperationID: "testtransfermt",
	}
	result, err := client.MT.TransferMT("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz", "5bc9da432cb9771a70060f0e0258f61b8efd75a8a5328e7a12261474d0ec19bc", params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 编辑 MT 接口示例
func TestEditMT(t *testing.T) {
	params := &models.EditMTReq{
		Data:        "测试mt",
		OperationID: "thisisatesteditmtforalex11",
	}

	result, err := client.MT.EditMT("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz", "5bc9da432cb9771a70060f0e0258f61b8efd75a8a5328e7a12261474d0ec19bc", params)
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

	result, err := client.MT.BurnMT("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz", "5bc9da432cb9771a70060f0e0258f61b8efd75a8a5328e7a12261474d0ec19bc", params)
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
	result, err := client.MT.QueryMT("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "5bc9da432cb9771a70060f0e0258f61b8efd75a8a5328e7a12261474d0ec19bc")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 操作记录接口示例
func TestQueryMTHistory(t *testing.T) {
	result, err := client.MT.QueryMTHistory("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "5bc9da432cb9771a70060f0e0258f61b8efd75a8a5328e7a12261474d0ec19bc", nil)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 MT 余额接口示例
func TestQueryMTBalance(t *testing.T) {
	result, err := client.MT.QueryMTBalance("20f89a93e6a20f2001bb2e70ea89b0e87be051e8b8ab74fa7d29effe93d8c348", "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz", nil)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
