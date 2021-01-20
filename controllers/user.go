package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	Token string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

type newToken struct {
	AccessToken  string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

type link struct {
	Link string `json:"link,omitempty" bson:"link,omitempty"`
}

var token Token

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var authen models.Authen
	err := json.NewDecoder(request.Body).Decode(&authen)
	if err != nil {
		utils.SentMessage(response, false, "invalid syntax")
	} else {
		var password = authen.Password
		authen, err = repository.FindEmail(authen.Email)
		if err != nil {
			utils.SentMessage(response, false, "user not found")
		} else {
			err = utils.Decrypt(password, authen.Password)
			if err != nil {
				utils.SentMessage(response, false, "invalid password")
			} else {
				authen.AccessToken, err = utils.GenerateToken(authen, "access")
				authen.RefreshToken, err = utils.GenerateToken(authen, "refresh")
				var rs models.RespondAuthen
				authen.Password = ""
				rs.Data = authen
				rs.Result = true
				var ms models.Message
				ms.Th = "ล็อกอินสำเร็จ"
				ms.En = "login success"
				rs.Message = ms
				json.NewEncoder(response).Encode(rs)
				token.Token = authen.RefreshToken
			}
		}
	}

}

func VerifyAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		authHeader := request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			tokenValid, _ := utils.ValidAccessToken(bearerToken[1])
			if tokenValid.Valid {
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

func VarifyRefresh(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		json.NewDecoder(request.Body).Decode(&token)
		tokenValid, _ := utils.ValidRefreshToken(token.Token)
		if tokenValid.Valid {
			ID, valid := utils.ParseJson(token.Token)
			if valid == false {
				utils.SentMessage(response, false, "parse token failed")
			}
			objID, _ := primitive.ObjectIDFromHex(ID)
			var authen models.Authen
			authen, _ = repository.FindAuthen(objID)
			val, err := repository.GetToken(authen.ID.Hex())
			if val != token.Token {
				utils.SentMessage(response, false, "old token")
			} else {
				if err != nil {
					utils.SentMessage(response, false, "token not found")
				} else {
					next.ServeHTTP(response, request)
				}
			}
		} else {
			utils.SentMessage(response, false, "invalid token")
			return
		}
	})
}

func Logout(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewDecoder(request.Body).Decode(&token)
	tokenValid, _ := utils.ValidRefreshToken(token.Token)
	if tokenValid.Valid {
		ID, valid := utils.ParseJson(token.Token)
		if valid == false {
			utils.SentMessage(response, false, "parse token failed")
		}
		objID, _ := primitive.ObjectIDFromHex(ID)
		var authen models.Authen
		authen, _ = repository.FindAuthen(objID)
		val, err := repository.GetToken(authen.ID.Hex())
		if val != token.Token {
			utils.SentMessage(response, false, "old token")
		} else {
			if err != nil {
				utils.SentMessage(response, false, "token not found")
			} else {
				repository.DeleteToken(authen.ID.Hex())
				utils.SentMessage(response, true, "logout success")
			}
		}
	} else {
		utils.SentMessage(response, true, "invalid token")
		return
	}
}
