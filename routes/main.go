// Main route handler
// define routes and pass to webserver.go

package routes

import (
	"net/http"
)

func HandleRoutes() {
	routes := GqlHandler()

	for _, route := range routes {
		http.HandleFunc(route.Key, route.Value)
	}
}
