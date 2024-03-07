package db

import (
	"errors"
	"strconv"
)

// GetChirps returns all chirps in the database
func (db *DB) GetChirps() ([]Chirp, error) {
	dbMem, err := db.loadDB()
	if err != nil {
		return nil, nil
	}
	acc := []Chirp{}

	for _, chirp := range dbMem.Chirps {
		acc = append(acc, chirp)
	}
	return acc, nil
}

func (db *DB) GetChirpsById(desiredId string) (Chirp, error) {
	searchId, err := strconv.Atoi(desiredId)

	if err != nil {
		return Chirp{}, err
	}

	dbMem, err := db.loadDB()
	if err != nil {
		return Chirp{}, nil
	}

	for _, chirp := range dbMem.Chirps {
		if chirp.Id == searchId {
			return chirp, nil
		}
	}

	return Chirp{}, errors.New("Coudln't find chirp")
}
