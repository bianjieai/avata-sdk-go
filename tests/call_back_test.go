package tests

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/utils"
)

// 应用方接收回调实现示例
func CallBack(r *http.Request, w http.ResponseWriter) error {
	// 定义一个闭包函数，用于处理回调操作
	app := func(ctx context.Context, version, apiSecret, path string, r *http.Request) error {
		// 业务逻辑
		return nil
	}

	var err = utils.OnCallBack(context.Background(), utils.APIVersionV1, "1", "", r, w, app)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
