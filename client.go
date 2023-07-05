package sdk

import (
	"github.com/bianjieai/avata-sdk-go/configs"
	"github.com/bianjieai/avata-sdk-go/models"
	"github.com/bianjieai/avata-sdk-go/services"
	"github.com/bianjieai/avata-sdk-go/utils"
)

type AvataClient struct {
	Account  services.AccountService
	Tx       services.TxService
	NFT      services.NFTService
	Record   services.RecordService
	Order    services.OrderService
	NS       services.NSService
	Contract services.ContractService
	MT       services.MTService
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

	baseParams := models.BaseParams{
		Domain:    domain,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	// 初始化
	httpClient := utils.NewHttpClient(cfg.HttpTimeout, baseParams)

	return &AvataClient{
		Account:  services.NewAccountService(cfg.Logger, httpClient),
		Tx:       services.NewTxService(cfg.Logger, httpClient),
		NFT:      services.NewNFTService(cfg.Logger, httpClient),
		Record:   services.NewRecordService(cfg.Logger, httpClient),
		Order:    services.NewOrderService(cfg.Logger, httpClient),
		NS:       services.NewNSService(cfg.Logger, httpClient),
		Contract: services.NewContractService(cfg.Logger, httpClient),
		MT:       services.NewMTService(cfg.Logger, httpClient),
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
