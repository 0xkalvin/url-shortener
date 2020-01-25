package controllers

import (
	"url-shortener-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"time"
	"os"
	"context"
)

type URL = models.URL


type ShortController struct {}

func (s ShortController) Create (c *gin.Context) {
	
	var url URL
	
	c.BindJSON(&url)

	collection := getCollection(c)
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	url.Short = createShortURL(100000000)
	url.CreatedAt = time.Now().Unix()

	result, err := collection.InsertOne(ctx, url)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, result)
}


func (s ShortController) Show (c *gin.Context){
		
	short_url := c.Param("short_url")
	fmt.Println(short_url)


	collection := getCollection(c)
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	

	cursor, err := collection.Find(ctx, bson.D{{}})

	urls := []URL{}
	if err == nil {
		for cursor.Next(context.Background()) {
			url := URL{}
			cursor.Decode(&url)
			urls = append(urls, url)
		}
	} else {
		print(err.Error())
	}

	c.JSON(200, urls)
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



func getCollection(c *gin.Context) (*mongo.Collection){

	db := c.MustGet("db").(*mongo.Database)
	collection := db.Collection(os.Getenv("MONGO_COLLECTION"))

	return collection
}

