package controllers

import (
	"encoding/json"
	"fmt"

	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllReport(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allReport, err := repository.FindAllReport()
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allReport.Report)
	}
}

func GetReportByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	report, err := repository.FindReport(objID)
	if err != nil {
		fmt.Println(err)
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(report)
	}
}

func PostReport(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var report models.Report
	err := json.NewDecoder(request.Body).Decode(&report)
	report.ID = primitive.NewObjectID()
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		err = repository.InsertReport(report)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutReport(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindReport(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var report models.Report
		json.NewDecoder(request.Body).Decode(&report)
		err = repository.UpdateReport(report, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelReport(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteHuman(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
