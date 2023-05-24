package api

import (
	"log"
	"net/http"
)

// healthCheck is a handler function that returns a 200 status code
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Health check!"))
	log.Printf("Health check")
}

// return index hellow world page
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// write message back to user passed in via URL
	w.Write([]byte("Hello World!"))
	log.Printf("Hello World")
}
