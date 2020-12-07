package repositories

import (
	"context"

	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/models"
)

// UserRepository abstraction
type UserRepository struct {
	Database *mongo.Database
	Cache    *redis.Client
}

// NewUserRepository creates an repository with each database layer
func NewUserRepository(database *mongo.Database, cache *redis.Client) *UserRepository {
	return &UserRepository{
		Database: database,
		Cache:    cache,
	}
}

// Create method persists an user object into the database
func (repository *UserRepository) Create(user *models.User) (*models.User, error) {
	logger := log.GetLogger()

	collection := repository.Database.Collection("posts")

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to store user on collection")
	}

	objectID, _ := insertResult.InsertedID.(primitive.ObjectID)

	user.ID = objectID.Hex()

	logger.Info("Successfully stored user on mongoDB")

	return user, nil
}
