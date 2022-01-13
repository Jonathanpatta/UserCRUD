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

func GetUser(id string) (user User, err error) {
	for _, val := range UserStore {
		if val.UUID == id {
			return val, err
		}
	}
	return user, errors.New("no user with that id exists")
}

func UpdateUser(id string, updatedUser User) (user User, err error) {
	for i, val := range UserStore {
		if val.UUID == id {
			if updatedUser.EmailAddress == "" {
				err = errors.New("email address cannot be empty")
				return user, err
			} else if updatedUser.FirstName == "" {
				err = errors.New("first name cannot be empty")
				return user, err
			}
			updatedUser.UUID = id
			UserStore[i] = updatedUser
			return updatedUser, err
		}
	}
	return user, errors.New("no user with that id exists")
}

func ListUsers() []User {
	return UserStore
}

func DeleteUser(id string) (err error) {
	for i, val := range UserStore {
		if val.UUID == id {
			UserStore[i] = UserStore[len(UserStore)-1]
			UserStore = UserStore[:len(UserStore)-1]
		}
	}
	return err
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

	fmt.Println(ListUsers())

}
