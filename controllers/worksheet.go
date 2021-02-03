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

func GetAllWorksheet(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	// allBoatsDevice, err := repository.FindAllWorksheet()

	// if err != nil {
	// 	message := models.Get_data_error()
	// 	utils.SentMessage(response, message)
	// } else {
	// 	// json.NewEncoder(response).Encode(allBoatsDevice.BoatDevice)
	// }
}

func GetWorksheetByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	boatDevice, err := repository.FindWorksheet(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(boatDevice)
	}
}

func PostWorksheet(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var boatDevice models.BoatDevice
	err := json.NewDecoder(request.Body).Decode(&boatDevice)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		// err = repository.InsertWorksheet(boatDevice)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutWorksheet(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindWorksheet(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var boatDevice models.BoatDevice
		json.NewDecoder(request.Body).Decode(&boatDevice)
		// err = repository.UpdateWorksheet(boatDevice, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelWorksheet(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteWorksheet(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
