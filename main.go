package main

import (
	"log" // simple logging of errors and info
	"my-first-server/router"
	"net/http" // HTTP server and client implementations
	// third‑party router for advanced URL matching
)

// Main entry point
func main() {
	r := router.SetupRouter()
	//Run Server //Fatal(from log package): if fails
	log.Fatal(http.ListenAndServe(":8000", r))
}


