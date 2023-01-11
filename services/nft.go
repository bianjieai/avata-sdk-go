package services

import (
	"github.com/sirupsen/logrus"

	"avata-sdk-go/utils"
)

type NftService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}
