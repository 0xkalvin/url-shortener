package middlewares

import (
	"time"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// Request struct
type Request struct {
	Message   string    `json:"message"`
	URL       string    `json:"url"`
	Method    string    `json:"method"`
	StartTime time.Time `json:"start_time"`
	From      string    `json:"from"`
}

// Response struct
type Response struct {
	Message   string        `json:"message"`
	URL       string        `json:"url"`
	Method    string        `json:"method"`
	StartTime time.Time     `json:"start_time"`
	From      string        `json:"from"`
	Latency   time.Duration `json:"latency"`
	Status    int           `json:"status"`
}

// HTTPLogger for logging requests and responses
func HTTPLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger := log.GetLogger()

		request := new(Request)

		request.Message = "Request received"
		request.URL = context.Request.URL.Path
		request.Method = context.Request.Method
		request.StartTime = time.Now()
		request.From = "Request"

		logger.WithFields(logrus.Fields{
			"request": request,
		}).Info()

		context.Next()

		response := new(Response)

		response.Message = "Request ended"
		response.URL = context.Request.URL.Path
		response.Method = context.Request.Method
		response.StartTime = time.Now()
		response.From = "Response"
		response.Latency = time.Since(request.StartTime)
		response.Status = context.Writer.Status()

		responseLogger := logger.WithFields(logrus.Fields{
			"response": response,
		})

		if response.Status >= 400 && response.Status < 500 {

			responseLogger.Warn()

		} else if response.Status >= 500 {

			responseLogger.Error()

		} else {

			responseLogger.Info()
		}

	}
}
