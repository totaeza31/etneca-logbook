package controllers

import (
	"encoding/json"

	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllHuman(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allHuman, err := repository.FindAllHuman()
	if err != nil {
		respond = models.Get_data_error()
		utils.SentMessage(response, respond)
	} else {
		var message models.MessageAllHuman
		message.AllHuman = allHuman
		message.Result = true
		message.Message = "get data success"
		json.NewEncoder(response).Encode(message)
	}
}

func GetHumanByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	human, err := repository.FindHuman(objID)
	if err != nil {
		respond = models.Get_data_error()
		utils.SentMessage(response, respond)
	} else {
		var message models.MessageHuman
		message.Human = human
		message.Result = true
		message.Message = "get data success"
		json.NewEncoder(response).Encode(message)
	}
}

func PostHuman(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var human models.Human
	err := json.NewDecoder(request.Body).Decode(&human)
	human.ID = primitive.NewObjectID()

	err = repository.InsertHuman(human)
	if err != nil {
		respond = models.Insert_error()
		utils.SentMessage(response, respond)
	} else {
		respond = models.Insert_success()
		utils.SentMessage(response, respond)
	}
}

func PutHuman(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindHuman(objID)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		var human models.Human
		json.NewDecoder(request.Body).Decode(&human)
		err = repository.UpdateHuman(human, objID)
		if err != nil {
			respond = models.Update_success()
			utils.SentMessage(response, respond)
		} else {
			respond = models.Update_success()
			utils.SentMessage(response, respond)
		}
	}
}

func DelHuman(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteHuman(objID)
	if err != nil {
		respond = models.Delete_error()
		utils.SentMessage(response, respond)
	} else {
		respond = models.Delete_success()
		utils.SentMessage(response, respond)
	}
}
