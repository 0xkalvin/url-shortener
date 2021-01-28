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
	mongoDB := database.InitializeMongoDB()

	redisClient := database.InitializeRedis()

	userRepository := repositories.NewUserRepository(mongoDB)
	shortURLRepository := repositories.NewShortURLRepository(mongoDB, redisClient)

	userService := services.NewUserService(*userRepository)
	shortURLService := services.NewShortURLService(*shortURLRepository, *userRepository)

	healthCheckController := new(controllers.HealthCheckController)
	UserController := controllers.NewUserController(*userService)
	shortURLController := controllers.NewShortURLController(*shortURLService)

	router := gin.New()

	router.GET("/_health_check", healthCheckController.Show)

	router.Use(middlewares.HTTPLogger())

	v1 := router.Group("v1")
	{
		v1.POST("/users", UserController.Create)
		v1.GET("/users/:id", UserController.Show)

		v1.POST("/short_urls", shortURLController.Create)
		v1.GET("/short_urls/:hash", shortURLController.Show)
		v1.GET("/short_urls/:hash/redirect", shortURLController.Redirect)
	}

	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Resource not found",
		})
	})

	return router
}
