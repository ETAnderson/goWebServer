package routes

import (
	"net/http"
)

// MyNewRouteHandler returns the route and its handler function.
func MyNewRouteHandler() (string, func(w http.ResponseWriter, r *http.Request)) {
	route := "/my_new_route"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("my_new_route placeholder"))
	}

	return route, handler
}

// init registers the MyNewRouteHandler.
func init() {
	route, handler := MyNewRouteHandler()
	RegisterRoute(route, handler)
}
