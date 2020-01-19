package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
)


type LongUrl struct{
    Long_url string `json:"long_url" binding:"required"`
}


func createShortURL(long_url string) string {
	return "short"
}


func main() {
	router := gin.Default()

	router.Use(cors.Default())
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Up and kicking",
		})
	})

	router.POST("/short", func(c *gin.Context) {
		var input LongUrl
		
		c.BindJSON(&input)

		var short_url = createShortURL(input.Long_url)
		
		c.JSON(200, gin.H{ "short_url": short_url , "long_url": input.Long_url })
	})



	router.Run() 
}