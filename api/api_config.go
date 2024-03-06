package api

import (
	"fmt"
	"net/http"
)

type ApiConfig struct {
	FileserverHits int
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits += 1
		next.ServeHTTP(w, r)
	})

}

func (cfg *ApiConfig) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type: text/plain; charset=utf-8", "*")
	stringHits := fmt.Sprintf("%v", cfg.FileserverHits)
	cfgAsBytes := []byte(stringHits)

	w.Write(cfgAsBytes)
}
