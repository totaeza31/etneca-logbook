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

func GetTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allTech, err := repository.FindAllTech()
	if err != nil {
		utils.SentNewMessage(response, false, "can not found")
	} else {

		json.NewEncoder(response).Encode(allTech.Tech)
	}
}

func GetTechByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	tech, err := repository.FindTech(objID)
	if err != nil {
		utils.SentNewMessage(response, false, "can not found")
	} else {
		json.NewEncoder(response).Encode(tech)
	}
}

func PostTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var tech models.Tech
	err := json.NewDecoder(request.Body).Decode(&tech)
	tech.ID = primitive.NewObjectID()

	err = repository.InsertTech(tech)
	if err != nil {
		utils.SentNewMessage(response, false, "can not save data")
	} else {
		respond = models.Insert_success()
		utils.SentNewMessage(response, true, "save data success")
	}
}

func PutTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindTech(objID)
	if err != nil {
		utils.SentNewMessage(response, false, "id not found")
	} else {
		var tech models.Tech
		json.NewDecoder(request.Body).Decode(&tech)
		err = repository.UpdateTech(tech, objID)
		if err != nil {
			utils.SentNewMessage(response, false, "update failed")
		} else {
			utils.SentNewMessage(response, true, "update success")
		}
	}
}

func DelTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteTech(objID)
	if err != nil {
		utils.SentNewMessage(response, false, "delete failed")
	} else {
		utils.SentNewMessage(response, true, "delete success")
	}
}