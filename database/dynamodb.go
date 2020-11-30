package database

import (
	"log"
	"os"

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

	client := dynamodb.New(session)

	log.Println("DynamoDB is connected")

	return client
}
