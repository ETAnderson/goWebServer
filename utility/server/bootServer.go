package server

import (
	"fmt"
	"net/http"
	"os"
)

func Serve() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run webserver.go <port>")
		return
	}
	port := os.Args[1]
	fmt.Println("Starting server on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
