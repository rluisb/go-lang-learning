package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all routes structure from the API
type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
	AuthorizationRequired bool
}

//Configure router by adding all routes inside mux.Router
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}