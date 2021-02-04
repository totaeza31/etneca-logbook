package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func TechnicianRoute(router *mux.Router) {
	router.HandleFunc("/tech", controllers.VerifyAccess(controllers.GetTech)).Methods("GET")
	router.HandleFunc("/tech/{id}", controllers.VerifyAccess(controllers.GetTechByID)).Methods("GET")
	router.HandleFunc("/tech", controllers.VerifyAccess(controllers.PostTech)).Methods("POST")
	router.HandleFunc("/tech/{id}", controllers.VerifyAccess(controllers.PutTech)).Methods("PUT")
	router.HandleFunc("/tech/{id}", controllers.VerifyAccess(controllers.DelTech)).Methods("DELETE")
}
