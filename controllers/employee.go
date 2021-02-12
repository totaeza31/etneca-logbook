package controllers

import (
	"encoding/json"

	"etneca-logbook/models"
	"etneca-logbook/repository"
	"etneca-logbook/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEmployees(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	employee, err := repository.FindAllEmployee()
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(employee.GetEmployee)
	}
}

func GetEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	emp, err := repository.FindEmployee(id)
	if err != nil {
		message := models.Get_data_error()
		utils.SentMessage(response, message)
	} else {
		json.NewEncoder(response).Encode(emp)
	}
}

func PostEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var emp models.Employee
	err := json.NewDecoder(request.Body).Decode(&emp)
	if err != nil {
		message := models.Invalid_syntax()
		utils.SentMessage(response, message)
	} else {
		emp.Password = utils.Encrypt(emp.Password)

		emp.BirthdayTime = utils.TimeFormat(emp.Birthday)
		emp.StartDateTime = utils.TimeFormat(emp.StartDate)
		emp.EndDateTime = utils.TimeFormat(emp.EndDate)
		emp.EnsureDateTime = utils.TimeFormat(emp.EnsureDate)
		id, _ := utils.GenerateEmpID(emp)
		emp.ID = id
		err = repository.InsertEmployee(emp)
		if err != nil {
			message := models.Update_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Update_success()
			utils.SentMessage(response, message)
		}
	}
}

func PutEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	_, err := repository.FindEmployee(id)
	if err != nil {
		message := models.User_not_found()
		utils.SentMessage(response, message)
	} else {
		var emp models.Employee
		json.NewDecoder(request.Body).Decode(&emp)
		emp.Password = utils.Encrypt(emp.Password)

		err = repository.UpdateEmployee(emp, id)
		if err != nil {
			message := models.Edit_error()
			utils.SentMessage(response, message)
		} else {
			message := models.Edit_success()
			utils.SentMessage(response, message)
		}
	}
}

func DelEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	id := param["id"]
	err := repository.DeleteEmployee(id)
	if err != nil {
		message := models.Delete_error()
		utils.SentMessage(response, message)
	} else {
		message := models.Delete_success()
		utils.SentMessage(response, message)
	}
}
