package routers

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/gorilla/mux"
)

// Route is a struct that represents a route of the API
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// Configure configures the routes of the API
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	for _, route := range routes {
		// Print the Function name
		println(route.URI + " " + route.Method + " -> " + runtime.FuncForPC(reflect.ValueOf(route.Function).Pointer()).Name())
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
