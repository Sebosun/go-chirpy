package db

import (
	"encoding/json"
	"os"
)

// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	file, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, err
	}

	defDatabse := DBStructure{}
	err = json.Unmarshal(file, &defDatabse)

	if err != nil {
		return DBStructure{}, err
	}

	return defDatabse, nil
}
