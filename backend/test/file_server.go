package main

import (
	_ "flag"
	"log"
	"net/http"
)

func main() {
	// Define command-line flags for directory and port
	dir := "./sample/feeds" //flag.String("dir", ".", "the directory to serve files from")
	port := "5555"          //flag.String("port", "5555", "the port to serve on")
	//flag.Parse()

	// Create a file server handler
	fs := http.FileServer(http.Dir(dir))

	// Start the HTTP server
	log.Printf("Serving %s on HTTP port: %s\n", dir, port)
	if err := http.ListenAndServe(":"+port, fs); err != nil {
		log.Fatal(err)
	}
}
