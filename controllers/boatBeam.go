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

func GetAllBoatBeam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allBoatsBeam, err := repository.FindAllBoatBeam()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allBoatsBeam.BoatBeam)
	}
}

func GetBoatBeamByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	boatBeam, err := repository.FindBoatBeam(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(boatBeam)
	}
}

func PostBoatBeam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var boatBeam models.BoatBeam
	err := json.NewDecoder(request.Body).Decode(&boatBeam)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertBoatBeam(boatBeam)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutBoatBeam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindBoatBeam(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var boatBeam models.BoatBeam
		json.NewDecoder(request.Body).Decode(&boatBeam)
		err = repository.UpdateBoatBeam(boatBeam, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelBoatBeam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteBoatBeam(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
