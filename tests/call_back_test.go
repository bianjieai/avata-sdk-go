/*
 * @description:
 * @param:
 * @return:
 */
package tests

import (
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/utils"
)

//回调服务
func CallBack(r *http.Request) {
	a := utils.ReceiveMessages("", "", "", r)
	fmt.Print(a)
}
