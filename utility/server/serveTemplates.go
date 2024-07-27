package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const templatesDir = "templates"

func LoadTemplate(route string) (*template.Template, error) {
	// Normalize the route
	if route == "/" {
		route = "index.html"
	} else {
		route = strings.TrimPrefix(route, "/")
	}

	// Create the full path for the requested template
	templatePath := filepath.Join(templatesDir, route)
	headerPath := filepath.Join(templatesDir, "header.html")

	// Log the template path being checked
	fmt.Printf("Checking template paths: %s and %s\n", templatePath, headerPath)

	// Check if the requested template file exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		// Log and return an error if the file does not exist
		return nil, fmt.Errorf("file does not exist: %s", templatePath)
	}

	// Parse both the header and the requested template files
	tmpl, err := template.New("base").ParseFiles(headerPath, templatePath)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func ServeMatchingRouteTemplate(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Load the template for the route
		tmpl, err := LoadTemplate(route)
		if err != nil {
			// Log the error if the template is not found
			fmt.Printf("No template found for %s\n", route)
			http.NotFound(w, r)
			return
		}

		// Log the template being executed
		fmt.Printf("Rendering template for route: %s\n", route)

		// Render the template
		err = tmpl.ExecuteTemplate(w, filepath.Base(tmpl.Name()), nil)
		if err != nil {
			// Handle rendering errors gracefully
			fmt.Printf("Error rendering template: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
