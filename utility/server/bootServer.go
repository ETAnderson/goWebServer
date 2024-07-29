package server

import (
	"fmt"
	"net/http"
)

func Serve(port string) error {
	fmt.Println("Starting server on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	return nil
}
