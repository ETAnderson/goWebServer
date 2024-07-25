package main

import (
	routes "goWebServer/routes"
	boot "goWebServer/utility/server"
)

func main() {
	routes.HandleRoutes()
	boot.Serve()
}
