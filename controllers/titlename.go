package controllers

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
)

func GetTitles(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	allTitle, err := repository.FindAllTitle()
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {

		json.NewEncoder(response).Encode(allTitle.Title)
	}
}
