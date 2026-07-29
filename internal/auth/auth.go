package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API Key from the headers of an HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("No Auth Info Found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Auth Header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed First Part of Auth Header")
	}

	return vals[1], nil
}