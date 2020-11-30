package middlewares

import (
	"time"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// HTTPLogger for logging requests and responses
func HTTPLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()

		baseFields := logrus.Fields{
			"url":        context.Request.URL.Path,
			"method":     context.Request.Method,
			"start_time": startTime,
		}

		baseRequestlogger := log.GetLogger().WithFields(baseFields)

		baseRequestlogger.WithFields(logrus.Fields{
			"from": "Request",
		}).Info("Request received")

		context.Next()

		status := context.Writer.Status()

		responseLogger := baseRequestlogger.WithFields(logrus.Fields{
			"from":    "Response",
			"latency": time.Since(startTime),
			"status":  status,
		})

		if status >= 400 && status < 500 {

			responseLogger.Warn("Request ended with client error")

		} else if status >= 500 {

			responseLogger.Error("Request ended with internal error")

		} else {

			responseLogger.Info("Request ended successfully")
		}

	}
}
