package avata_sdk_go

import (
	"github.com/bianjieai/avata-sdk-go/configs"
	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/services"
	"github.com/bianjieai/avata-sdk-go/utils"
)

type AvataClient struct {
	Account services.AccountService
	Tx      services.TxService
	NFT     services.NFTService
	MT      services.MTService
	Record  services.RecordService
	Order   services.OrderService
}

func NewClient(domain, apiKey, apiSecret string, options ...configs.Options) *AvataClient {
	// 校验必填参数
	checkBaseParams(domain, apiKey, apiSecret)
	// 设置默认配置
	cfg := configs.SetDefaultConfig()
	// 遍历调用函数，调整配置
	for _, option := range options {
		option(cfg)
	}
	// 初始化日志
	logger := utils.Logger(cfg.Level)

	baseParams := models.BaseParams{
		Domain:    domain,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	// 初始化
	httpClient := utils.NewHttpClient(cfg.HttpTimeout, baseParams)

	return &AvataClient{
		Account: services.NewAccountService(logger, httpClient),
		Tx:      services.NewTxService(logger, httpClient),
		NFT:     services.NewNFTService(logger, httpClient),
		MT:      services.NewMTService(logger, httpClient),
		Record:  services.NewRecordService(logger, httpClient),
		Order:   services.NewOrderService(logger, httpClient),
	}
}

func checkBaseParams(domain, apiKey, apiSecret string) {
	if domain == "" {
		panic(models.ErrDomain)
	}
	if apiKey == "" {
		panic(models.ErrAPIKey)
	}
	if apiSecret == "" {
		panic(models.ErrAPISecret)
	}
}
