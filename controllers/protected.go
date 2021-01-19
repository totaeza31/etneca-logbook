package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/utils"
	"etneca-logbook/repository"
	"net/http"
)

func Profile(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var users models.User
	json.NewDecoder(request.Body).Decode(&users)
	user, err := repository.FindUser(users.ID)
	if err != nil {
		utils.SentMessage(response, false, "user not found")
	}
	json.NewEncoder(response).Encode(user)
}
