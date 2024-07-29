package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	routes "goWebServer/routes"
	server "goWebServer/utility/server"

	"github.com/fsnotify/fsnotify"
)

func main() {
	err := server.LoadEnv()
	if err != nil {
		log.Fatal("Error loading environment variables:", err)
	}

	// Create a log file
	logFile, err := os.OpenFile("http_server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Set up logging to both standard output and the log file
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Initialize routes only once
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

	// Read the WATCHER environment variable from .env
	watcherDir := os.Getenv("WATCHER")
	if watcherDir == "" {
		log.Fatal("Environment variable WATCHER not set")
	}

	// Make sure the directory exists
	if _, err := os.Stat(watcherDir); os.IsNotExist(err) {
		log.Fatalf("Specified directory does not exist: %s", watcherDir)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating file watcher:", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		restartChan := make(chan bool)
		go startServer(restartChan)

		err := filepath.Walk(watcherDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if err := watcher.Add(path); err != nil {
					log.Printf("Error adding watcher to directory %s: %v", path, err)
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal("Error walking the file tree:", err)
		}

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Check if the event is a write and log it
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

	log.Printf("Watching for file changes in the '%s' directory...", watcherDir)

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
