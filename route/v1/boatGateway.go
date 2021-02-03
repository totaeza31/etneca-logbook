package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatGatewayRoute(router *mux.Router) {
	router.HandleFunc("/boatGateway", controllers.GetAllBoatGateway).Methods("GET")
	router.HandleFunc("/boatGateway/{id}", controllers.GetBoatGatewayByID).Methods("GET")
	router.HandleFunc("/boatGateway", controllers.PostBoatGateway).Methods("POST")
	router.HandleFunc("/boatGateway/{id}", controllers.PutBoatGateway).Methods("PUT")
	router.HandleFunc("/boatGateway/{id}", controllers.DelBoatGateway).Methods("DELETE")
}
