package main

import (
	"fmt"
	"testing"
)

func TestUserCreate(t *testing.T) {
	DBConnect()
	_, err := CreateUser("", "", "", "", "")
	if err != errEmptyEmailError {
		t.Errorf("")
	}
	_, err = CreateUser("asdfsdfsdf", "", "", "", "")
	if err != errEmptyFirstNameError {
		t.Errorf("")
	}

	_, err = CreateUser("asdfsdfsdf", "asdfasdf", "", "", "")

	fmt.Println("hello")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	// DBConnect()
	user, createError := CreateUser("asdfasdf", "asdfasdfsadf", "asdfasdfasdf", "asdfasdf", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}

	normalUser := User{EmailAddress: "asdfasdf", FirstName: "sadfasdf", PhoneNumber: "asdfoasdfjoasdf"}

	_, err := UpdateUser(user.UUID, &normalUser)
	if err != nil {
		t.Errorf("Error with updating function")
	}
}

func TestGetUser(t *testing.T) {
	// DBConnect()
	user, createError := CreateUser("asdfsadf", "asdfasdf", "asdf", "0", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}
	id := user.UUID
	_, err := GetUser(id)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestListUser(t *testing.T) {
	// DBConnect()
	_, err := ListUsers()
	if err != nil {
		t.Errorf("error listing elements")
	}
}

func TestDeleteUser(t *testing.T) {
	// DBConnect()
	user, createError := CreateUser("asdfs", "asdf", "asdf", "", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}

	err := DeleteUser(user.UUID)

	if err != nil {
		t.Errorf(err.Error())
	}

}
