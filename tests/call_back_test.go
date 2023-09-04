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
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/utils"
)

// 回调服务
func CallBack(r *http.Request, w http.ResponseWriter) error {
	// 定义一个闭包函数，用于处理回调操作
	app := func(ctx context.Context, a *http.Request) {
		// 业务逻辑
	}
	//result 需要返回给消息推送端
	err := utils.OnCallBack(context.Background(), utils.APIVersionV1, "1", "", r, w, app)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
