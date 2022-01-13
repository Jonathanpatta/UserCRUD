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

const EmptyEmailError string = "email address cannot be empty"
const EmptyFirstNameError string = "first name cannot be empty"
const UserDoesNotExistError string = "user does not exist"

var UserStore []User

func CreateUser(emailAddr string, firstName string, lastName string, phoneNumber string, dob string) (user User, err error) {
	if emailAddr == "" {
		err = errors.New(EmptyEmailError)
		return user, err
	} else if firstName == "" {
		err = errors.New(EmptyFirstNameError)
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

func GetUser(id string) (user User, err error) {
	for _, val := range UserStore {
		if val.UUID == id {
			return val, err
		}
	}
	return user, errors.New(UserDoesNotExistError)
}

func UpdateUser(id string, updatedUser User) (user User, err error) {
	for i, val := range UserStore {
		if val.UUID == id {
			if updatedUser.EmailAddress == "" {
				err = errors.New(EmptyEmailError)
				return user, err
			} else if updatedUser.FirstName == "" {
				err = errors.New(EmptyFirstNameError)
				return user, err
			}
			updatedUser.UUID = id
			UserStore[i] = updatedUser
			return updatedUser, err
		}
	}
	//asdfsdf
	return user, errors.New(UserDoesNotExistError)
}

func ListUsers() (user []User, err error) {
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
	return errors.New(UserDoesNotExistError)
}

func main() {
	newUser, _ := CreateUser("jonathan.patta@gmail.com", "jonathan", "patta", "", "")
	fmt.Println(newUser)
	returnedUser, _ := GetUser(newUser.UUID)
	fmt.Println(returnedUser)
	CreateUser("sadfasdfsadf", "sadfsdf", "safasdfadsf", "asdfasdfasdf", "asfasdf")
	CreateUser("asfadsfasdf", "asdfa", "asdfsadf", "asdf", "asdf")
	CreateUser("asfasdfsadf", "asdfasdf", "asdfsad", "asdf", "asdf")
	CreateUser("sadf", "sadfasdf", "adfsasd", "asdfsadf", "sadfsadf")

	UpdateUser(newUser.UUID, User{FirstName: "biscuit", EmailAddress: "b@c.com"})

	returnedUser, _ = GetUser(newUser.UUID)
	fmt.Println(returnedUser)

	DeleteUser(newUser.UUID)

	fmt.Println(ListUsers())

}
