package utils

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey retrieves the API key from the provided HTTP headers.
//
// This function expects the "x-api-key" header to be in the format: "ApiKey <key>".
// It returns the extracted API key as a string if the header is valid. If the header
// is missing, malformed, or does not start with "ApiKey", an error is returned.
//
// Parameters:
//   headers (http.Header): The HTTP headers from which the API key is extracted.
//
// Returns:
//   string: The extracted API key if the header is valid.
//   error: An error if the header is missing, malformed, or does not start with "ApiKey".
//
// Example usage:
//   key, err := GetAPIKey(req.Header)
//   if err != nil {
//       // handle error
//   }
//   // use the API key
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("x-api-key")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}