package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/muhamash/gin-serverless/cmd/api/handlers"
	"github.com/muhamash/gin-serverless/internel/db"
	"github.com/muhamash/gin-serverless/internel/utils"
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
		DB:   db.NewDynamoClient(),
	}
	
	// app.DB = db.NewDynamoClient()

	app.handlers = struct {
		GetUser    gin.HandlerFunc
		CreateUser gin.HandlerFunc
		UpdateUser gin.HandlerFunc
		DeleteUser gin.HandlerFunc
	}{
		GetUser:    utils.GinifyHandler(handlers.GetUser, app.DB),
		CreateUser: utils.GinifyHandler(handlers.CreateUser, app.DB),
		UpdateUser: utils.GinifyHandler(handlers.UpdateUser, app.DB),
		DeleteUser: utils.GinifyHandler(handlers.DeleteUser, app.DB),
	}
	

	if app.IsLambda {
		lambda.Start(app.Handler)
	} else {
		if err := app.Serve(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}
