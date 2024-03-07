package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sebosun/chirpy/api"
)

func appRouterFS() (http.Handler, http.Handler) {
	fs := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	fs2 := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	return fs, fs2
}

func apiRouter(cfg *api.ApiConfig) *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.HandleFunc("/reset", cfg.HandleReset)
	apiRouter.Get("/healthz", api.HandleHealthz)

	apiRouter.Get("/chirps", api.HandleGetChirp)
	apiRouter.Post("/chirps", api.HandleCreateChirp)

	apiRouter.Get("/chirps/{id}", api.HandleGetChirpById)

	return apiRouter
}

func adminRouter(cfg *api.ApiConfig) *chi.Mux {
	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", cfg.HandleMetrics)
	return adminRouter
}
