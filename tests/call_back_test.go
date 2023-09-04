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
	"context"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/utils"
)

// 回调服务
func CallBack(r *http.Request) (string, error) {
	//result 需要返回给消息推送端
	result, err := utils.OnCallBack(context.Background(), utils.APIVersionV1, "", "", r, func(ctx context.Context, r *http.Request) {
		
	})
	return result, err
}
