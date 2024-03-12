package api

import (
	"encoding/json"
	"errors"
	"github.com/sebosun/chirpy/auth"
	"github.com/sebosun/chirpy/db"
	"net/http"
)

type WebhookPayload struct {
	Event string      `json:"event"`
	Data  WebhookData `json:"data"`
}

type WebhookData struct {
	UserId int `json:"user_id"`
}

func (api *ApiConfig) HandlePolkaWebhook(w http.ResponseWriter, r *http.Request) {
	err := auth.ParsePolka(r.Header)
	if err != nil {
		RespondWithError(w, 401, "Invalid API key")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := WebhookPayload{}
	err = decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 500, "Error decoding json")
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(200)
		return
	}

	err = api.DB.UpgradeUserPremium(params.Data.UserId)

	if errors.Is(err, db.ErrorUserDoesntExist()) {
		RespondWithError(w, 404, err.Error())
	} else if err != nil {
		RespondWithError(w, 500, "Error decoding json")
		return
	}

	w.WriteHeader(200)
}
