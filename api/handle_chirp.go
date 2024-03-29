package api

import (
	"encoding/json"
	"github.com/sebosun/chirpy/auth"
	"log"
	"net/http"
)

type ChirpParams struct {
	Message string `json:"message"`
}

func (cfg *ApiConfig) HandleCreateChirp(w http.ResponseWriter, r *http.Request) {
	usrId, err := auth.HandleAuthUser(r)
	if err != nil {
		RespondWithError(w, 401, err.Error())
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

	if len(params.Message) == 0 {
		RespondWithError(w, 400, "Chirp is too short")
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
