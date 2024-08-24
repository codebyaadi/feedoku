package internals

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func RespondWithError(w http.ResponseWriter, status int, msg string) {
	if status >= 499 {
		log.Printf("Server error: %s", msg)
	}

	payload := map[string]string{"error": msg}

	RespondWithJSON(w, status, payload)
}
