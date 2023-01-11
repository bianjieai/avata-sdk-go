package services

import (
	"avata-sdk-go/utils"
	"github.com/sirupsen/logrus"
)

type NftService struct {
	*logrus.Logger // 日志
	*utils.HttpClient
}
