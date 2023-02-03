# avata-sdk-go

## 快速开始

### 引入依赖

```
 go get -u github.com/bianjieai/avata-sdk-go
```

### 创建和使用客户端

```
import (
	"time"

	sdk "github.com/bianjieai/avata-sdk-go"
	"github.com/bianjieai/avata-sdk-go/configs"
	"github.com/sirupsen/logrus"
)

// 非必填参数(不填写将使用默认值)
log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	options := []configs.Options{
		configs.Logger(log), // 日志
		configs.HttpTimeout(15 * time.Second), // 响应超时时间
	}

	client = sdk.NewClient("域名", "项目参数 API KEY", "项目参数 API SECRET", options...)
```

- [创建客户端示例代码](./tests/client_test.go)
- [链账户接口示例代码](./tests/account_test.go)
- [交易结果查询接口示例代码](./tests/tx_test.go)
- [NFT 接口示例代码](./tests/nft_test.go)
- [MT 接口示例代码](./tests/mt_test.go)
- [充值接口示例代码](./tests/order_test.go)
- [存证接口示例代码](./tests/record_test.go)