package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bianjieai/avata-sdk-go/v3/models"
	"github.com/bianjieai/avata-sdk-go/v3/utils"
)

// 应用方接收回调实现示例
func Callback(r *http.Request, w http.ResponseWriter) error {
	// 定义一个闭包函数，用于处理回调操作
	app := func(ctx context.Context, version, kind string, res interface{}) error {
		switch version {
		case models.APIVersionV1:
			// V1 版本回调结果
			result := res.(*models.OnCallbackResV1)
			fmt.Println(result)
			nftV1 := models.NftV1{}
			if err := json.Unmarshal([]byte(result.Nft), &nftV1); err != nil {
				return err
			}
			fmt.Println(nftV1.ClassId, nftV1.NftId)
		case models.APIVersionsOther:
			switch kind {
			case models.Native:
				// V2 及以上版本原生模块接口回调结果
				result := res.(*models.OnCallbackResNative)
				fmt.Println(result.Nft.Id)
			case models.EVM:
				// V2 及以上版本 EVM 模块接口回调结果
				result := res.(*models.OnCallbackResEVM)
				fmt.Println(result.Nft.Id)
			}

		}

		// 业务逻辑
		return nil
	}

	var err = utils.OnCallback(context.Background(), models.APIVersionV1, "1", "", r, w, app)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
