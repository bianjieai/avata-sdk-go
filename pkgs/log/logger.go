package log

import (
	log "github.com/sirupsen/logrus"
)

func Logger() *log.Logger {
	logger := log.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(log.ErrorLevel)
	return logger
}
