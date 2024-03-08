package api

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (cfg *ApiConfig) HandleGetChirp(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.DB.GetChirps()

	if err != nil {
		log.Println("Error accessing db", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	RespondWithJSON(w, 200, chirps)
}

func (cfg *ApiConfig) HandleGetChirpById(w http.ResponseWriter, r *http.Request) {
	chirpId := chi.URLParam(r, "id")
	chirp, err := cfg.DB.GetChirpsById(chirpId)

	if err != nil {
		RespondWithError(w, 404, err.Error())
		return
	}

	RespondWithJSON(w, 200, chirp)
}
