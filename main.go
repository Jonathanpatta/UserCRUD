package main

import "errors"

type User struct {
	EmailAddress string
	FirstName    string
	LastName     string
	PhoneNumber  int64
	DOB          string
	UUID         string
}

var UserStore []User

func CreateUser(emailAddr string, firstName string, lastName string, phoneNumber int64, dob string) (user User, err error) {
	if emailAddr == "" {
		err = errors.New("Email Address cannot be empty")
	}
}

func main() {
}
