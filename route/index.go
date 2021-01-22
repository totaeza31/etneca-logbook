package route

import (
	v1 "etneca-logbook/route/v1"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func IndexRoute() {
	router := mux.NewRouter()

	v1.AuthenRoute(router.PathPrefix("/v1").Subrouter())
	v1.OwnerRoute(router.PathPrefix("/v1").Subrouter())

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
