package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

const staticFilesDir = "public/static/"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Info Actuator",
		"GET",
		"/actuator/info",
		InfoHandler,
	},
	Route{
		"Health Actuator",
		"GET",
		"/actuator/health",
		HealthHandler,
	},
}

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFilesDir)))
	// router.NotFoundHandler = http.HandlerFunc(NotFound)

	return router
}
