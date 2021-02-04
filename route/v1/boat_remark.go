package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func BoatRemarkRoute(router *mux.Router) {
	// router.HandleFunc("/boat/{id}/remark", controllers.GetAllBoatType).Methods("GET")
	// router.HandleFunc("/boat/{id}/remark/{r_id}", controllers.GetBoatTypeByID).Methods("GET")
	router.HandleFunc("/boat/{id}/remark", controllers.PostBoatRemark).Methods("POST")
	// router.HandleFunc("/boatType/{id}", controllers.PutBoatType).Methods("PUT")
	router.HandleFunc("/boatType/{id}/remark", controllers.DelBoatRemark).Methods("DELETE")
}
