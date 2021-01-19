package utils

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Alert struct {
	Result  bool   "false"
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

func SentMessage(response http.ResponseWriter, result bool, message string) {
	var alert Alert
	alert.Result = result
	alert.Message = message
	response.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(response).Encode(alert)
}

func Decrypt(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func Encrypt(text string) string {

	bytes, _ := bcrypt.GenerateFromPassword([]byte(text), 14)

	fmt.Println(bcrypt.Cost(bytes))
	return string(bytes)

}

func GenerateToken(authen models.Authen, types string) (string, error) {
	var secret string
	var expires int64

	if types == "access" {
		secret = os.Getenv("ACCESS_TOKEN")
		expires = time.Now().Add(time.Second * 30).Unix()
	} else if types == "refresh" {
		secret = os.Getenv("REFRESH_TOKEN")
		expires = time.Now().Add(time.Minute * 3).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       authen.ID,
		"email":    authen.Email,
		"password": authen.Password,
		"exp":      expires,
	})

	tokenString, _ := token.SignedString([]byte(secret))

	if types == "refresh" {

		tk := time.Unix(expires, 0)
		sub := time.Now()
		stringID := authen.ID.Hex()
		err := repository.SetToken(stringID, tokenString, tk, sub)
		if err != nil {
			return "", err
		}
	}
	return tokenString, nil
}