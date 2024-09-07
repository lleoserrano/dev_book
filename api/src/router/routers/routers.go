package routers

import (
	"api/src/middlewares"
	"net/http"

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
	routes = append(routes, loginRoute)

	for _, route := range routes {

		if route.NeedAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Function),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)

		}

	}
	return r
}
