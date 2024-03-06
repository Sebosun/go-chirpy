package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type parameters struct {
	Message string `json:"body"`
}

func validateMsg(msg string) string {
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

func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
	const maxMsgLen = 140

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	if len(params.Message) > maxMsgLen {
		RespondWithError(w, 400, "Chirp is too long")
		return
	}

	succesBody := SuccessMsg{
		Body: validateMsg(params.Message),
	}

	RespondWithJSON(w, 200, succesBody)
}
