package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/bianjieai/avata-sdk-go/models"
)

var orderID = fmt.Sprintf("orderID_%v", time.Now().Unix())

// 购买能量值/业务费接口示例
func TestCreateOrder(t *testing.T) {
	params := &models.CreateOrderReq{
		Account:   "0x7982C2FEEECCB2A86C5346762AF9DCAC4DF79219",
		Amount:    10100,
		OrderType: "gas",
		OrderId:   orderID,
	}

	result := client.Order.CreateOrder(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var orderRes models.OrderRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &orderRes)

	t.Logf("%+v \n", orderRes)
}

// 查询能量值/业务费购买结果列表接口
func TestQueryOrders(t *testing.T) {
	params := &models.QueryOrdersReq{
		Limit: "1",
	}
	result := client.Order.QueryOrders(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var orderRes models.QueryOrdersRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &orderRes)

	t.Logf("%+v \n", orderRes)
}

// 查询能量值/业务费购买结果接口示例
func TestQueryOrder(t *testing.T) {
	result := client.Order.QueryOrder("orderID_1673342033")
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var orderRes models.QueryOrderRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &orderRes)

	t.Logf("%+v \n", orderRes)
}

// 批量购买能量值接口示例
func TestCreateBatchOrder(t *testing.T) {
	var list []models.List
	list = append(list, models.List{
		Account: "iaa1cz8c3ka0wskwdxdm204jvxrmzxmd3yuy7tm7k9",
		Amount:  100,
	})
	params := &models.BatchCreateOrderReq{
		List:    list,
		OrderId: orderID,
	}

	result := client.Order.BatchCreateOrder(params)
	if result.Code != 0 {
		t.Log(result)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result)
		return
	}
	t.Logf("%+v \n", result.Data)

	var orderRes models.OrderRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &orderRes)

	t.Logf("%+v \n", orderRes)
}
