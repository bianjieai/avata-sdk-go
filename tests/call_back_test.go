package tests

import (
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/utils"
)

//回调服务
//v1版本
func CallBack1() {
	http.HandleFunc("路由", getRequest1)
	http.ListenAndServe("端口", nil)
}
func getRequest1(w http.ResponseWriter, r *http.Request) {
	a := utils.CallBackV1(r, "apiSecret")
	if a == "SUCCESS" {
		//该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
		//TODO
	}
	_, err := fmt.Fprintf(w, a)
	if err != nil {
		fmt.Print(err)
	}
}

//v2版本
func CallBack2() {
	http.HandleFunc("路由", getRequest2)
	http.ListenAndServe("端口", nil)
}
func getRequest2(w http.ResponseWriter, r *http.Request) {
	a := utils.CallBackV2(r, "路由地址", "apiSecret")
	if a == "SUCCESS" {
		//该笔推送消息属于文昌链上链完成所推送消息，请及时存储数据
		//TODO
	}
	_, err := fmt.Fprintf(w, a)
	if err != nil {
		fmt.Print(err)
	}
}
