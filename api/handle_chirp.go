package api

import (
	"encoding/json"
	"github.com/sebosun/chirpy/auth"
	"log"
	"net/http"
	"os"
)

type ChirpParams struct {
	Message string `json:"body"`
}

func (cfg *ApiConfig) HandleCreateChirp(w http.ResponseWriter, r *http.Request) {
	authToken, err := auth.ParseBearer(r.Header)
	if err != nil {
		RespondWithError(w, 401, err.Error())
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	usrId, err := auth.ValidateAccessJWT(authToken, jwtSecret)

	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	const maxMsgLen = 140

	decoder := json.NewDecoder(r.Body)
	params := ChirpParams{}
	err = decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	if len(params.Message) > maxMsgLen {
		RespondWithError(w, 400, "Chirp is too long")
		return
	}

	item, err := cfg.DB.CreateChirp(parseMsg(params.Message), usrId)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
	}

	RespondWithJSON(w, 201, item)
}
