package services

import (
	"github.com/sirupsen/logrus"

	"avata-sdk-go/utils"
)

type NFTService interface {
}

type nftService struct {
	*logrus.Logger // 日志
	utils.HttpClient
}

func NewNFTService(log *logrus.Logger, httpClient utils.HttpClient) *nftService {
	return &nftService{
		Logger:     log,
		HttpClient: httpClient,
	}
}
