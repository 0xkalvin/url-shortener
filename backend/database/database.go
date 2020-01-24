package database


import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"log"
	"os"
	"context"
)

func SetupDatabase() (*mongo.Database) {
    ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	
    if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}
	
	fmt.Println("Connected to MongoDB!")
	
	return client.Database(os.Getenv("MONGO_DATABASE"))
}


func Inject(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}