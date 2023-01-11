package avata_sdk_go

import (
	"avata-sdk-go/configs"
	"avata-sdk-go/models"
	"avata-sdk-go/services"
	"avata-sdk-go/utils"
)

type AvataClient struct {
	Account services.AccountService
	Tx      services.TxService
	NFT     services.NftService
	MT      services.MtService
	Record  services.RecordService
	Order   services.OrderService
}

func NewClient(domain, apiKey, apiSecret string, options ...configs.Options) *AvataClient {
	// 校验必填参数
	checkBaseParams(domain, apiKey, apiSecret)
	// 设置默认配置
	cfg := configs.SetDefaultConfig(&configs.Config{})
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
		Account: services.AccountService{Logger: logger, HttpClient: httpClient},
		Tx:      services.TxService{Logger: logger, HttpClient: httpClient},
		NFT:     services.NftService{Logger: logger, HttpClient: httpClient},
		MT:      services.MtService{Logger: logger, HttpClient: httpClient},
		Record:  services.RecordService{Logger: logger, HttpClient: httpClient},
		Order:   services.OrderService{Logger: logger, HttpClient: httpClient},
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
