package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func PositionRoute(router *mux.Router) {
	router.HandleFunc("/pst", controllers.GetPositions).Methods("GET")
	// router.HandleFunc("/pst/{id}", controllers.GetPosition).Methods("GET")
	// router.HandleFunc("/pst", controllers.PostPosition).Methods("POST")
	// router.HandleFunc("/pst/{id}", controllers.PutPosition).Methods("PUT")
	// router.HandleFunc("/pst/{id}", controllers.DelPosition).Methods("DELETE")
}
