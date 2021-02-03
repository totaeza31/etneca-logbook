package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatDeviceRoute(router *mux.Router) {
	router.HandleFunc("/boatDevice", controllers.GetAllBoatDevice).Methods("GET")
	router.HandleFunc("/boatDevice/{id}", controllers.GetBoatDeviceByID).Methods("GET")
	router.HandleFunc("/boatDevice", controllers.PostBoatDevice).Methods("POST")
	router.HandleFunc("/boatDevice/{id}", controllers.PutBoatDevice).Methods("PUT")
	router.HandleFunc("/boatDevice/{id}", controllers.DelBoatDevice).Methods("DELETE")
}
