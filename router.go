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
	apiRouter.Get("/healthz", cfg.HandleHealthz)

	apiRouter.Get("/chirps", cfg.HandleGetChirp)
	apiRouter.Post("/chirps", cfg.HandleCreateChirp)
	apiRouter.Get("/chirps/{id}", cfg.HandleGetChirpById)

	apiRouter.Post("/users", cfg.HandleCreateUsers)
	apiRouter.Get("/users", cfg.HandleGetUsers)

	apiRouter.Post("/login", cfg.HandleLogin)

	return apiRouter
}

func adminRouter(cfg *api.ApiConfig) *chi.Mux {
	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", cfg.HandleMetrics)
	return adminRouter
}
