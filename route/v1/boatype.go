package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatTypeRoute(router *mux.Router) {
	router.HandleFunc("/boatType", controllers.GetAllBoatType).Methods("GET")
	router.HandleFunc("/boatType/{id}", controllers.GetBoatTypeByID).Methods("GET")
	router.HandleFunc("/boatType", controllers.PostBoatType).Methods("POST")
	router.HandleFunc("/boatType/{id}", controllers.PutBoatType).Methods("PUT")
	router.HandleFunc("/boatType/{id}", controllers.DelBoatType).Methods("DELETE")
}
