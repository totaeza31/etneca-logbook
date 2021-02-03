package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func WorksheetRoute(router *mux.Router) {
	router.HandleFunc("/worksheet", controllers.GetAllWorksheet).Methods("GET")
	router.HandleFunc("/worksheet/{id}", controllers.GetWorksheetByID).Methods("GET")
	router.HandleFunc("/worksheet", controllers.PostWorksheet).Methods("POST")
	router.HandleFunc("/worksheet/{id}", controllers.PutWorksheet).Methods("PUT")
	router.HandleFunc("/worksheet/{id}", controllers.DelWorksheet).Methods("DELETE")
}
