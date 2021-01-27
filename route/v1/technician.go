package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func TechnicianRoute(router *mux.Router) {
	router.HandleFunc("/tech", controllers.VerifyAccess(controllers.GetTech)).Methods("GET")
	router.HandleFunc("/tech/{id}", controllers.GetTechByID).Methods("GET")
	router.HandleFunc("/tech", controllers.PostTech).Methods("POST")
	router.HandleFunc("/tech/{id}", controllers.PutTech).Methods("PUT")
	router.HandleFunc("/tech/{id}", controllers.DelTech).Methods("DELETE")
}