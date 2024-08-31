package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON sends a JSON response with the specified status code and payload.
//
// Parameters:
//   - w: The http.ResponseWriter to send the response to.
//   - status: The HTTP status code to set in the response.
//   - payload: The data to be marshaled into JSON and sent in the response body.
//
// The function attempts to marshal the provided payload into a JSON format. If successful,
// it sets the "Content-Type" header to "application/json", writes the status code, and sends
// the JSON response. If JSON marshalling fails, an error is logged, and an HTTP 500 (Internal Server Error)
// response is sent back to the client.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal json response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondWithError sends a JSON response with an error message and the specified status code.
//
// Parameters:
//   - w: The http.ResponseWriter to send the response to.
//   - status: The HTTP status code to set in the response. If the status code is 500 or higher,
//     the error message will be logged as a server error.
//   - msg: The error message to include in the response body.
//
// The function creates a JSON payload containing the error message and sends it using the
// RespondWithJSON function. If the status code indicates a server error (status >= 499),
// the error message is logged.
func RespondWithError(w http.ResponseWriter, status int, msg string) {
	if status >= 499 {
		log.Printf("Server error: %s", msg)
	}

	payload := map[string]string{"error": msg}

	RespondWithJSON(w, status, payload)
}
