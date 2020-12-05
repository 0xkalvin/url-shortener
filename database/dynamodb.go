package database

import (
	"os"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// InitializeDynamoDB sets up a connection with dynamoDB servers
func InitializeDynamoDB() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region:   aws.String(os.Getenv("DYNAMODB_REGION")),
		Endpoint: aws.String(os.Getenv("DYNAMODB_ENDPOINT")),
	}

	session := session.Must(session.NewSession(config))

	dynamoDBClient := dynamodb.New(session)

	logger := log.GetLogger()

	logger.Info("DynamoDB is connected")

	return dynamoDBClient
}
