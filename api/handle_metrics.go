package api

import (
	"fmt"
	"net/http"
)

func HandleMetrics(w http.ResponseWriter, r *http.Request, cfg *ApiConfig) {
	w.Header().Set("Content-Type: text/plain; charset=utf-8", "*")
	stringHits := fmt.Sprintf("%v", cfg.FileserverHits)
	cfgAsBytes := []byte(stringHits)

	w.Write(cfgAsBytes)
}
