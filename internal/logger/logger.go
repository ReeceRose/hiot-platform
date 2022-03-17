package logger

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

//Instance returns a singleton logrus instance
func Instance() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	return logger
}
