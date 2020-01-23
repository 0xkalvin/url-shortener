package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"fmt"
)

var counter int


func main() {

	loadConfig()
	
	db := setupDatabase()
	
	fmt.Println(db)
	counter = 0

	short := new(ShortController)
	long := new(LongController)
	
	router := gin.Default()
	
	router.Use(cors.Default())
		
	router.GET("/",indexHandler)
	router.POST("/short", short.create)
	router.GET("/short", short.index)
	router.GET("/long/:short_url", long.show)
	
	router.NoRoute(notFoundHandler)

	router.Run() 
}

func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Page not found"})
}

func indexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Up and kicking",
	})
}