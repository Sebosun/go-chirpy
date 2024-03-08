package db

import (
	"errors"
	"sort"
	"strconv"
)

/* // createChirp creates a new chirp and saves it to disk */
func (db *DB) CreateChirp(body string) (Chirp, error) {
	chirps, err := db.GetChirps()

	if err != nil {
		return Chirp{}, err
	}

	newId := 1

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].Id < chirps[j].Id
	})

	if len(chirps) > 0 {
		lastIndx := len(chirps) - 1
		newId = chirps[lastIndx].Id + 1
	}

	newChirp := Chirp{
		Id:   newId,
		Body: body,
	}

	dbMap, err := db.loadDB()
	dbMap.Chirps[newId] = newChirp

	err = db.writeDB(dbMap)
	if err != nil {
		return Chirp{}, err
	}

	return newChirp, nil
}

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
