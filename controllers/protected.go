package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
)

func GetProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var users models.User
	json.NewDecoder(request.Body).Decode(&users)
	user, err := repository.FindUser(users.ID)
	if err != nil {
		utils.SentMessage(response, false, "user not found")
	}
	json.NewEncoder(response).Encode(user)
}
