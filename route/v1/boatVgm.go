package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatVgmRoute(router *mux.Router) {
	router.HandleFunc("/boatVgm", controllers.GetAllBoatVgm).Methods("GET")
	router.HandleFunc("/boatVgm/{id}", controllers.GetBoatVgmByID).Methods("GET")
	router.HandleFunc("/boatVgm", controllers.PostBoatVgm).Methods("POST")
	router.HandleFunc("/boatVgm/{id}", controllers.PutBoatVgm).Methods("PUT")
	router.HandleFunc("/boatVgm/{id}", controllers.DelBoatVgm).Methods("DELETE")
}
