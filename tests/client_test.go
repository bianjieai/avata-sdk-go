package tests

import (
	"os"
	"testing"
	"time"

	client2 "github.com/bianjieai/avata-sdk-go"
	"github.com/bianjieai/avata-sdk-go/configs"
	"github.com/sirupsen/logrus"
)

var client *client2.AvataClient

func initClient() *client2.AvataClient {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	options := []configs.Options{
		configs.Logger(log),
		configs.HttpTimeout(15 * time.Second),
	}

	client = client2.NewClient("域名", "项目参数 API KEY", "项目参数 API SECRET", options...)
	return client
}

func TestMain(m *testing.M) {
	initClient()
	os.Exit(m.Run())
}
