package main

import (
	"etneca-logbook/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	router.HandleFunc("/forgot", controllers.GetNewPassword).Methods("POST")
	router.HandleFunc("/reset/{email}", controllers.ResetPassword).Methods("GET")

	router.HandleFunc("/deleteUser/{id}", controllers.DeleteUser).Methods("GET")
	router.HandleFunc("/updateUser/{id}", controllers.UpdateUser).Methods("GET")

	router.HandleFunc("/package", controllers.GetPackage).Methods("GET")

	// router.HandleFunc("/owner", controllers.GetPackage).Methods("POST")
	router.HandleFunc("/owner", controllers.GetOwner).Methods("GET")
	router.HandleFunc("/owner/{id}", controllers.GetOwnerByID).Methods("GET")
	router.HandleFunc("/owner", controllers.PostOwner).Methods("POST")

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
