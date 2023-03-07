package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rwiteshbera/microservices_demo/brokerService/api"
)

func main() {
	logger := log.New(os.Stdout, "broker_api\t", log.LstdFlags)
	broker := api.NewBroker(logger)

	router := mux.NewRouter()

	PostRouter := router.Methods(http.MethodPost).Subrouter()
	PostRouter.HandleFunc("/broker", broker.CallBroker)

	server := &http.Server{
		Addr:         "localhost:9090",
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
