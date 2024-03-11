package auth

import (
	"errors"
	"net/http"
	"strings"
)

func ParseBearer(r http.Header) (string, error) {
	header := r.Get("Authorization")

	if header == "" {
		return "", errors.New("Invalid authorization token")
	}

	bearerSplit := strings.Split(header, " ")

	if len(bearerSplit) != 2 || bearerSplit[0] != "Bearer" {
		return "", errors.New("Invalid authorization token")
	}

	return bearerSplit[1], nil
}
