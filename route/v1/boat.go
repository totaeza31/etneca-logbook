package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatRoute(router *mux.Router) {
	router.HandleFunc("/boat", controllers.VerifyAccess(controllers.GetAllBoat)).Methods("GET")
	router.HandleFunc("/boat/{id}", controllers.GetBoatByID).Methods("GET")
	router.HandleFunc("/boat", controllers.PostBoat).Methods("POST")

	router.HandleFunc("/boat/{id}", controllers.PutBoat).Methods("PUT")
	router.HandleFunc("/boat/{id}", controllers.DelBoat).Methods("DELETE")

	router.HandleFunc("/search", controllers.GetBoatByName).Methods("POST")
}
