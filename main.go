package main

import (
	"errors"
	"fmt"
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

var UserStore []User

func CreateUser(emailAddr string, firstName string, lastName string, phoneNumber string, dob string) (user User, err error) {
	if emailAddr == "" {
		err = errors.New("email address cannot be empty")
		return user, err
	} else if firstName == "" {
		err = errors.New("first name cannot be empty")
		return user, err
	}
	id, idCreationError := uuid.NewRandom()
	if idCreationError != nil {
		return user, idCreationError
	}
	user.UUID = id.String()
	user.EmailAddress = emailAddr
	user.FirstName = firstName
	user.LastName = lastName
	user.PhoneNumber = phoneNumber
	user.DOB = dob

	UserStore = append(UserStore, user)

	return user, err
}

func main() {
	fmt.Println(CreateUser("jonathan.patta@gmail.com", "jonathan", "patta", "", ""))
}
