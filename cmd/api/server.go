package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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
