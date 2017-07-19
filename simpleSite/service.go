package simpleSite

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type SaveToFileService struct {
}

type SaveToDBService struct {
	Repository *Repository
}

type Service interface {
	save(*User) error
	getUser(string) (*User, error)
}

func (svc SaveToFileService) save(u *User) error {
	filename := u.Username + ".txt"
	marshaledUser, err := json.Marshal(&u)

	if err != nil {
		errors.New("Error marshalling the user information")
	}

	return ioutil.WriteFile("data/"+filename, marshaledUser, 0600)
}

func (svc SaveToFileService) getUser(username string) (*User, error) {
	filename := username + ".txt"
	userAsJson, err := ioutil.ReadFile("data/" + filename)

	if err != nil {
		return nil, err
	}

	var u *User
	err = json.Unmarshal(userAsJson, &u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (svc SaveToDBService) save(u *User) error {
	user, _ := svc.Repository.findUserByUsername(u.Username)

	if user.Id == 0 {
		id, err := svc.Repository.createUser(u)
		u.Id = id
		return err
	} else {
		u.Id = user.Id
		err := svc.Repository.updateUser(u)
		return err
	}
}

func (svc SaveToDBService) getUser(username string) (*User, error) {
	return svc.Repository.findUserByUsername(username)
}
