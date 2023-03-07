package handlers

import (
	"net/http"

	"github.com/rwiteshbera/microservices_demo/loggerService/database"
	"github.com/rwiteshbera/microservices_demo/loggerService/helpers"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func WriteLog(res http.ResponseWriter, req *http.Request) {
	// Read JSON
	var requestPayload JSONPayload
	helpers.ReadJSON(res, req, &requestPayload)

	// Insert Data
	event := database.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := database.Insert(event)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInsufficientStorage)
		return
	}

	response := JSONResponse{
		Error:   false,
		Message: "logged",
	}

	helpers.WriteJSON(res, http.StatusOK, response)

}
