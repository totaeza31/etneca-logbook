package route

import (
	test "etneca-logbook/route/testRedux"
	v1 "etneca-logbook/route/v1"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	v1.BoatTypeRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatGatewayRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatBeamRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatDeviceRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatFinanceRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatVgmRoute(router.PathPrefix("/v1").Subrouter())
	v1.ReportRoute(router.PathPrefix("/v1").Subrouter())
	v1.GoodsRoute(router.PathPrefix("/v1").Subrouter())
	v1.WorksheetRoute(router.PathPrefix("/v1").Subrouter())
	v1.BoatRemarkRoute(router.PathPrefix("/v1").Subrouter())
	v1.EmployeeRoute(router.PathPrefix("/v1").Subrouter())
	v1.CompanyRoute(router.PathPrefix("/v1").Subrouter())
	v1.TitlenameRoute(router.PathPrefix("/v1").Subrouter())
	v1.PositionRoute(router.PathPrefix("/v1").Subrouter())
	v1.GenderRoute(router.PathPrefix("/v1").Subrouter())

	test.AuthenRoute(router.PathPrefix("/test").Subrouter())
	test.OwnerRoute(router.PathPrefix("/test").Subrouter())

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(router))
}
