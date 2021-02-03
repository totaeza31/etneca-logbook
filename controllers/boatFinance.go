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

func GetAllBoatFinance(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allBoatsFinance, err := repository.FindAllBoatFinance()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allBoatsFinance.BoatFinance)
	}
}

func GetBoatFinanceByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	boatFinance, err := repository.FindBoatFinance(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(boatFinance)
	}
}

func PostBoatFinance(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var boatFinance models.BoatFinance
	err := json.NewDecoder(request.Body).Decode(&boatFinance)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertBoatFinance(boatFinance)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutBoatFinance(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindBoatFinance(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var boatFinance models.BoatFinance
		json.NewDecoder(request.Body).Decode(&boatFinance)
		err = repository.UpdateBoatFinance(boatFinance, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelBoatFinance(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteBoatFinance(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
