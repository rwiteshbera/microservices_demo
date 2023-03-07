package api

import (
	"log"
	"net/http"
)

type Broker struct {
	logger *log.Logger
}

type RequestPayLoad struct {
	Action string      `json:"action:`
	Auth   AuthPayLoad `json:"auth,omitempty"`
	Log    LogPayLoad  `json:"log,omitempty"`
}

type AuthPayLoad struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayLoad struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (broker *Broker) CallBroker(res http.ResponseWriter, req *http.Request) {
	broker.logger.Println("I am a Broker Service")
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = broker.writeJSON(res, http.StatusOK, payload)
}

func (broker *Broker) LogItem(res http.ResponseWriter, entry LogPayLoad) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	broker.writeJSON(res, http.StatusOK, payload)
}

func NewBroker(logger *log.Logger) *Broker {
	return &Broker{logger: logger}
}
