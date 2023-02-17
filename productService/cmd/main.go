package main

import (
	"github.com/gorilla/mux"
	"github.com/rwiteshbera/microservices_demo/productService/handlers"
	"github.com/rwiteshbera/microservices_demo/productService/middlewares"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product_api\t", log.LstdFlags)
	//Create the handler
	productHandler := handlers.NewProductHandler(logger)

	// Create a serve mux and register the handler
	router := mux.NewRouter()

	// Get SubRouter to handle GET Request
	GetRouter := router.Methods(http.MethodGet).Subrouter()
	GetRouter.HandleFunc("/api/products", productHandler.GetProducts)

	// Post SubRouter to handle POST
	PostRouter := router.Methods(http.MethodPost).Subrouter()
	PostRouter.HandleFunc("/api/products", productHandler.AddProduct)
	PostRouter.Use(middlewares.MiddlewareProductValidator)

	GetRouter.HandleFunc("/{id:[0-9]+}", productHandler.GetProduct)
	GetRouter.Use(middlewares.MiddlewareProductValidator)

	server := &http.Server{
		Addr:         "localhost:9091",
		Handler:      router,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic("failed to run the broker service: ", err.Error())
	}
}
