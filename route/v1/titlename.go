package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func TitlenameRoute(router *mux.Router) {
	router.HandleFunc("/title", controllers.GetTitles).Methods("GET")
	// router.HandleFunc("/title/{id}", controllers.GetTitle).Methods("GET")
	// router.HandleFunc("/title", controllers.PostTitle).Methods("POST")
	// router.HandleFunc("/title/{id}", controllers.PutTitle).Methods("PUT")
	// router.HandleFunc("/title/{id}", controllers.DelTitle).Methods("DELETE")
}
