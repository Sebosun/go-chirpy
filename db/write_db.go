package db

import (
	"encoding/json"
	"os"
)

// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	dat, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	ok := os.WriteFile(db.path, dat, 0666)

	if ok != nil {
		return err
	}

	return nil
}
