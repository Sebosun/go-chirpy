package db

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

type CreatedUserReturnVal struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func (db *DB) CreateUser(email string, password string) (CreatedUserReturnVal, error) {
	users, err := db.GetUsers()
	if err != nil {
		return CreatedUserReturnVal{}, err
	}

	newId := 1

	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})

	if len(users) > 0 {
		lastIndx := len(users) - 1
		newId = users[lastIndx].Id + 1
	}

	hashedPaswd, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return CreatedUserReturnVal{}, err
	}

	newUser := User{
		Id:       newId,
		Email:    email,
		Password: string(hashedPaswd),
	}

	dbMap, err := db.loadDB()
	dbMap.Users[newId] = newUser

	err = db.writeDB(dbMap)

	if err != nil {
		return CreatedUserReturnVal{}, err
	}

	userReturned := CreatedUserReturnVal{
		Id:    newUser.Id,
		Email: newUser.Email,
	}

	return userReturned, nil
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

func (db *DB) GetUserByEmail(email string) (User, error) {
	dbMem, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	for _, val := range dbMem.Users {
		if val.Email == email {
			return val, nil
		}
	}

	return User{}, nil
}

func (db *DB) UpdateUsers(id int, email string, password string) error {
	users, err := db.GetUsers()

	if err != nil {
		return err
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})

	userExists := false
	var curUser User

	for _, v := range users {
		if v.Id == id {
			userExists = true
			curUser = v
		}
	}

	if !userExists {
		return errors.New("User doesnt exist")
	}

	if email != "" {
		curUser.Email = email
	}

	if password != "" {
		hashedPaswd, err := bcrypt.GenerateFromPassword([]byte(password), 4)
		if err != nil {
			return err
		}
		curUser.Password = string(hashedPaswd)
	}

	dbMap, err := db.loadDB()

	if err != nil {
		return err
	}

	dbMap.Users[id] = curUser

	err = db.writeDB(dbMap)

	if err != nil {
		return err
	}

	return nil
}
