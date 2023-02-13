package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data, omitempty"`
}

type Broker struct {
	logger *log.Logger
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	broker.logger.Println("I am a Broker Service")
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	rw.Write(out)
}

func NewBroker(logger *log.Logger) *Broker {
	return &Broker{logger: logger}
}
