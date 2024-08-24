package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the request headers
// Example:
// Authorization: API_KEY <key>
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")

	if apiKey == "" {
		return "", errors.New("no authentication header found")
	}

	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("malfomed authentication header")
	}
	if vals[0] != "API_KEY" {
		return "", errors.New("invalid authentication type")
	}

	return vals[1], nil
}
