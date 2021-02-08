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
	if err != nil || authen.Email == "" || authen.Password == "" {
		respond = models.Invalid_syntax()
		utils.SentMessage(response, respond)
	} else {
		var password = authen.Password
		authen, err = repository.FindEmail(authen.Email)
		if err != nil {
			respond = models.Email_invalid()
			utils.SentMessage(response, respond)
		} else {
			err = utils.Decrypt(password, authen.Password)
			if err != nil {
				respond = models.Email_invalid()
				utils.SentMessage(response, respond)
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
		response.Header().Set("Access-Control-Allow-Origin", "*")
		authHeader := request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			tokenValid, _ := utils.ValidAccessToken(bearerToken[1])
			if tokenValid.Valid {
				next.ServeHTTP(response, request)
			} else {
				respond = models.Token_expired()
				utils.SentMessage(response, respond)
				return
			}
		} else {
			respond = models.Invalid_token()
			utils.SentMessage(response, respond)
			return
		}
	})
}

func VarifyRefresh(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewDecoder(request.Body).Decode(&token)
		tokenValid, _ := utils.ValidRefreshToken(token.Token)
		if tokenValid.Valid {
			ID, valid := utils.ParseJson(token.Token)
			if valid == false {
				respond = models.User_not_found()
				utils.SentMessage(response, respond)
			} else {
				objID, _ := primitive.ObjectIDFromHex(ID)
				var authen models.Authen
				authen, _ = repository.FindAuthen(objID)
				val, err := repository.GetToken(authen.ID.Hex())
				if val != token.Token {
					respond = models.Token_expired()
					utils.SentMessage(response, respond)
				} else {
					if err != nil {
						respond = models.Invalid_token()
						utils.SentMessage(response, respond)
					} else {
						next.ServeHTTP(response, request)
					}
				}
			}
		} else {
			respond = models.Invalid_token()
			utils.SentMessage(response, respond)
			return
		}
	})
}

func Logout(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(request.Body).Decode(&token)
	tokenValid, _ := utils.ValidRefreshToken(token.Token)
	if tokenValid.Valid {
		ID, valid := utils.ParseJson(token.Token)
		if valid == false {
			respond = models.User_not_found()
			utils.SentMessage(response, respond)
		} else {
			objID, _ := primitive.ObjectIDFromHex(ID)
			var authen models.Authen
			authen, _ = repository.FindAuthen(objID)
			val, err := repository.GetToken(authen.ID.Hex())
			if val != token.Token {
				respond = models.User_not_found()
				utils.SentMessage(response, respond)
			} else {
				if err != nil {
					respond = models.Token_expired()
					utils.SentMessage(response, respond)
				} else {
					repository.DeleteToken(authen.ID.Hex())
					respond = models.Logout_success()
					utils.SentMessage(response, respond)
				}
			}
		}
	} else {
		respond = models.Invalid_token()
		utils.SentMessage(response, respond)
		return
	}
}
