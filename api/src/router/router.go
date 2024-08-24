package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

// return an Router with the routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.Configure(r)
}
