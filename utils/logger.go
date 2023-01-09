package utils

import (
	log "github.com/sirupsen/logrus"
)

func Logger(level log.Level) *log.Logger {
	logger := log.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(level)
	return logger
}
