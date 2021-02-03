package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBoat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allBoats, err := repository.FindAllBoat()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(allBoats.Boat)
	}
}

func GetBoatByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	boat, err := repository.FindBoat(objID)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(boat)
	}
}

func PostBoat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var boat models.Boat
	err := json.NewDecoder(request.Body).Decode(&boat)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		boat.Anniversary += "T00:00:00.000Z"
		boat.Anniversary_date, _ = time.Parse("2006-01-02T15:04:05.000Z", boat.Anniversary)
		boat.Anniversary = ""

		boat.WarrantyExp += "T00:00:00.000Z"
		boat.WarrantyExp_date, _ = time.Parse("2006-01-02T15:04:05.000Z", boat.WarrantyExp)
		boat.WarrantyExp = ""

		boat.ReportDate += "T00:00:00.000Z"
		boat.ReportDate_date, _ = time.Parse("2006-01-02T15:04:05.000Z", boat.ReportDate)
		boat.ReportDate = ""

		err = repository.InsertBoat(boat)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutBoat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindBoat(objID)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var boat models.Boat

		json.NewDecoder(request.Body).Decode(&boat)
		boat.Anniversary += "T00:00:00.000Z"
		boat.Anniversary_date, _ = time.Parse("2006-01-02T15:04:05.000Z", boat.Anniversary)
		boat.Anniversary = ""

		boat.WarrantyExp += "T00:00:00.000Z"
		boat.WarrantyExp_date, _ = time.Parse("2006-01-02T15:04:05.000Z", boat.WarrantyExp)
		boat.WarrantyExp = ""

		boat.ReportDate += "T00:00:00.000Z"
		boat.ReportDate_date, _ = time.Parse("2006-01-02T15:04:05.000Z", boat.ReportDate)
		boat.ReportDate = ""

		err = repository.UpdateBoat(boat, objID)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelBoat(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteBoat(objID)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}

type boatName struct {
	Text string `json:"text"`
}

func GetBoatByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var name boatName
	err := json.NewDecoder(request.Body).Decode(&name)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		boat, err := repository.FindBoatByName(name.Text)
		if err != nil {
			fmt.Println(err)
			message := models.Get_data_error()
			utils.SentMessage(response, message)
		} else {
			json.NewEncoder(response).Encode(boat.Boat)
		}
	}
}
