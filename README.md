# avata-sdk-go

## 快速开始

### 引入依赖

```
 go get -u github.com/bianjieai/avata-sdk-go
```

### 创建和使用客户端

```
// 非必填参数(不填写将使用默认值)
options := []configs.Options{
		configs.Level(logrus.DebugLevel), // 日志级别
		configs.HttpTimeout(15 * time.Second), // 响应超时时间
	}
client := NewClient("域名", "项目参数 API KEY", "项目参数 API SECRET", options...)
```

- [链账户接口示例代码](./tests/account_test.go)
- [交易结果查询接口示例代码](./tests/tx_test.go)
- [NFT 接口示例代码](./tests/nft_test.go)
- [MT 接口示例代码](./tests/mt_test.go)
- [充值接口示例代码](./tests/order_test.go)
- [存证接口示例代码](./tests/record_test.go)