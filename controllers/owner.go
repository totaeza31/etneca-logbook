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

func PutOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindOwner(objID)
	if err != nil {
		utils.SentMessage(response, false, "this user not found")
	} else {
		var owner models.Owner
		json.NewDecoder(request.Body).Decode(&owner)
		owner.Birthday += "T00:00:00.000Z"
		owner.Birthday_date, _ = time.Parse("2006-01-02T15:04:05.000Z", owner.Birthday)
		owner.Birthday = ""
		err = repository.UpdateOwer(owner, objID)
		if err != nil {
			utils.SentMessage(response, true, "update error")
		} else {
			utils.SentMessage(response, true, "update success")
		}
	}
}

func DelOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteOwner(objID)
	if err != nil {
		utils.SentMessage(response, false, "delete error")
	} else {
		utils.SentMessage(response, true, "delete success")
	}
}

func PatchOwnerCredit(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindOwner(objID)
	if err != nil {
		utils.SentMessage(response, false, "this user not found")
	} else {
		var owner models.Owner
		json.NewDecoder(request.Body).Decode(&owner)
		err = repository.UpadateOwnerCredit(owner.Credit, objID)
		if err != nil {
			utils.SentMessage(response, true, "update error")
		} else {
			utils.SentMessage(response, true, "update success")
		}
	}
}
