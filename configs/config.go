package configs

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const httpTimeout = 10

type config struct {
	Level       log.Level     // 日志级别
	HttpTimeout time.Duration // 响应超时时间
}

type Options func(*config)

// SetDefaultConfig 默认配置
func SetDefaultConfig() *config {
	return &config{
		Level:       log.ErrorLevel,
		HttpTimeout: httpTimeout,
	}
}

// Level 日志级别
func Level(level log.Level) Options {
	return func(config *config) {
		config.Level = level
	}
}

// HttpTimeout 响应超时时间
func HttpTimeout(httpTimeout time.Duration) Options {
	return func(config *config) {
		config.HttpTimeout = httpTimeout
	}
}
