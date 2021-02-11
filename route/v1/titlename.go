package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func TitlenameRoute(router *mux.Router) {
	router.HandleFunc("/title", controllers.VerifyAccess(controllers.GetTitles)).Methods("GET")
	router.HandleFunc("/title/{id}", controllers.VerifyAccess(controllers.GetTitle)).Methods("GET")
	router.HandleFunc("/title", controllers.VerifyAccess(controllers.PostTitle)).Methods("POST")
	router.HandleFunc("/title/{id}", controllers.VerifyAccess(controllers.PutTitle)).Methods("PUT")
	router.HandleFunc("/title/{id}", controllers.VerifyAccess(controllers.DelTitle)).Methods("DELETE")
}
