package db

import "sync"

type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps     map[int]Chirp     `json:"chirps"`
	Users      map[int]User      `json:"users"`
	RevokedJWT map[string]string `json:"revoked_jwt"`
}

type Chirp struct {
	Id     int    `json:"id"`
	Body   string `json:"body"`
	UserID int    `json:"author_id"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatedUserReturnVal struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
