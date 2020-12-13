package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/schemas"
	"github.com/0xkalvin/url-shortener/services"
)

// ShortURLController struct
type ShortURLController struct {
	Service services.ShortURLService
}

// NewShortURLController creates an controller layer
func NewShortURLController(service services.ShortURLService) *ShortURLController {
	return &ShortURLController{
		Service: service,
	}
}

// Create an URL
func (controller ShortURLController) Create(context *gin.Context) {
	var urlPayload schemas.ShortURLPostSchema

	err := context.Bind(&urlPayload)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error_type": "Invalid request body"},
		)

		return
	}

	createdURL, err := controller.Service.CreateURL(urlPayload)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error_type": "Internal server error"},
		)

		return
	}
	logger := log.GetLogger()

	logger.Info("Finished to create URL")

	context.JSON(http.StatusOK, createdURL)
}

// Show returns an URL
func (controller ShortURLController) Show(context *gin.Context) {
	logger := log.GetLogger()

	hash := context.Param("hash")

	originalURL, err := controller.Service.FindOneURLByHash(hash)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error_type": "Internal server error"},
		)

		return
	}

	logger.Info("Finished to find URL")

	context.JSON(
		http.StatusOK,
		gin.H{"original_url": originalURL},
	)
}
