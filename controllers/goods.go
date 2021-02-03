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

func GetAllGoods(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allGoods, err := repository.FindAllGoods()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allGoods.Goods)
	}
}

func GetGoodsByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	goods, err := repository.FindGoods(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(goods)
	}
}

func PostGoods(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var goods models.Goods
	err := json.NewDecoder(request.Body).Decode(&goods)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertGoods(goods)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutGoods(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindGoods(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var goods models.Goods
		json.NewDecoder(request.Body).Decode(&goods)
		err = repository.UpdateGoods(goods, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelGoods(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteGoods(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
