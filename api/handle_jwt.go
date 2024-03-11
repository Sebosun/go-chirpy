package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/sebosun/chirpy/auth"
)

type Token struct {
	Token string `json:"token"`
}

func (api *ApiConfig) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	authToken, err := auth.ParseBearer(r.Header)
	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	userId, err := auth.ValidateRefreshJWT(authToken, jwtSecret)
	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	// TODO: put str conversion into createJWT access
	idInt, err := strconv.Atoi(userId)
	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	jwtString, err := auth.CreateAccessJWT(idInt, jwtSecret)
	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	RespondWithJSON(w, 200, Token{Token: jwtString})
}

/* func (api *ApiConfig) HandleRevoke(w http.ResponseWriter, r *http.Request) { */
/* } */
