// Main route handler
// define routes and pass to webserver.go

package routes

import (
	"fmt"
	"net/http"
)

func HandleRoutes() {
	routes := GetRoutes()

	for route, handler := range routes {
		fmt.Println("Setting up route:", route)
		http.HandleFunc(route, handler)
	}
}
