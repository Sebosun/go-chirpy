package api

import (
	"encoding/json"
	"github.com/sebosun/chirpy/db"
	"log"
	"net/http"
)

type parameters struct {
	Message string `json:"body"`
}

func HandleCreateChirp(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewDB("./database.json")
	if err != nil {
		log.Println("Error reading from the database", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	const maxMsgLen = 140

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
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

	item, err := db.CreateChirp(parseMsg(params.Message))

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
	}

	RespondWithJSON(w, 201, item)
}

func HandleGetChirp(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewDB("./database.json")
	if err != nil {
		log.Println("Error accessing db", err)
		RespondWithError(w, 500, "Something went wrong")
	}

	chirps, err := db.GetChirps()

	if err != nil {
		log.Println("Error accessing db", err)
		RespondWithError(w, 500, "Something went wrong")
	}

	RespondWithJSON(w, 200, chirps)
}
