// route for create-route interface layer
package routes

import (
	"net/http"
)

func CreateRouteHandler() (string, func(w http.ResponseWriter, r *http.Request)) {
	route := "/create_route"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("create_route placeholder"))
	}

	return route, handler
}

func init() {
	route, handler := CreateRouteHandler()
	RegisterRoute(route, handler)
}
