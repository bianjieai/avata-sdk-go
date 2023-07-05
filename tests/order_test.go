package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/bianjieai/avata-sdk-go/models"
)

var operationID = fmt.Sprintf("orderID_%v", time.Now().Unix())

// 购买能量值/业务费接口示例
func TestCreateOrder(t *testing.T) {
	params := &models.CreateOrderReq{
		Account:     "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
		Amount:      10100,
		OrderType:   1,
		OperationID: operationID,
	}

	result, err := client.Order.CreateOrder(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询能量值/业务费购买结果列表接口
func TestQueryOrders(t *testing.T) {
	params := &models.QueryOrdersReq{
		Limit: "1",
	}
	result, err := client.Order.QueryOrders(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询能量值/业务费购买结果接口示例
func TestQueryOrder(t *testing.T) {
	result, err := client.Order.QueryOrder("orderID_1673342033")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 批量购买能量值接口示例
func TestCreateBatchOrder(t *testing.T) {
	var list []models.List
	list = append(list, models.List{
		Account: "iaa1cz8c3ka0wskwdxdm204jvxrmzxmd3yuxy7tm7k9",
		Amount:  100,
	})
	params := &models.BatchCreateOrderReq{
		List:        list,
		OperationID: operationID,
	}

	result, err := client.Order.BatchCreateOrder(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
