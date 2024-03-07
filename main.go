package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sebosun/chirpy/api"
	"github.com/sebosun/chirpy/db"
)

func main() {
	const port = "8080"
	apiConfig := api.ApiConfig{}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	corsMux := middlewareCors(router)
	db, err := db.NewDB("./database.json")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateChirp("test chirp")

	chirps, err := db.GetChirps()

	fmt.Println(chirps)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	fs, fs2 := appRouterFS()

	router.Handle("/app/*", apiConfig.MiddlewareMetricsInc(fs))
	router.Handle("/app", apiConfig.MiddlewareMetricsInc(fs2))

	router.Mount("/api/", apiRouter(&apiConfig))
	router.Mount("/admin/", adminRouter(&apiConfig))

	fmt.Printf("Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
