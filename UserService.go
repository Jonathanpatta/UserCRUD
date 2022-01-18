package main

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	EmailAddress string
	FirstName    string
	LastName     string
	PhoneNumber  string
	DOB          string
	UUID         string
}

var errEmptyEmailError = errors.New("email address cannot be empty")
var errEmptyFirstNameError = errors.New("first name cannot be empty")
var errUserDoesNotExistError = errors.New("user does not exist")

var UserStore []*User

func CreateUser(emailAddr string, firstName string, lastName string, phoneNumber string, dob string) (user *User, err error) {
	if emailAddr == "" {
		return user, errEmptyEmailError
	} else if firstName == "" {
		return user, errEmptyFirstNameError
	}
	id, idCreationError := uuid.NewRandom()
	if idCreationError != nil {
		return user, idCreationError
	}
	myuser := User{}
	myuser.UUID = id.String()
	myuser.EmailAddress = emailAddr
	myuser.FirstName = firstName
	myuser.LastName = lastName
	myuser.PhoneNumber = phoneNumber
	myuser.DOB = dob

	UserStore = append(UserStore, &myuser)

	return &myuser, err
}

func GetUser(id string) (user *User, err error) {
	for _, val := range UserStore {
		if val.UUID == id {
			return val, err
		}
	}
	return user, errUserDoesNotExistError
}

func UpdateUser(id string, updatedUser *User) (user *User, err error) {

	for i, val := range UserStore {
		if val.UUID == id {
			if updatedUser.EmailAddress == "" {
				return user, errEmptyEmailError
			} else if updatedUser.FirstName == "" {
				return user, errEmptyFirstNameError
			}
			updatedUser.UUID = id
			UserStore[i] = updatedUser
			return updatedUser, err
		}
	}

	return user, errUserDoesNotExistError
}

func ListUsers() (user []*User, err error) {
	return UserStore, err
}

func DeleteUser(id string) (err error) {
	for i, val := range UserStore {
		if val.UUID == id {
			UserStore[i] = UserStore[len(UserStore)-1]
			UserStore = UserStore[:len(UserStore)-1]
			return err
		}
	}
	return errUserDoesNotExistError
}
