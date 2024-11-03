package utils

import (
	"errors"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

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

func GetAccessToken(headers http.Header) (string, error) {
	authHeader := headers.Get("x-auth-key")
	if authHeader == "" {
		return "", errors.New("no auth info found")
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return "", errors.New("malformed auth header")
	}

	if authHeaderParts[0] != "Bearer" {
		return "", errors.New("malformed first part of auth header")
	}

	return authHeaderParts[1], nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
