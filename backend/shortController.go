package main

import (
	"github.com/gin-gonic/gin"
)

type LongUrl struct {
	Long_url string `json:"long_url" binding:"required"`
}

type ShortController struct {}

func (s ShortController) create (c *gin.Context) {
	var input LongUrl
	
	c.BindJSON(&input)

	var short_url = createShortURL(input.Long_url)

	c.JSON(200, gin.H{ "short_url": short_url , "long_url": input.Long_url })
}

/*	Get all long URLs	*/
func (s ShortController) index (c *gin.Context){
		
	short_url := c.Param("short_url")

	c.JSON(200, gin.H{ "short_url": short_url })
}


func (s ShortController) show (c *gin.Context){
		
	short_url := c.Param("short_url")

	c.JSON(200, gin.H{ "short_url": short_url })
}

func createShortURL(long_url string) string {
	return "short"
}
