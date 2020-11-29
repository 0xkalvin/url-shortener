package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// InitializeDynamoDB sets up a connection with dynamoDB servers
func InitializeDynamoDB() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://dynamodb:8000"),
	}

	session := session.Must(session.NewSession(config))

	client := dynamodb.New(session)

	return client
}
