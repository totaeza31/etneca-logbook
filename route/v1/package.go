package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func PackageRoute(router *mux.Router) {
	router.HandleFunc("/package", controllers.GetPackage).Methods("GET")
}
