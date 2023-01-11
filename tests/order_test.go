package tests

import (
	"fmt"
	"testing"
	"time"

	"avata-sdk-go/models"
)

var orderID = fmt.Sprintf("orderID_%v", time.Now().Unix())

// 购买能量值/业务费接口示例
func TestCreateOrder(t *testing.T) {
	client := GetClient()

	params := &models.CreateOrderReq{
		Account:   "0x7982C2FEEECCB2A86C5346762AF9DCAC4DF79219",
		Amount:    100,
		OrderType: "gas",
		OrderId:   orderID,
	}

	result := client.Order.CreateOrder(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询能量值/业务费购买结果列表接口
func TestGetOrders(t *testing.T) {
	client := GetClient()

	result := client.Order.GetOrders(nil)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 查询能量值/业务费购买结果接口示例
func TestGetOrder(t *testing.T) {
	client := GetClient()

	result := client.Order.GetOrder("orderID_1673342033")
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}

// 批量购买能量值接口示例
func TestCreateBatchOrder(t *testing.T) {
	client := GetClient()

	var list []models.List
	list = append(list, models.List{
		Account: "iaa1cz8c3ka0wskwdxdm204jvxrmzxmd3yuy7tm7k9",
		Amount:  100,
	})
	params := &models.CreateBatchOrderReq{
		List:    list,
		OrderId: orderID,
	}

	result := client.Order.CreateBatchOrder(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}

	t.Logf("%+v \n", result)
}
