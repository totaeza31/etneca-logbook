package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func GoodsRoute(router *mux.Router) {
	router.HandleFunc("/goods", controllers.GetAllGoods).Methods("GET")
	router.HandleFunc("/goods/{id}", controllers.GetGoodsByID).Methods("GET")
	router.HandleFunc("/goods", controllers.PostGoods).Methods("POST")
	router.HandleFunc("/goods/{id}", controllers.PutGoods).Methods("PUT")
	router.HandleFunc("/goods/{id}", controllers.DelGoods).Methods("DELETE")
}
