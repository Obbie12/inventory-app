package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

// RespondWithError sends an error response to the client
func RespondWithError(w http.ResponseWriter, statusCode int, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
		log.Printf("Error: %v", err)
	}

	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      errorMsg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// RespondWithJSON sends a JSON response to the client
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}