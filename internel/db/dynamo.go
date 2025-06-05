package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/muhamash/gin-serverless/internel/env"
)

func NewDynamoClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(env.GetEnvString("AWS_REGION")),
	}))
	return dynamodb.New(sess)
}
