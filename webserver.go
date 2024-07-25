package main

import (
	boot "goWebServer/utility/server"
)

func main() {
	routes.HandleRoutes()
	boot.Serve()
}
