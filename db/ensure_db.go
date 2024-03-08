package db

import (
	"encoding/json"
	"os"
)

func (db *DB) ensureDB() error {
	_, err := os.ReadFile(db.path)
	if err != nil {

		defDatabse := DBStructure{
			Chirps: make(map[int]Chirp),
			Users:  make(map[int]User),
		}

		dat, err := json.Marshal(defDatabse)

		if err != nil {
			return err
		}

		ok := os.WriteFile(db.path, dat, 0666)
		if ok != nil {
			return err
		}
	}

	return nil
}
