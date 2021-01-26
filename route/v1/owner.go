package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func OwnerRoute(router *mux.Router) {

	router.HandleFunc("/forgot", controllers.GetNewPassword).Methods("POST")
	router.HandleFunc("/reset/{email}", controllers.ResetPassword).Methods("GET")

	router.HandleFunc("/package", controllers.GetPackage).Methods("GET")

	router.HandleFunc("/owner", controllers.VerifyAccess(controllers.GetOwner)).Methods("GET")
	router.HandleFunc("/owner/{id}", controllers.GetOwnerByID).Methods("GET")
	router.HandleFunc("/owner", controllers.PostOwner).Methods("POST")
	router.HandleFunc("/owner/{id}", controllers.PutOwner).Methods("PUT")
	router.HandleFunc("/owner/{id}", controllers.DelOwner).Methods("DELETE")
	router.HandleFunc("/owner/credit/{id}", controllers.PatchOwnerCredit).Methods("PATCH")
}
