package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	redis "github.com/go-redis/redis/v8"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/models"
)

// UserRepository abstraction
type UserRepository struct {
	Database *dynamodb.DynamoDB
	Cache    *redis.Client
}

// NewUserRepository creates an repository with each database layer
func NewUserRepository(database *dynamodb.DynamoDB, cache *redis.Client) *UserRepository {
	return &UserRepository{
		Database: database,
		Cache:    cache,
	}
}

// Create method persists an user object into the database
func (repository *UserRepository) Create(user *models.User) (*models.User, error) {
	logger := log.GetLogger()

	item, err := dynamodbattribute.MarshalMap(user)

	if err != nil {

		logger.Error("Failed to marshal user attribute")

		return nil, err
	}

	logger.Info("Preparing to store user on dynamoDB")

	inputItem := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Users"),
	}

	_, err = repository.Database.PutItem(inputItem)

	logger.Info("Successfully stored user on dynamoDB")

	return user, nil
}
