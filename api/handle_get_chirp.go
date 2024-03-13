package api

import (
	"errors"
	"log"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/sebosun/chirpy/auth"
	"github.com/sebosun/chirpy/db"
)

func (cfg *ApiConfig) HandleGetChirp(w http.ResponseWriter, r *http.Request) {
	var chirps []db.Chirp
	queryAuthId := r.URL.Query().Get("author_id")
	querySort := r.URL.Query().Get("sort")
	if queryAuthId == "" {
		c, err := cfg.DB.GetChirps()
		chirps = c
		if err != nil {
			log.Println("Error accessing db", err)
			RespondWithError(w, 500, "Something went wrong")
			return
		}
	} else {
		c, err := cfg.DB.GetChirpsByAuthorId(queryAuthId)
		chirps = c
		if err != nil {
			log.Println("Error accessing db", err)
			RespondWithError(w, 500, "Something went wrong")
			return
		}

	}

	if querySort == "asc" {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].Id < chirps[j].Id
		})
	}

	if querySort == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].Id > chirps[j].Id
		})
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
