/*
 * @description:
 * @param:
 * @return:
 */
/*
 * @description:
 * @param:
 * @return:
 */
package tests

import (
	"net/http"

	"github.com/bianjieai/avata-sdk-go/utils"
)

//回调服务
func CallBack(r *http.Request) string {
	//result 需要返回给消息推送端
	result := utils.OnCallBack("", "", "", r, func() {

	})
	return result
}
