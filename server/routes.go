package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0xkalvin/url-shortener/repositories"

	"github.com/0xkalvin/url-shortener/controllers"
	"github.com/0xkalvin/url-shortener/database"
	"github.com/0xkalvin/url-shortener/middlewares"
	"github.com/0xkalvin/url-shortener/services"
)

func initializeRouter() *gin.Engine {
	dynamoDBClient := database.InitializeDynamoDB()
	redisClient := database.InitializeRedis()

	userRepository := repositories.NewUserRepository(dynamoDBClient, redisClient)

	userService := services.NewUserService(*userRepository)

	healthCheckController := new(controllers.HealthCheckController)
	UserController := controllers.NewUserController(*userService)

	router := gin.New()

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
