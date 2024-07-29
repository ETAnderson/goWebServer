package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run create_route.go <filename>")
		os.Exit(1)
	}

	// Get the filename from the arguments and ensure it has a .go extension
	filename := os.Args[1]
	if filepath.Ext(filename) != ".go" {
		filename += ".go"
	}

	// Create the route file
	routeFilePath := filepath.Join("routes", filename)
	routeFile, err := os.Create(routeFilePath)
	if err != nil {
		log.Fatalf("Failed to create route file: %v", err)
	}
	defer routeFile.Close()

	// Create the content for the route file
	routeContent := fmt.Sprintf(`// route for %s interface layer
package routes

import (
	"net/http"
)

func %sHandler() (string, func(w http.ResponseWriter, r *http.Request)) {
	route := "/%s"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("%s placeholder"))
	}

	return route, handler
}

func init() {
	route, handler := %sHandler()
	RegisterRoute(route, handler)
}
`, filename[:len(filename)-3], filename[:len(filename)-3], filename[:len(filename)-3], filename[:len(filename)-3], filename[:len(filename)-3])

	_, err = routeFile.WriteString(routeContent)
	if err != nil {
		log.Fatalf("Failed to write to route file: %v", err)
	}

	fmt.Printf("Route file '%s' created successfully.\n", routeFilePath)

	// Create the HTML template file
	templateFilePath := filepath.Join("templates", filename[:len(filename)-3]+".html")
	templateFile, err := os.Create(templateFilePath)
	if err != nil {
		log.Fatalf("Failed to create template file: %v", err)
	}
	defer templateFile.Close()

	// Create the content for the HTML template file
	templateContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s Page</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
</head>
<body>
    <header>
        <h1>%s Page</h1>
        <nav>
            <a href="/">Home</a>
            <!-- Add more navigation links as needed -->
        </nav>
    </header>
    <main>
        <h2>Welcome to the %s Page</h2>
        <p>This is the template for the %s route.</p>
    </main>
    <footer>
        <p>&copy; 2024 My Web Server</p>
    </footer>
</body>
</html>
`, filename[:len(filename)-3], filename[:len(filename)-3], filename[:len(filename)-3], filename[:len(filename)-3])

	_, err = templateFile.WriteString(templateContent)
	if err != nil {
		log.Fatalf("Failed to write to template file: %v", err)
	}

	fmt.Printf("Template file '%s' created successfully.\n", templateFilePath)
}
