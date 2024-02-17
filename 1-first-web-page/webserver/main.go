package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	// Set up command line flags
	var (
		port   = flag.Int("port", 8080, "port to listen on")
		wwwdir = flag.String("www", "www", "directory to serve")
	)

	// Parse command line flags
	flag.Parse()

	// Create a file server
	fileServer := http.FileServer(http.Dir(*wwwdir))

	// Mount the file server at the root
	http.Handle("/", fileServer)

	// Start the web server in the background
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Log the web server's address
	log.Printf("web server listening on http://localhost:%d", *port)

	// log the directory being served
	fullpath, err := filepath.Abs(*wwwdir)
	if err != nil {
		log.Printf("error getting absolute path of %s: %v", *wwwdir, err)
	}
	log.Printf("serving files from %s", fullpath)

	// Wait for a signal to stop the web server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
