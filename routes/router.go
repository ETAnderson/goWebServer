package routes

import (
	"net/http"
	"sync"
)

var (
	routes = make(map[string]func(http.ResponseWriter, *http.Request))
	mu     sync.Mutex
)

func RegisterRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	mu.Lock()
	defer mu.Unlock()
	routes[route] = handler
}

func GetRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return routes
}
