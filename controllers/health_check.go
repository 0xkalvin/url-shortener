package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckController struct
type HealthCheckController struct{}

// Show method
func (h HealthCheckController) Show(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "Up and kicking",
	})
}
