package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func OwnerRoute(router *mux.Router) {

	router.HandleFunc("/human", controllers.GetAllHuman).Methods("GET")
	router.HandleFunc("/human/{id}", controllers.GetHumanByID).Methods("GET")
	router.HandleFunc("/human", controllers.PostHuman).Methods("POST")
	router.HandleFunc("/human/{id}", controllers.PutHuman).Methods("PUT")
	router.HandleFunc("/human/{id}", controllers.DelHuman).Methods("DELETE")
}
