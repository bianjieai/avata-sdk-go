package tests

import (
	"testing"
)

// EVM 模块上链交易结果查询接口请求示例
func TestQueryTxResult(t *testing.T) {
	result, err := client.Tx.QueryTxResult("abc071100")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 原生模块上链交易结果查询接口请求示例
func TestQueryNativeTxResult(t *testing.T) {
	result, err := client.Tx.QueryNativeTxResult("abc071100")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// EVM 模块查询枚举值列表接口请求实例
func TestQueryTxTypes(t *testing.T) {
	result, err := client.Tx.QueryTxTypes()
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 原生模块查询枚举值列表接口请求实例
func TestQueryNativeTxTypes(t *testing.T) {
	result, err := client.Tx.QueryNativeTxTypes()
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
