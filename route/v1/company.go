package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func CompanyRoute(router *mux.Router) {
	router.HandleFunc("/comp", controllers.GetAllCompany).Methods("GET")
	// router.HandleFunc("/comp/{id}", controllers.GetCompany).Methods("GET")
	// router.HandleFunc("/comp", controllers.PostCompany).Methods("POST")
	// router.HandleFunc("/comp/{id}", controllers.PutCompany).Methods("PUT")
	// router.HandleFunc("/comp/{id}", controllers.DelCompany).Methods("DELETE")
}
