package controllers

import (
	"url-shortener-api/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"time"
	"os"
	"context"
)

type URL = models.URL
var counter_seed uint64 = 100000000

type ShortController struct {}

func (s ShortController) Create (c *gin.Context) {
	
	var url URL
	
	c.BindJSON(&url)

	collection := getCollection(c)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"long_url", url.Long}}).Decode(&url)
	
	if err != nil {
		fmt.Println(err)
		
		url.Short = createShortURL(counter_seed)
		counter_seed++
		url.CreatedAt = time.Now().Unix()
	
		result, _ := collection.InsertOne(ctx, url)
		fmt.Println(result)
	}

	c.JSON(200, gin.H{ "short_url": url.Short })
	return
}


func (s ShortController) GetAll (c *gin.Context){
		
	short_url := c.Param("short_url")
	fmt.Println(short_url)


	collection := getCollection(c)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
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

