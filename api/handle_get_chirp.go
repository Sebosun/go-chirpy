package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sebosun/chirpy/auth"
	"github.com/sebosun/chirpy/db"
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

func (cfg *ApiConfig) HandleDeleteChirpById(w http.ResponseWriter, r *http.Request) {
	usrId, err := auth.HandleAuthUser(r)
	if err != nil {
		RespondWithError(w, 401, err.Error())
		return
	}

	chirpId := chi.URLParam(r, "id")

	err = cfg.DB.DeleteChirpsById(chirpId, usrId)

	if errors.Is(err, db.MsgNotBelongUser()) {
		RespondWithError(w, 403, err.Error())
		return
	} else if err != nil {
		RespondWithError(w, 403, "Unauthorized request")
		return
	}

	w.WriteHeader(200)
}
