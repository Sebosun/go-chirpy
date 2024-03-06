package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type parameters struct {
	Message string `json:"body"`
}

type ErrorMsg struct {
	Error string `json:"error"`
}

type SuccessMsg struct {
	Valid bool `json:"valid"`
}

func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
	const maxMsgLen = 140

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		errBody := ErrorMsg{
			Error: "Something went wrong",
		}
		data, err := json.Marshal(errBody)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

		return

	}

	if len(params.Message) > maxMsgLen {
		errBody := ErrorMsg{
			Error: "Chirp is too long",
		}
		data, err := json.Marshal(errBody)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	succesBody := SuccessMsg{
		Valid: true,
	}

	data, err := json.Marshal(succesBody)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
