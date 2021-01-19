package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"net/http"
)

func Login(response http.ResponseWriter, request *http.Request) {
	var message models.Message
	message.Message = "Hello"
	message.Result = true

	json.NewEncoder(response).Encode(message)
}
