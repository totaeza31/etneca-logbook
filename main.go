package main

import (
	"etneca-logbook/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/profile", controllers.VerifyAccess(controllers.GetProfile)).Methods("GET")
	router.HandleFunc("/token", controllers.VarifyRefresh(controllers.GetNewToken)).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
