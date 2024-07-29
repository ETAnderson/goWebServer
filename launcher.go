package main

import (
	routes "goWebServer/routes"
	server "goWebServer/utility/server"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	routes.HandleRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.Serve()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
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
					log.Println("modified file:", event.Name)
					restartChan <- true
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Watch current directory and its subdirectories
	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			return watcher.Add(path)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	<-done
}

func startServer(restartChan chan bool) {
	cmd := exec.Command("go", "run", "main.go")

	for {
		log.Println("Starting server...")
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		go func() {
			if err := cmd.Wait(); err != nil {
				log.Println("Server exited with error:", err)
			}
		}()

		<-restartChan
		log.Println("Restarting server...")
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("Failed to kill process:", err)
		}

		cmd = exec.Command("go", "run", "main.go")
		time.Sleep(1 * time.Second) // Give some time for the process to release resources
	}
}
