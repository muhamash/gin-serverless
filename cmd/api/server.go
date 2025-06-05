package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (app *Application) InitAWS() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		panic(err)
	}

	app.DB = dynamodb.New(sess)
}

func (app *Application) Serve() error {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", app.Port),
		Handler:        app.routes(),
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 3 << 20,
	}

	log.Printf("Starting server at port %d", app.Port)
	return server.ListenAndServe()
}
