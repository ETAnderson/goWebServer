package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	routes "goWebServer/routes"
	server "goWebServer/utility/server"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create a log file
	logFile, err := os.OpenFile("http_server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Set up logging to both standard output and the log file
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Initializing routes...")
	routes.HandleRoutes()

	log.Println("Setting up static file server...")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Starting server...")
	go func() {
		if len(os.Args) < 2 {
			log.Fatal("Usage: go run webserver.go <port>")
		}
		port := os.Args[1]

		err := server.Serve(port)
		if err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating file watcher:", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		restartChan := make(chan bool)
		go startServer(restartChan)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file:", event.Name)
					restartChan <- true
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	log.Println("Watching for file changes in the 'templates' directory...")
	err = watcher.Add("templates")
	if err != nil {
		log.Fatal("Error adding watcher to templates directory:", err)
	}

	<-done
}

func startServer(restartChan chan bool) {
	for {
		<-restartChan
		log.Println("Restarting server...")

		cmd := exec.Command("go", "build", "-o", "server", ".")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to build server: %v", err)
		}

		cmd = exec.Command("./server")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Start()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Printf("Server exited with error: %v", err)
		}
	}
}
