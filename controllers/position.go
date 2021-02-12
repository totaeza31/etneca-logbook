package controllers

import (
	"encoding/json"

	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
)

func GetPositions(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	position, err := repository.FindAllPosition()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(position.Position)
	}
}
