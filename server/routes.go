package server

import (
	"net/http"

	"github.com/0xkalvin/url-shortener/controllers"
	"github.com/0xkalvin/url-shortener/middlewares"

	"github.com/gin-gonic/gin"
)

func initializeRouter() *gin.Engine {
	router := gin.New()

	healthCheckController := new(controllers.HealthCheckController)
	UserController := new(controllers.UserController)

	router.GET("/_health_check", healthCheckController.Show)

	router.Use(middlewares.HTTPLogger())

	v1 := router.Group("v1")
	{
		v1.POST("/users", UserController.Create)
	}

	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Resource not found",
		})
	})

	return router
}
