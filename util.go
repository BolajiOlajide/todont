package main

import (
	"encoding/json"
	"net/http"
)

func respond(w http.ResponseWriter, data any, statusCode int, isText bool) {
	var contentType string
	if isText {
		contentType = "text/plain"
	} else {
		contentType = "application/json"
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", contentType)

	if isText {
		w.Write([]byte("Welcome to todont"))
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
