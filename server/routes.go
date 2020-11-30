package server

import (
	"net/http"

	"github.com/0xkalvin/url-shortener/controllers"
	"github.com/gin-gonic/gin"
)

func initializeRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	healthCheckController := new(controllers.HealthCheckController)
	UserController := new(controllers.UserController)

	router.GET("/_health_check", healthCheckController.Show)

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
