package auth

import (
	"net/http"
	"os"
)

// chcks if user can use given resource
func HandleAuthUser(r *http.Request) (string, error) {
	authToken, err := ParseBearer(r.Header)
	if err != nil {
		return "", err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	usrId, err := ValidateAccessJWT(authToken, jwtSecret)

	if err != nil {
		return "", err
	}

	return usrId, nil
}
