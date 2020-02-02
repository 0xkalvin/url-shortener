package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener-api/src/controllers"

)


func SetupRoutes(router *gin.Engine){

	short := new(controllers.ShortController)
	long := new(controllers.LongController)
	
	router.GET("/health", healthHandler)
	
	router.POST("/short", short.Create)
	router.GET("/short", short.GetAll)
	
	router.GET("/long/:short_url", long.GetOne)
}

func healthHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Up and kicking",
	})
}