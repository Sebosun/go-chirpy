package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type ErrorMsg struct {
	Error string `json:"error"`
}

type SuccessMsg struct {
	Body string `json:"cleaned_body"`
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

func parseMsg(msg string) string {
	bannedWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	wordsSplit := strings.Split(msg, " ")
	acc := []string{}
	for _, word := range wordsSplit {

		appendWord := word
		for _, bannedWrd := range bannedWords {
			if strings.ToLower(word) == strings.ToLower(bannedWrd) {
				appendWord = "****"
			}
		}
		acc = append(acc, appendWord)
	}

	return strings.Join(acc, " ")
}
