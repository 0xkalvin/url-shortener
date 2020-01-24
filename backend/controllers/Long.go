package controllers

import (
	"github.com/gin-gonic/gin"
)

type LongController struct {
	
}

func (l LongController) Index (c *gin.Context){
		
	short_url := c.Param("short_url")

	c.JSON(200, gin.H{ "short_url": short_url })
}


func (l LongController) Show (c *gin.Context){
		
	short_url := c.Param("short_url")

	c.JSON(200, gin.H{ "short_url": short_url })
}

