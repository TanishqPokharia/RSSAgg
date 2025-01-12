package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract api key from headers

func GetApiKey(headers *http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", errors.New("api key not found")
	}

	authSlice := strings.Split(auth, " ")

	if len(authSlice) != 2 {
		return "", errors.New("invalid authorization format")
	}

	if authSlice[0] != "ApiKey" {
		return "", errors.New("malformed Authorization header")
	}

	return authSlice[1], nil
}
