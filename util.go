package main

import (
	"encoding/json"
	"net/http"
)

type status string

const (
	successStatus status = "success"
	errorStatus   status = "error"
)

func respond(w http.ResponseWriter, r *http.Request, data any, statusCode int, ok bool) {
	resp := response{}

	if ok {
		resp.Status = successStatus
		resp.Data = data
		resp.Message = ""
	} else {
		resp.Status = errorStatus
		str, isString := data.(string)
		if !isString {
			resp.Message = "An error occurred while processing your request."
		} else {
			resp.Message = str
		}
		resp.Data = nil
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
