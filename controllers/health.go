package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController struct
type HealthController struct{}

// Show method
func (h HealthController) Show(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "Up and kicking",
	})
}
