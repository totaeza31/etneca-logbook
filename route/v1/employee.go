package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func EmployeeRoute(router *mux.Router) {
	router.HandleFunc("/emp", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/emp/{id}", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/emp", controllers.PostEmployee).Methods("POST")
	router.HandleFunc("/emp/{id}", controllers.PutEmployee).Methods("PUT")
	router.HandleFunc("/emp/{id}", controllers.DelEmployee).Methods("DELETE")
}
