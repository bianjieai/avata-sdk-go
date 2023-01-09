package avata_sdk_go

import (
	"avata-sdk-go/models"
	"avata-sdk-go/pkgs/log"
	"avata-sdk-go/services"
)

type AvataClient struct {
	Account services.AccountService
	Tx      services.TxService
	NFT     services.NftService
	MT      services.MtService
	Record  services.RecordService
	Order   services.OrderService
}

func NewClient(domain, apiKey, apiSecret string) *AvataClient {
	logger := log.Logger()

	baseParams := models.BaseParams{
		Domain:    domain,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	return &AvataClient{
		Account: services.AccountService{Logger: logger, BaseParams: baseParams},
		Tx:      services.TxService{Logger: logger, BaseParams: baseParams},
		NFT:     services.NftService{Logger: logger, BaseParams: baseParams},
		MT:      services.MtService{Logger: logger, BaseParams: baseParams},
		Record:  services.RecordService{Logger: logger, BaseParams: baseParams},
		Order:   services.OrderService{Logger: logger, BaseParams: baseParams},
	}
}
