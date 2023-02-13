package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "broker-api", log.LstdFlags)
	brokerHandler := NewBroker(logger)

	mux := http.NewServeMux()
	mux.Handle("/api/v1", brokerHandler)

	server := &http.Server{
		Addr:         "localhost:9090",
		Handler:      mux,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()

}
