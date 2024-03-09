package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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

type MyCustomClaims struct {
	Issuer  string `json:"Issuer"`
	Subject int    `json:"Subject"`
	jwt.RegisteredClaims
}

type EditParams struct {
	Message string `json:"body"`
}

type PutParameters struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *ApiConfig) HandlePutUsers(w http.ResponseWriter, r *http.Request) {
	headers := r.Header.Get("Authorization")

	if headers == "" {
		RespondWithError(w, 401, "Missing authorization token ")
		return
	}

	authToken, err := parseBearer(headers)
	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	token, err := jwt.ParseWithClaims(authToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	claims := token.Claims.(*MyCustomClaims)
	id, err := claims.GetSubject()

	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
	}

	if err != nil {
		RespondWithError(w, 401, "Invalid authorization token")
		return
	}

	fmt.Println(id)

	decoder := json.NewDecoder(r.Body)
	params := PutParameters{}
	err = decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	api.DB.UpdateUsers(idInt, params.Email, params.Password)

	w.WriteHeader(200)
}
