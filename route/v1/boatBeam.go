package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatBeamRoute(router *mux.Router) {
	router.HandleFunc("/boatBeam", controllers.GetAllBoatBeam).Methods("GET")
	router.HandleFunc("/boatBeam/{id}", controllers.GetBoatBeamByID).Methods("GET")
	router.HandleFunc("/boatBeam", controllers.PostBoatBeam).Methods("POST")
	router.HandleFunc("/boatBeam/{id}", controllers.PutBoatBeam).Methods("PUT")
	router.HandleFunc("/boatBeam/{id}", controllers.DelBoatBeam).Methods("DELETE")
}
