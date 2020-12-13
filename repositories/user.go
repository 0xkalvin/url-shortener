package repositories

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/models"
)

// UserRepository abstraction
type UserRepository struct {
	Database *mongo.Database
}

// NewUserRepository creates an repository with each database layer
func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

// Create method persists an user object into the database
func (repository *UserRepository) Create(user *models.User) (*models.User, error) {
	logger := log.GetLogger()

	collection := repository.Database.Collection("users")

	user.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to store user on collection")

		return nil, err
	}

	logger.Info("Successfully stored user on mongoDB")

	return user, nil
}

// FindByID method returns an user from mongo collection if it exists
func (repository *UserRepository) FindByID(objectID primitive.ObjectID) (*models.User, error) {
	logger := log.GetLogger()

	collection := repository.Database.Collection("users")

	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := collection.FindOne(
		ctx,
		bson.M{"_id": objectID},
	).Decode(&user)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to find user on collection")

		return nil, err
	}

	logger.Info("Successfully found user")

	return &user, nil
}
