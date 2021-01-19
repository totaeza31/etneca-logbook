package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var authen models.Authen
	json.NewDecoder(request.Body).Decode(&authen)
	var password = authen.Password
	authen, err := repository.FindEmail(authen.Email)
	if err != nil {
		utils.SentMessage(response, false, "user not found")
	}
	err = utils.Decrypt(password, authen.Password)
	if err != nil {
		utils.SentMessage(response, false, "invalid password")
	} else {
		authen.AccessToken, err = utils.GenerateToken(authen, "access")
		authen.RefreshToken, err = utils.GenerateToken(authen, "refresh")
		if err != nil {
			utils.SentMessage(response, false, "crete  token error")
		}
		json.NewEncoder(response).Encode(authen)
	}
}
