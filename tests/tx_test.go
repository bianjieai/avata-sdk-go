package tests

import (
	"testing"
)

// 以 EVM 方式上链交易结果查询示例
func TestQueryTxResult(t *testing.T) {
	result, err := client.Tx.QueryTxResult("v3_TestUseContract34411")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询上链交易结果查询示例
func TestQueryNativeTxResult(t *testing.T) {
	result, err := client.Tx.QueryNativeTxResult("abcderf111shiheng")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 枚举值列表查询实例
func TestQueryTxTypes(t *testing.T) {
	result, err := client.Tx.QueryTxTypes()
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询枚举值列表实例
func TestQueryNativeTxTypes(t *testing.T) {
	result, err := client.Tx.QueryNativeTxTypes()
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
