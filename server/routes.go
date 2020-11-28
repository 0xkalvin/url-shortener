package server

import (
	"github.com/0xkalvin/url-shortener/controllers"
	"github.com/gin-gonic/gin"
)

func initializeRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	healthController := new(controllers.HealthController)

	router.GET("/_health_check", healthController.Show)

	return router
}
