package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sebosun/chirpy/api"
)

func main() {
	const port = "8080"
	apiConfig := api.ApiConfig{}
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	fs := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", apiConfig.MiddlewareMetricsInc(fs))
	mux.HandleFunc("/reset", apiConfig.HandleReset)
	mux.HandleFunc("/metrics", apiConfig.HandleMetrics)

	mux.HandleFunc("/healthz", api.HandleHealthz)

	fmt.Printf("Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
