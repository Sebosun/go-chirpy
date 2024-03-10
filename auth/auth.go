package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateJWT(expirationTime int, userId int, secretToken string) (string, error) {
	expiresAt := time.Hour * 24

	if expirationTime > 0 {
		expiresAt = time.Duration(expirationTime)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresAt)),
		Subject:   fmt.Sprintf("%v", userId),
	})

	jwtString, err := jwtToken.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (string, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return "", err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return userIDString, nil
}
