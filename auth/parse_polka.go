package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ParsePolka(r http.Header) error {
	header := r.Get("Authorization")
	polkaSecret := os.Getenv("POLKA_KEY")
	fmt.Println(header, polkaSecret)

	if header == "" {
		return errors.New("Invalid authorization token")
	}

	headerSplit := strings.Split(header, " ")

	if len(headerSplit) != 2 || headerSplit[0] != "ApiKey" {
		return errors.New("Invalid authorization token")
	}
	if headerSplit[1] != polkaSecret {
		return errors.New("Invalid api key")
	}

	return nil
}
