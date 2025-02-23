package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	// code in the 400 range are client-side errors, so we dont need to know about them
	// code in the 500 range are server-side errors, so we should log them as it indicates a bug on our end
	if code > 499 {
		log.Println("Responding with 5XX error: ", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
		// telling json.Marshal() that JSON key will be "error" for this field, these are called JSON reflect tags
	}

	respondWithJson(w, code, errorResponse{
		Error: message,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data) // Write the JSON to the response body
}
