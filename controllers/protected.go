package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var accessToken = tokens.Token
	ID, valid := utils.ParseJson(accessToken)
	if valid == false {
		utils.SentMessage(response, false, "parse token failed")
	} else {
		objID, _ := primitive.ObjectIDFromHex(ID)
		var user models.User
		json.NewDecoder(request.Body).Decode(&user)
		user, err := repository.FindUser(objID)
		if err != nil {
			utils.SentMessage(response, false, "user not found")
		}
		json.NewEncoder(response).Encode(user)
	}

}
