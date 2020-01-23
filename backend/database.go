package main


import (
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"log"
	"os"
	"context"
)

func setupDatabase() (*mongo.Database) {
    ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	
    if err != nil {
        log.Fatal(err)
	}
	
	fmt.Println("Connected to MongoDB!")
	
	return client.Database(os.Getenv("MONGO_DATABASE"))
}



// func initDatabase(ctx context.Context) *mongo.Client {

// 	// Set client options
// 	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	return client
// }