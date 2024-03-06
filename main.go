package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sebosun/chirpy/api"
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	apiConfig := api.ApiConfig{}
	router := chi.NewRouter()
	corsMux := middlewareCors(router)
	router.Use(middleware.Logger)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	fs := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	router.Handle("/app/*", apiConfig.MiddlewareMetricsInc(fs))

	fs2 := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	router.Handle("/app", apiConfig.MiddlewareMetricsInc(fs2))

	apiRouter := chi.NewRouter()
	apiRouter.HandleFunc("/reset", apiConfig.HandleReset)
	apiRouter.Get("/metrics", apiConfig.HandleMetrics)
	apiRouter.Get("/healthz", api.HandleHealthz)

	router.Mount("/api", apiRouter)

	fmt.Printf("Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
