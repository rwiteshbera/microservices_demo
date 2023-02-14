package main

import (
	"github.com/rwiteshbera/microservices_demo/productService/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "broker_api\t", log.LstdFlags)
	brokerHandler := handlers.NewProducts(logger)

	mux := http.NewServeMux()
	mux.Handle("/api/products", brokerHandler)

	server := &http.Server{
		Addr:         "localhost:9091",
		Handler:      mux,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic("failed to run the broker service: ", err.Error())
	}
}
