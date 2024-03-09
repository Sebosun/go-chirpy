package api

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

type LoginParams struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	ExpiresSeconds int    `json:"expires_in_seconds"`
}

type LoginReturn struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (api *ApiConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := LoginParams{}
	err := decoder.Decode(&user)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	foundUser, err := api.DB.GetUserByEmail(user.Email)

	if err != nil {
		RespondWithError(w, 400, "Couldnt find user with given email")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		RespondWithError(w, 401, "Unauthorized")
		return
	}

	var expiresAt time.Time

	now := time.Now()
	if user.ExpiresSeconds == 0 {
		fullDay := time.Hour * 24
		expiresAt = time.Now().Add(fullDay)
	} else {
		expiresAt = time.Now().Add(time.Duration(user.ExpiresSeconds))
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Subject:   fmt.Sprintf("%v", foundUser.Id),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtString, err := jwtToken.SignedString([]byte(jwtSecret))

	if err != nil {
		fmt.Printf("Failed to parse jwtSecret %v", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	userRtrn := LoginReturn{
		Id:    foundUser.Id,
		Email: foundUser.Email,
		Token: jwtString,
	}

	RespondWithJSON(w, 200, userRtrn)

}
