package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func PositionRoute(router *mux.Router) {
	router.HandleFunc("/pst", controllers.VerifyAccess(controllers.GetPositions)).Methods("GET")
	router.HandleFunc("/pst/{id}", controllers.VerifyAccess(controllers.GetPosition)).Methods("GET")
	router.HandleFunc("/pst", controllers.VerifyAccess(controllers.PostPosition)).Methods("POST")
	router.HandleFunc("/pst/{id}", controllers.VerifyAccess(controllers.PutPosition)).Methods("PUT")
	router.HandleFunc("/pst/{id}", controllers.VerifyAccess(controllers.DelPosition)).Methods("DELETE")
}
