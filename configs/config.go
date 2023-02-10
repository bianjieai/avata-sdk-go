package configs

import (
	"time"

	"github.com/siddontang/go-log/loggers"
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/avata-sdk-go/utils"
)

const httpTimeout = 10

type config struct {
	HttpTimeout time.Duration    // 响应超时时间
	Logger      loggers.Advanced // 日志
}

type Options func(*config)

// SetDefaultConfig 默认配置
func SetDefaultConfig() *config {
	return &config{
		HttpTimeout: httpTimeout,
		Logger:      utils.Logger(log.InfoLevel),
	}
}

// HttpTimeout 响应超时时间
func HttpTimeout(httpTimeout time.Duration) Options {
	return func(config *config) {
		config.HttpTimeout = httpTimeout
	}
}

// Logger 日志
func Logger(logger loggers.Advanced) Options {
	return func(config *config) {
		config.Logger = logger
	}
}
