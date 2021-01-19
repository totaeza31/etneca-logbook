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
	router.HandleFunc("/profile", controllers.Profile).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
