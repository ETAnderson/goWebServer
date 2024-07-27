package main

import (
	routes "goWebServer/routes"
	server "goWebServer/utility/server"
)

func main() {
	routes.HandleRoutes()
	server.Serve()
}
