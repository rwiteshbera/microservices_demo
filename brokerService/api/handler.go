package api

import (
	"log"
	"net/http"
)

type Broker struct {
	logger *log.Logger
}

func (broker *Broker) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	broker.logger.Println("I am a Broker Service")
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = broker.writeJSON(res, http.StatusOK, payload)
}

func NewBroker(logger *log.Logger) *Broker {
	return &Broker{logger: logger}
}
