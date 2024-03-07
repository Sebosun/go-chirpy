package db

import (
	"sync"
)

func NewDB(path string) (*DB, error) {
	db := DB{
		path: "./database.json",
		mux:  &sync.RWMutex{},
	}

	err := db.ensureDB()

	if err != nil {
		return &DB{}, err
	}

	return &db, nil
}
