package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatRoute(router *mux.Router) {
	router.HandleFunc("/boat", controllers.VerifyAccess(controllers.GetTech)).Methods("GET")
	router.HandleFunc("/boat/{id}", controllers.GetTechByID).Methods("GET")
	router.HandleFunc("/boat", controllers.PostTech).Methods("POST")
	router.HandleFunc("/boat/{id}", controllers.PutTech).Methods("PUT")
	router.HandleFunc("/boat/{id}", controllers.DelTech).Methods("DELETE")
}