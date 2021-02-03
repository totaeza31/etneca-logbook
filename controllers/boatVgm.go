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

func GetAllBoatVgm(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allBoatsVgm, err := repository.FindAllBoatVgm()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allBoatsVgm.BoatVgm)
	}
}

func GetBoatVgmByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	boatVgm, err := repository.FindBoatVgm(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(boatVgm)
	}
}

func PostBoatVgm(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var boatVgm models.BoatVgm
	err := json.NewDecoder(request.Body).Decode(&boatVgm)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertBoatVgm(boatVgm)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutBoatVgm(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindBoatVgm(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var boatVgm models.BoatVgm
		json.NewDecoder(request.Body).Decode(&boatVgm)
		err = repository.UpdateBoatVgm(boatVgm, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelBoatVgm(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteBoatVgm(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
