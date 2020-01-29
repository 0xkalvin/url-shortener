package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"context"
	"fmt"
)

type LongController struct {
	
}


func (l LongController) Show (c *gin.Context){
		
	short_url := c.Param("short_url")

	var url URL

	collection := getCollection(c)
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	err := collection.FindOne(ctx, bson.D{{"short_url", short_url}}).Decode(&url)

	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{ "message":  "not found any long URL for "+ short_url  })
		return
	}

	c.JSON(200, gin.H{ "long_url": url.Long })

}

