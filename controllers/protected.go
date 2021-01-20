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

type password struct {
	Password        string `json:"password,omitempty" bson:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty" bson:"confirmPassword,omitempty"`
}

func GetProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var accessToken = token.Token
	ID, valid := utils.ParseJson(accessToken)
	if valid == false {
		utils.SentMessage(response, false, "parse token failed")
	} else {
		objID, _ := primitive.ObjectIDFromHex(ID)
		user, err := repository.FindUser(objID)
		if err != nil {
			utils.SentMessage(response, false, "user not found")
		}
		user.Password = ""
		json.NewEncoder(response).Encode(user)
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
		utils.SentMessage(response, false, "not user found")
	} else {
		repository.DeleteToken(authen.ID.Hex())
		newToken.AccessToken, err = utils.GenerateToken(authen, "access")
		newToken.RefreshToken, err = utils.GenerateToken(authen, "refresh")
		if err != nil {
			utils.SentMessage(response, false, "crete token error")
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
		utils.SentMessage(response, false, "email not found")
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
			utils.SentMessage(response, false, "email not found")
		} else {
			if password.Password == password.ConfirmPassword {
				password := utils.Encrypt(password.Password)
				err := repository.UpdatePassword(password, email)
				if err != nil {
					utils.SentMessage(response, false, "change password error")
				} else {
					utils.SentMessage(response, true, "change password success")
				}
			} else {
				utils.SentMessage(response, false, "Password do not match")
			}
		}
	} else {
		utils.SentMessage(response, false, "path error")
	}

}
