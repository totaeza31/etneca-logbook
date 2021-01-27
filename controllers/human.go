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
		utils.SentNewMessage(response, false, "can not found")
	} else {
		json.NewEncoder(response).Encode(allHuman.Human)
	}
}

func GetHumanByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	human, err := repository.FindHuman(objID)
	if err != nil {
		utils.SentNewMessage(response, false, "can not found")
	} else {
		json.NewEncoder(response).Encode(human)
	}
}

func PostHuman(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var human models.Human
	err := json.NewDecoder(request.Body).Decode(&human)
	human.ID = primitive.NewObjectID()
	if err != nil {
		message := models.Update_error()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertHuman(human)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}

}

func PutHuman(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindHuman(objID)
	if err != nil {
		utils.SentNewMessage(response, false, "id not found")
	} else {
		var human models.Human
		json.NewDecoder(request.Body).Decode(&human)
		err = repository.UpdateHuman(human, objID)
		if err != nil {
			utils.SentNewMessage(response, false, "update failed")
		} else {
			utils.SentNewMessage(response, true, "update success")
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
		utils.SentNewMessage(response, false, "delete failed")
	} else {
		utils.SentNewMessage(response, true, "delete success")
	}
}
