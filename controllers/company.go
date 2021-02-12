package controllers

import (
	"encoding/json"

	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"
)

func GetAllCompany(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	company, err := repository.FindAllCompany()

	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(company.Company)
	}
}
