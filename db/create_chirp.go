package db

import (
	"sort"
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
