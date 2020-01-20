package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"


)

func main() {

	loadConfig()
	initDatabase()
	
	router := gin.Default()
	
	router.Use(cors.Default())
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Up and kicking",
		})
	})

	router.POST("/short", generateShortUrl)
	router.GET("/short/:short_url", getLongUrl)

	router.Run() 
}