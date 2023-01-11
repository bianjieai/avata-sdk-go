package configs

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Level       log.Level // 日志级别
	HttpTimeout int       // 响应超时时间
}

type Options func(*Config)

// SetDefaultConfig 默认配置
func SetDefaultConfig(config *Config) *Config {
	config.Level = log.ErrorLevel
	config.HttpTimeout = 10
	return config
}

// Level 日志级别
func Level(level log.Level) Options {
	return func(config *Config) {
		config.Level = level
	}
}

// HttpTimeout 响应超时时间
func HttpTimeout(httpTimeout int) Options {
	return func(config *Config) {
		config.HttpTimeout = httpTimeout
	}
}
