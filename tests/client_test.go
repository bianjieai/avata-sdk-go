package tests

import (
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	sdk "github.com/bianjieai/avata-sdk-go"
	"github.com/bianjieai/avata-sdk-go/configs"
)

var client *sdk.AvataClient

func initClient() *sdk.AvataClient {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	options := []configs.Options{
		configs.Logger(log),
		configs.HttpTimeout(15 * time.Second),
	}

	client = sdk.NewClient("http://openapi.avata.qa.nj.bianjie.ai", "J2r350y6U2F8T1X0d0L534a819W2w0Zf", "2243Y0B642A8a1K0J0y5F4q8M90290B5", options...)
	return client
}

func TestMain(m *testing.M) {
	initClient()
	os.Exit(m.Run())
}
