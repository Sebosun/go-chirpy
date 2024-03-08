package db

import "sort"

func (db *DB) CreateUser(email string) (User, error) {
	users, err := db.GetUsers()
	if err != nil {
		return User{}, err
	}

	newId := 1

	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})

	if len(users) > 0 {
		lastIndx := len(users) - 1
		newId = users[lastIndx].Id + 1
	}

	newUser := User{
		Id:    newId,
		Email: email,
	}

	dbMap, err := db.loadDB()
	dbMap.Users[newId] = newUser

	err = db.writeDB(dbMap)

	if err != nil {
		return User{}, err
	}

	return newUser, nil

}

func (db *DB) GetUsers() ([]User, error) {
	dbMem, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	var acc []User
	for _, val := range dbMem.Users {
		acc = append(acc, val)
	}
	return acc, nil
}
