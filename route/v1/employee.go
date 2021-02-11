package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func EmployeeRoute(router *mux.Router) {
	router.HandleFunc("/emp", controllers.VerifyAccess(controllers.GetEmployees)).Methods("GET")
	router.HandleFunc("/emp/{id}", controllers.VerifyAccess(controllers.GetEmployee)).Methods("GET")
	router.HandleFunc("/emp", controllers.VerifyAccess(controllers.PostEmployee)).Methods("POST")
	router.HandleFunc("/emp/{id}", controllers.VerifyAccess(controllers.PutEmployee)).Methods("PUT")
	router.HandleFunc("/emp/{id}", controllers.VerifyAccess(controllers.DelEmployee)).Methods("DELETE")
}
