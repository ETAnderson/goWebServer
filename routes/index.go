// route for Indexing
package routes

import (
	"goWebServer/utility/server"
	"log"
	"net/http"
)

// IndexHandler function to generate the index page with route links
func IndexHandler() (string, http.HandlerFunc) {
	route := "/"
	indexRoutes := GetRoutes()

	handler := func(w http.ResponseWriter, r *http.Request) {
		var routes []string
		for route := range indexRoutes {
			routes = append(routes, route)
		}

		data := struct {
			Title       string
			Heading     string
			Description string
			Routes      []string
		}{
			Title:       "goWebServer Index",
			Heading:     "Welcome to goWebServer",
			Description: "A dynamically-generated route map",
			Routes:      routes,
		}

		// Load the template
		tmpl, err := server.LoadTemplate("index.html")
		if err != nil {
			log.Printf("Error loading template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Render the template
		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			// Only set the error response if no other response has been sent
			if w.Header().Get("Content-Type") != "" {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	}
	return route, handler
}

func init() {
	route, handler := IndexHandler()
	RegisterRoute(route, handler)
}
