package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Port     int
	DB       *dynamodb.DynamoDB
	IsLambda bool
	handlers struct {
		GetUser    gin.HandlerFunc
		CreateUser gin.HandlerFunc
		UpdateUser gin.HandlerFunc
		DeleteUser gin.HandlerFunc
	}
}


func (app *Application) Handler() {
	ginLambda := ginadapter.New(app.routes().(*gin.Engine))
	lambda.Start(ginLambda.ProxyWithContext)
}

func main() {
	app := &Application{
		Port:      8080,
	}
	
	app.InitAWS()

	if app.IsLambda {
		lambda.Start(app.Handler)
	} else {
		if err := app.Serve(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}
