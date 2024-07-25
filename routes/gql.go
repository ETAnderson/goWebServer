// route for gql interface layer
package gql

import (
	"fmt"
	"net/http"
)

func GqlHandler() (string, func(w http.ResponseWriter, r *http.Request)) {
	route := "/gql"

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gql placeholder")
	}

	return route, handler
}
