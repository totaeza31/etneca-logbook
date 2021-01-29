package route

import (
	test "etneca-logbook/route/testRedux"
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
	v1.TechnicianRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatRoute(router.PathPrefix("/v1").Subrouter())

	test.AuthenRoute(router.PathPrefix("/test").Subrouter())
	test.OwnerRoute(router.PathPrefix("/test").Subrouter())

	cors := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowCredentials: true,
	})
	handler := cors.Handler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
