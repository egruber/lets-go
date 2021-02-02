package main

import (
	"log"
	"net/http"
)

// Create the home function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippetbox. Golang course"))
}

func main() {
	// Start a new server mux
	mux := http.NewServeMux()

	// This handles anything to the / location
	// And calls the home function to address it
	mux.HandleFunc("/", home)

	// Because logging matters
	log.Println("Starting server on port 4000")

	//
	err := http.ListenAndServe(":4000", mux)
	log.Fatalf("Error: %s", err)
}
