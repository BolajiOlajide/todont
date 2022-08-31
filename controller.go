package main

import "net/http"

// IndexRoute returns a simple text indicating the service is running.
func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to todont"))
	// respond(w, r, "Welcome to todont", http.StatusOK, true)
}
