package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sebosun/chirpy/auth"
	"github.com/sebosun/chirpy/db"
)

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *ApiConfig) HandleCreateUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := UserCreateBody{}
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

type PutParameters struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *ApiConfig) HandlePutUsers(w http.ResponseWriter, r *http.Request) {
	headers := r.Header.Get("Authorization")

	if headers == "" {
		RespondWithError(w, 401, "Authorization token missing header")
		return
	}

	authToken, err := parseBearer(headers)
	fmt.Println(authToken)
	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	val, err := auth.ValidateAccessJWT(authToken, jwtSecret)

	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	idInt, err := strconv.Atoi(val)

	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := PutParameters{}
	err = decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	api.DB.UpdateUsers(idInt, params.Email, params.Password)
	usrRtnr := db.CreatedUserReturnVal{
		Id:    idInt,
		Email: params.Email,
	}

	RespondWithJSON(w, 200, usrRtnr)
}
