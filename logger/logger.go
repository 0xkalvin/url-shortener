package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GetLogger returns standardized logger
func GetLogger() *logrus.Entry {
	var logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetOutput(os.Stdout)

	loggerWithFields := logger.WithFields(logrus.Fields{
		"service": "url-shortener",
		"env":     os.Getenv("ENV"),
	})

	return loggerWithFields
}
