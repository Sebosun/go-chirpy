package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	fs := http.FileServer(http.Dir("."))
	mux.Handle("/app", fs)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type: text/plain; charset=utf-8", "*")
		var isOk []byte = []byte("OK")
		w.Write(isOk)

	})

	fmt.Printf("Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
