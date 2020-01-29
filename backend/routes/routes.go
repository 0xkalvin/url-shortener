package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener-api/controllers"

)


func SetupRoutes(router *gin.Engine){

	short := new(controllers.ShortController)
	long := new(controllers.LongController)
	
	router.GET("/", indexHandler)
	
	router.POST("/short", short.Create)
	router.GET("/short", short.Index)
	
	router.GET("/long/:short_url", long.Show)
	
}

func indexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Up and kicking",
	})
}