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
	w.Header().Set("Content-Type: text/html; charset=utf-8", "*")
	stringHits := fmt.Sprintf(
		`<html>

<body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
</body>

</html>
`, cfg.FileserverHits)
	cfgAsBytes := []byte(stringHits)

	w.Write(cfgAsBytes)
}

func (cfg *ApiConfig) HandleReset(w http.ResponseWriter, r *http.Request) {
	cfg.FileserverHits = 0
	w.Header().Set("Content-Type: text/plain; charset=utf-8", "*")
	w.Write([]byte("OK"))
}
