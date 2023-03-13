package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("Starting email service: ")
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRT-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// GetROUTER := router.Methods(http.MethodGet).Subrouter()

	server := &http.Server{
		Addr:         "localhost:7001",
		Handler:      c.Handler(router),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}
