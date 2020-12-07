package database

import (
	"context"
	"fmt"

	"os"
	"time"

	log "github.com/0xkalvin/url-shortener/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitializeMongoDB starts a mongodb connection
func InitializeMongoDB() *mongo.Database {
	logger := log.GetLogger()

	connectionURL := fmt.Sprintf(
		"mongodb://%s:%s@%s",
		os.Getenv("MONGODB_USERNAME"),
		os.Getenv("MONGODB_PASSWORD"),
		os.Getenv("MONGODB_ENDPOINT"),
	)

	logger.Info("Attempting to connect to mongodb...")

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURL))

	if err != nil {
		logger.Errorf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		logger.Errorf("Failed to connect to mongodb cluster: %v", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Errorf("Failed to ping mongodb cluster: %v", err)
	}

	database := client.Database("urlshortener")

	logger.Info("Successfully connect to mongoDB")

	return database
}
