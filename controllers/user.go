package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
	"strings"
)

type Token struct {
	Token string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

var tokens Token

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var authen models.Authen
	json.NewDecoder(request.Body).Decode(&authen)
	var password = authen.Password
	authen, err := repository.FindEmail(authen.Email)
	if err != nil {
		utils.SentMessage(response, false, "user not found")
	}
	err = utils.Decrypt(password, authen.Password)
	if err != nil {
		utils.SentMessage(response, false, "invalid password")
	} else {
		authen.AccessToken, err = utils.GenerateToken(authen, "access")
		authen.RefreshToken, err = utils.GenerateToken(authen, "refresh")
		if err != nil {
			utils.SentMessage(response, false, "crete  token error")
		}
		json.NewEncoder(response).Encode(authen)
		tokens.Token = authen.RefreshToken
	}
}

func VerifyAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		authHeader := request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			token, _ := utils.ValidToken(bearerToken[1])
			if token.Valid {
				next.ServeHTTP(response, request)
			} else {
				utils.SentMessage(response, false, "invalid token")
				return
			}
		} else {
			utils.SentMessage(response, false, "invalid token")
			return
		}
	})
}
