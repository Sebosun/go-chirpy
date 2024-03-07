package db

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
