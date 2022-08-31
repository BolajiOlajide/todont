package main

import "net/http"

// IndexRoute returns a simple text indicating the service is running.
func indexRoute(w http.ResponseWriter, r *http.Request) {
	respond(w, "Welcome to todont", http.StatusOK, true)
}
