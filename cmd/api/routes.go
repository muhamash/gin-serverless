package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Application) routes() http.Handler {
	g := gin.Default()

	// CORS setup
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8088", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := g.Group("/api/v1")
	{
		api.GET("/", app.handlers.GetUser)
		api.POST("/", app.handlers.CreateUser)
		api.PUT("/", app.handlers.UpdateUser)
		api.DELETE("/", app.handlers.DeleteUser)
	}

	return g
}