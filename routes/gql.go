// route for gql interface layer
package routes

import (
	"net/http"
)

func GqlHandler() (string, func(w http.ResponseWriter, r *http.Request)) {
	route := "/gql"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("gql placeholder"))
	}

	return route, handler
}

func init() {
	route, handler := GqlHandler()
	RegisterRoute(route, handler)
}
