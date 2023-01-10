package services

import (
	"avata-sdk-go/models"
	"github.com/sirupsen/logrus"
)

type NftService struct {
	*logrus.Logger    // 日志
	models.BaseParams // 域名和项目参数
}
