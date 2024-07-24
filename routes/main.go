// Main route handler
// define routes and pass to webserver.go

package routes

import (
	"fmt"
	"net/http"
)

func HandleRoutes() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Web Server!")
	}

	http.HandleFunc("/", handler)
}
