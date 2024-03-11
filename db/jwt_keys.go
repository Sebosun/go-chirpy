package db

import (
	"errors"
)

func (db *DB) GetRevokedJWTs() ([]string, error) {
	dbMem, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	var acc []string
	for _, val := range dbMem.RevokedJWT {
		acc = append(acc, val)
	}
	return acc, nil
}

func (db *DB) CheckIsJWTRevoked(jwt string) (bool, error) {
	hasJWT := false

	jwts, err := db.GetRevokedJWTs()
	if err != nil {
		return false, err
	}

	for _, val := range jwts {
		if val == jwt {
			hasJWT = true
			break
		}
	}

	return hasJWT, nil
}

func (db *DB) CreateRevokedJWTs(revokedJWT string) error {
	hasJwt, err := db.CheckIsJWTRevoked(revokedJWT)

	if err != nil {
		return err
	}

	if hasJwt {
		return errors.New("JWT has already been revoked")
	}

	dbMap, err := db.loadDB()
	if err != nil {
		return err
	}
	dbMap.RevokedJWT[revokedJWT] = revokedJWT
	err = db.writeDB(dbMap)
	if err != nil {
		return err
	}

	return nil
}
