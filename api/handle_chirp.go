package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type parameters struct {
	Message string `json:"body"`
}

func (cfg *ApiConfig) HandleCreateChirp(w http.ResponseWriter, r *http.Request) {
	const maxMsgLen = 140

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	if len(params.Message) > maxMsgLen {
		RespondWithError(w, 400, "Chirp is too long")
		return
	}

	item, err := cfg.DB.CreateChirp(parseMsg(params.Message))

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
	}

	RespondWithJSON(w, 201, item)
}
