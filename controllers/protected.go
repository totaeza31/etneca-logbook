package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type password struct {
	Password        string `json:"password,omitempty" bson:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty" bson:"confirmPassword,omitempty"`
}

func GetProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var accessToken = token.Token
	ID, valid := utils.ParseJson(accessToken)
	if valid == false {
		respond = models.Get_data_error()
		utils.SentMessage(response, respond)
	} else {
		objID, _ := primitive.ObjectIDFromHex(ID)
		user, err := repository.FindUser(objID)
		if err != nil {
			respond = models.Get_data_error()
			utils.SentMessage(response, respond)
		} else {
			user.Password = ""
			json.NewEncoder(response).Encode(user)
		}
	}
}
func GetNewToken(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var authen models.Authen
	var newToken newToken
	var refreshToken = token.Token
	ID, _ := utils.ParseJson(refreshToken)
	objID, _ := primitive.ObjectIDFromHex(ID)
	authen, err := repository.FindAuthen(objID)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		repository.DeleteToken(authen.ID.Hex())
		newToken.AccessToken, err = utils.GenerateToken(authen, "access")
		newToken.RefreshToken, err = utils.GenerateToken(authen, "refresh")
		if err != nil {
			respond = models.Create_token_error()
			utils.SentMessage(response, respond)
		} else {
			json.NewEncoder(response).Encode(newToken)
		}
	}
}

func GetNewPassword(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var authen models.Authen
	json.NewDecoder(request.Body).Decode(&authen)
	_, err := repository.FindEmail(authen.Email)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		path := utils.SentMail(authen.Email)
		var link link
		link.Link = path
		json.NewEncoder(response).Encode(link)
	}
}

func ResetPassword(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var password password
	json.NewDecoder(request.Body).Decode(&password)
	param := mux.Vars(request)
	path := param["email"]
	token, _ := utils.ValidPath(path)
	if token.Valid {
		email, _ := utils.ParseEmail(path)
		_, err := repository.FindEmail(email)
		if err != nil {
			respond = models.User_not_found()
			utils.SentMessage(response, respond)
		} else {
			if password.Password == password.ConfirmPassword {
				password := utils.Encrypt(password.Password)
				err := repository.UpdatePassword(password, email)
				if err != nil {
					respond = models.Change_password_error()
					utils.SentMessage(response, respond)
				} else {
					respond = models.Change_password_success()
					utils.SentMessage(response, respond)
				}
			} else {
				respond = models.Password_not_match()
				utils.SentMessage(response, respond)
			}
		}
	} else {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	}

}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repository.DeleteUser(objID)
	if err != nil {
		respond = models.Delete_error()
		utils.SentMessage(response, respond)
	} else {
		respond = models.Delete_success()
		utils.SentMessage(response, respond)
	}
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := repository.FindAuthen(objID)
	if err != nil {
		respond = models.User_not_found()
		utils.SentMessage(response, respond)
	} else {
		var user models.User
		json.NewDecoder(request.Body).Decode(&user)
		err = repository.UpdateUser(user, objID)
		if err != nil {
			fmt.Println(err)
		} else {
			respond = models.Update_success()
			utils.SentMessage(response, respond)
		}
	}
}

func GetPackage(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var packages models.Data
	packages, err := repository.GetPackageAllPackage()
	if err != nil {
		respond = models.Get_data_success()
		utils.SentMessage(response, respond)
	} else {
		json.NewEncoder(response).Encode(packages)
	}

}
