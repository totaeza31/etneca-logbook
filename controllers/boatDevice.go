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

func GetAllBoatDevice(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allBoatsDevice, err := repository.FindAllBoatDevice()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allBoatsDevice.BoatDevice)
	}
}

func GetBoatDeviceByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	boatDevice, err := repository.FindBoatDevice(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(boatDevice)
	}
}

func PostBoatDevice(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var boatDevice models.BoatDevice
	err := json.NewDecoder(request.Body).Decode(&boatDevice)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertBoatDevice(boatDevice)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutBoatDevice(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindBoatDevice(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var boatDevice models.BoatDevice
		json.NewDecoder(request.Body).Decode(&boatDevice)
		err = repository.UpdateBoatDevice(boatDevice, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelBoatDevice(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteBoatDevice(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
