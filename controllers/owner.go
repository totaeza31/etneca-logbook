package controllers

import (
	"encoding/json"

	"etneca-logbook/helpers"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var respond models.Constants

func GetOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	allOwner, err := repository.FindAllOwner()
	if err != nil {
		respond = models.Get_data_error()
		utils.SentMessage(response, respond)
	} else {
		json.NewEncoder(response).Encode(allOwner.Owner)
	}
}

func GetOwnerByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	owner, err := repository.FindOwner(objID)
	if err != nil {
		respond = models.Get_data_error()
		utils.SentMessage(response, respond)
	} else {
		json.NewEncoder(response).Encode(owner)
	}
}

func PostOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var owner models.Owner
	err := json.NewDecoder(request.Body).Decode(&owner)
	owner.ID = primitive.NewObjectID()
	stringObjectID := owner.ID.Hex()
	if owner.Picture == "" {
		owner.Picture = "default.png"
	} else {
		owner.Picture, err = helpers.SavePicture(owner.Picture, stringObjectID)
		if err != nil {
			respond = models.Save_picture_error()
			utils.SentMessage(response, respond)
		} else {
			owner.Birthday += "T00:00:00.000Z"
			owner.Birthday_date, _ = time.Parse("2006-01-02T15:04:05.000Z", owner.Birthday)
			owner.Birthday = ""
			err = repository.InsertOwner(owner)
			if err != nil {
				respond = models.Insert_error()
				utils.SentMessage(response, respond)
			} else {
				respond = models.Insert_success()
				utils.SentMessage(response, respond)
			}
		}
	}

}

func PutOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindOwner(objID)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		var owner models.Owner
		json.NewDecoder(request.Body).Decode(&owner)
		if owner.Picture == "" {
			owner.Picture = "default.png"
		} else {
			owner.Picture, err = helpers.SavePicture(owner.Picture, id)
			if err != nil {
				respond = models.Save_picture_error()
				utils.SentMessage(response, respond)
			} else {
				owner.Birthday += "T00:00:00.000Z"
				owner.Birthday_date, _ = time.Parse("2006-01-02T15:04:05.000Z", owner.Birthday)
				owner.Birthday = ""
				err = repository.UpdateOwer(owner, objID)
				if err != nil {
					// respond = models.Update_success()
					// utils.SentMessage(response, respond)
				} else {
					// respond = models.Update_success()
					// utils.SentMessage(response, respond)
				}
			}
		}
	}
}

func DelOwner(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	owner, err := repository.FindOwner(objID)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		if owner.Picture != "defalt.png" {
			err := <-helpers.DeleteFile(owner.Picture)
			if err != nil {
				respond = models.Save_picture_error()
				utils.SentMessage(response, respond)
				return
			}
		}
		err = repository.DeleteOwner(objID)
		if err != nil {
			respond = models.Delete_error()
			utils.SentMessage(response, respond)
		} else {
			respond = models.Delete_success()
			utils.SentMessage(response, respond)
		}
	}
}

func PatchOwnerCredit(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindOwner(objID)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		var owner models.Owner
		json.NewDecoder(request.Body).Decode(&owner)
		err = repository.UpadateOwnerCredit(owner.Credit, objID)
		if err != nil {
			// respond = models.Update_error()
			// utils.SentMessage(response, respond)
		} else {
			// respond = models.Update_success()
			// utils.SentMessage(response, respond)
		}
	}
}
