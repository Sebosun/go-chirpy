package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sebosun/chirpy/api"
)

func main() {
	const port = "8080"
	apiConfig := api.ApiConfig{}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	corsMux := middlewareCors(router)
	db, err := api.NewDB("./databse.json")
	if err != nil {
		log.Fatal(err)
	}
	dbMem, err := db.LoadDB()

	if err != nil {
		log.Fatal()
	}

	for _, chirp := range dbMem.Chirps {
		fmt.Println(chirp)
	}

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
