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

func GetTitles(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	allTitle, err := repository.FindAllTitle()
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {

		json.NewEncoder(response).Encode(allTitle.Title)
	}
}

func GetTitle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	title, err := repository.FindTitle(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(title)
	}
}

func PostTitle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var title models.Title
	err := json.NewDecoder(request.Body).Decode(&title)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		title.ID = primitive.NewObjectID()
		err = repository.InsertTitle(title)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}

}

func PutTitle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindTitle(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var title models.Title
		json.NewDecoder(request.Body).Decode(&title)
		err = repository.UpdateTitle(title, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelTitle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteTitle(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
