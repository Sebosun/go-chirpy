package api

import (
	"encoding/json"
	"fmt"
	"github.com/sebosun/chirpy/auth"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReturn struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	TokenRefresh string `json:"refresh_token"`
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

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtAccessStr, err := auth.CreateAccessJWT(foundUser.Id, jwtSecret)
	jwtRefreshStr, err := auth.CreateRefreshJWT(foundUser.Id, jwtSecret)

	if err != nil {
		fmt.Printf("Failed to parse jwtSecret %v", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	userRtrn := LoginReturn{
		Id:           foundUser.Id,
		Email:        foundUser.Email,
		Token:        jwtAccessStr,
		TokenRefresh: jwtRefreshStr,
	}

	RespondWithJSON(w, 200, userRtrn)
}
