package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type userBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *ApiConfig) HandleCreateUsers(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	user := userBody{}
	err := decoder.Decode(&user)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	createdUser, err := api.DB.CreateUser(user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating user:  %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	RespondWithJSON(w, 201, createdUser)
}

func (api *ApiConfig) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.DB.GetUsers()
	if err != nil {
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	RespondWithJSON(w, 200, users)
}
