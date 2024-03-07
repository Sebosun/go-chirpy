package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sebosun/chirpy/db"
)

func HandleGetChirp(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewDB("./database.json")

	if err != nil {
		log.Println("Error accessing db", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	chirps, err := db.GetChirps()

	if err != nil {
		log.Println("Error accessing db", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	RespondWithJSON(w, 200, chirps)
}

func HandleGetChirpById(w http.ResponseWriter, r *http.Request) {
	chirpId := chi.URLParam(r, "id")
	db, err := db.NewDB("./database.json")
	if err != nil {
		log.Println("Error accessing db", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	chirp, err := db.GetChirpsById(chirpId)
	if err != nil {
		RespondWithError(w, 404, err.Error())
		return
	}

	RespondWithJSON(w, 200, chirp)
}
