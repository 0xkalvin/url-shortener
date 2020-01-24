package controllers

import (
	"github.com/gin-gonic/gin"
)

type LongUrl struct {
	Long_url string `json:"long_url" binding:"required"`
}

type ShortController struct {}

func (s ShortController) Create (c *gin.Context) {
	var input LongUrl
	
	c.BindJSON(&input)

	var short_url = createShortURL(input.Long_url)

	c.JSON(200, gin.H{ "short_url": short_url , "long_url": input.Long_url })
}

/*	Get all long URLs	*/
func (s ShortController) Index (c *gin.Context){
		
	short_url := c.Param("short_url")

	c.JSON(200, gin.H{ "short_url": short_url })
}


func (s ShortController) Show (c *gin.Context){
		
	short_url := c.Param("short_url")

	c.JSON(200, gin.H{ "short_url": short_url })
}

func createShortURL(long_url string) string {
	seed_number := 0

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base := len(alphabet)
	hash := ""

	for seed_number > 0 {
		hash += string(alphabet[seed_number % base])
		seed_number = seed_number / base
	}

	return hash
}