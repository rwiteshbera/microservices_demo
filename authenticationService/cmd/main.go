package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rwiteshbera/microservices_demo/authenticationService/api"
	"github.com/rwiteshbera/microservices_demo/authenticationService/handlers"
)

func main() {
	log.Println("Starting authentication service")

	router := mux.NewRouter()

	PostROUTER := router.Methods(http.MethodPost).Subrouter()
	PostROUTER.HandleFunc("/signup", handlers.SignupRouter)

	server := &http.Server{
		Addr:         api.AuthHandlerInstance.Env.SERVER_HOST + ":" + api.AuthHandlerInstance.Env.SERVER_PORT,
		Handler:      router,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
