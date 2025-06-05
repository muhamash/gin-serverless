package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/muhamash/gin-serverless/internal/env"
)

func main() {
	region := env.GetEnvString("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Use awsSession here
	log.Println("AWS session created successfully:", awsSession)
}
