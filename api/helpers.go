package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorMsg struct {
	Error string `json:"error"`
}

type SuccessMsg struct {
	Valid bool `json:"valid"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	errBody := ErrorMsg{Error: msg}
	RespondWithJSON(w, code, errBody)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)

}
