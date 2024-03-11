package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const AccessName = "chirpy-access"
const RefreshName = "chirpy-refresh"

func CreateAccessJWT(userId int, secretToken string) (string, error) {
	expirationTime := time.Hour * 24 // day
	val, err := CreateJWT(AccessName, expirationTime, userId, secretToken)
	if err != nil {
		return "", err
	}
	return val, nil
}

func CreateRefreshJWT(userId int, secretToken string) (string, error) {
	expirationTime := time.Hour * 24 * 60 // 60 days
	val, err := CreateJWT(RefreshName, expirationTime, userId, secretToken)
	if err != nil {
		return "", err
	}
	return val, nil
}

func CreateJWT(issuerName string, expirationTime time.Duration, userId int, secretToken string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuerName,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expirationTime)),
		Subject:   fmt.Sprintf("%v", userId),
	})

	jwtString, err := jwtToken.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func ValidateAccessJWT(tokenString, tokenSecret string) (string, error) {
	val, err := ValidateJWT(tokenString, tokenSecret, AccessName)
	if err != nil {
		return "", err
	}
	return val, nil
}

func ValidateRefreshJWT(tokenString, tokenSecret string) (string, error) {
	val, err := ValidateJWT(tokenString, tokenSecret, RefreshName)
	if err != nil {
		return "", err
	}
	return val, nil
}

func ValidateJWT(tokenString, tokenSecret, validInsuerName string) (string, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return "", err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", err
	}

	if issuer != validInsuerName {
		return "", errors.New("Issuer does not grant access")
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return userIDString, nil
}
