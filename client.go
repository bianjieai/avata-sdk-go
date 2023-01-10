package avata_sdk_go

import (
	"avata-sdk-go/models"
	"avata-sdk-go/pkgs/configs"
	"avata-sdk-go/pkgs/errors"
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

	return &AvataClient{
		Account: services.AccountService{Logger: logger, BaseParams: baseParams, Config: cfg},
		Tx:      services.TxService{Logger: logger, BaseParams: baseParams},
		NFT:     services.NftService{Logger: logger, BaseParams: baseParams},
		MT:      services.MtService{Logger: logger, BaseParams: baseParams},
		Record:  services.RecordService{Logger: logger, BaseParams: baseParams},
		Order:   services.OrderService{Logger: logger, BaseParams: baseParams},
	}
}

func checkBaseParams(domain, apiKey, apiSecret string) {
	if domain == "" {
		panic(errors.ErrDomain)
	}
	if apiKey == "" {
		panic(errors.ErrAPIKey)
	}
	if apiSecret == "" {
		panic(errors.ErrAPISecret)
	}
}
