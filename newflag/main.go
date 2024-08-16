package main

import (
	"log"
	"net/http"
)

func main() {
	// Set the directory to serve static files
	fs := http.FileServer(http.Dir("."))

	// Handle all routes with the FileServer
	http.Handle("/", fs)

	// Define the port to listen on
	port := ":8080"
	log.Printf("Starting server on http://localhost%s\n", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
