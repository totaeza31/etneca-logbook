package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatFinanceRoute(router *mux.Router) {
	router.HandleFunc("/boatFinance", controllers.GetAllBoatFinance).Methods("GET")
	router.HandleFunc("/boatFinance/{id}", controllers.GetBoatFinanceByID).Methods("GET")
	router.HandleFunc("/boatFinance", controllers.PostBoatFinance).Methods("POST")
	router.HandleFunc("/boatFinance/{id}", controllers.PutBoatFinance).Methods("PUT")
	router.HandleFunc("/boatFinance/{id}", controllers.DelBoatFinance).Methods("DELETE")
}
