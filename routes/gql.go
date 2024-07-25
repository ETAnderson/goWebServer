// route for gql interface layer
package routes

import (
	"net/http"
)

type KeyValuePair struct {
	Key   string
	Value func(http.ResponseWriter, *http.Request)
}

func GqlHandler() []KeyValuePair {
	route := "/gql"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("gql placeholder"))
	}

	return []KeyValuePair{
		{Key: route, Value: handler},
	}
}
