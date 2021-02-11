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
	response.Header().Set("Access-Control-Allow-Origin", "*")
	allTech, err := repository.FindAllTech()
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {

		json.NewEncoder(response).Encode(allTech.Tech)
	}
}

func GetTechByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	tech, err := repository.FindTech(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(tech)
	}
}

func PostTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var tech models.Tech
	err := json.NewDecoder(request.Body).Decode(&tech)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		tech.ID = primitive.NewObjectID()
		err = repository.InsertTech(tech)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}

}

func PutTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindTech(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var tech models.Tech
		json.NewDecoder(request.Body).Decode(&tech)
		err = repository.UpdateTech(tech, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelTech(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteTech(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
