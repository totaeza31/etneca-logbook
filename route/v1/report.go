package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func ReportRoute(router *mux.Router) {
	router.HandleFunc("/report", controllers.GetAllReport).Methods("GET")
	router.HandleFunc("/report/{id}", controllers.GetReportByID).Methods("GET")
	router.HandleFunc("/report", controllers.PostReport).Methods("POST")
	// router.HandleFunc("/report/{id}", controllers.PutReport).Methods("PUT")
	// router.HandleFunc("/report/{id}", controllers.DelReport).Methods("DELETE")
}
