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

	
	var short_url = createShortURL(100000000)

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

func createShortURL(counter uint64) string {
	
	var seed_number uint64 = counter

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var base uint64 = uint64(len(alphabet))
	var hash string = ""

	for seed_number > 0 {
		hash += string(alphabet[seed_number % base])
		seed_number = seed_number / base
	}

	return hash
}
