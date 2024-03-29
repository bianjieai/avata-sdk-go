package tests

import (
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	sdk "github.com/bianjieai/avata-sdk-go/v3"
	"github.com/bianjieai/avata-sdk-go/v3/configs"
)

var client *sdk.AvataClient

func initClient() *sdk.AvataClient {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	options := []configs.Options{
		configs.Logger(log),
		configs.HttpTimeout(15 * time.Second),
	}

	client = sdk.NewClient("域名，不同环境对应不同的域名", "项目参数 API KEY", "项目参数 API SECRET", options...)
	return client
}

func TestMain(m *testing.M) {
	initClient()
	os.Exit(m.Run())
}
