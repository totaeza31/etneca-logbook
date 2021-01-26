package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func AuthenRoute(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/profile", controllers.VerifyAccess(controllers.GetProfile)).Methods("GET")
	router.HandleFunc("/token", controllers.VarifyRefresh(controllers.GetNewToken)).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	router.HandleFunc("/forgot", controllers.GetNewPassword).Methods("POST")
	router.HandleFunc("/reset/{email}", controllers.ResetPassword).Methods("GET")

}
