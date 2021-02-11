package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func CompanyRoute(router *mux.Router) {
	router.HandleFunc("/comp", controllers.VerifyAccess(controllers.GetAllCompany)).Methods("GET")
	router.HandleFunc("/comp/{id}", controllers.VerifyAccess(controllers.GetCompany)).Methods("GET")
	router.HandleFunc("/comp", controllers.VerifyAccess(controllers.PostCompany)).Methods("POST")
	router.HandleFunc("/comp/{id}", controllers.VerifyAccess(controllers.PutCompany)).Methods("PUT")
	router.HandleFunc("/comp/{id}", controllers.VerifyAccess(controllers.DelCompany)).Methods("DELETE")
}
