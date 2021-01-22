package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allOwner, err := repository.FindAllOwner()
	if err != nil {
		utils.SentMessage(response, false, "get data error")
	} else {
		var message models.MessageAllOwner
		message.AllOwner = allOwner
		message.Result = true
		message.Message = "get data success"
		json.NewEncoder(response).Encode(message)
	}
}

func GetOwnerByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	owner, err := repository.FindOwner(objID)
	if err != nil {
		utils.SentMessage(response, false, "get data error")
	} else {
		var message models.MessageOwner
		message.Owner = owner
		message.Result = true
		message.Message = "get data success"
		json.NewEncoder(response).Encode(message)
	}
}

func PostOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var owner models.Owner
	err := json.NewDecoder(request.Body).Decode(&owner)
	owner.Birthday += "T00:00:00.000Z"
	owner.Birthday_date, _ = time.Parse("2006-01-02T15:04:05.000Z", owner.Birthday)
	owner.Birthday = ""
	err = repository.InsertOwner(owner)
	if err != nil {
		utils.SentMessage(response, false, "Insert Error")
	} else {
		utils.SentMessage(response, true, "Insert Success")
	}
}
