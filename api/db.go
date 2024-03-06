package api

import (
	"encoding/json"
	"os"
	"sync"
)

type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

type Chirp struct {
	Id   uint64 `json:"id"`
	Body string `json:"body"`
}

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

/**/
/* // CreateChirp creates a new chirp and saves it to disk */
/* func (db *DB) CreateChirp(body string) (Chirp, error) { */
/**/
/* } */
/**/
/* // GetChirps returns all chirps in the database  */
/* func (db *DB) GetChirps() ([]Chirp, error) { */
/**/
/* } */
/**/
/* // ensureDB creates a new database file if it doesn't exist */
func (db *DB) ensureDB() error {
	_, err := os.ReadFile(db.path)
	if err != nil {
		defDatabse := DBStructure{}
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

// loadDB reads the database file into memory
func (db *DB) LoadDB() (DBStructure, error) {
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

/**/
/* // writeDB writes the database file to disk */
/* func (db *DB) writeDB(dbStructure DBStructure) error { */
/**/
/* } */
