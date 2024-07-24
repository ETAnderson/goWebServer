package main


import (
	"goWebServer/routes"
	"goWebServer/utility/server"
)

func main() {
	routes.HandleRoutes()

	boot.Serve()
}
