package main

import (
	"testing"
)

func TestUserCreate(t *testing.T) {
	_, err := CreateUser("", "", "", "", "")
	if err.Error() != EmptyEmailError {
		t.Errorf("")
	}
	_, err = CreateUser("asdfsdfsdf", "", "", "", "")
	if err.Error() != EmptyFirstNameError {
		t.Errorf("")
	}
	_, err = CreateUser("asdfsdfsdf", "asdfasdf", "", "", "")
	if err != nil {
		t.Errorf("Error occured")
	}
}

func TestGetUser(t *testing.T) {
	user, createError := CreateUser("asdfsadf", "asdfasdf", "asdf", "0", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}
	user, err := GetUser(user.UUID)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestListUser(t *testing.T) {
	_, err := ListUsers()
	if err != nil {
		t.Errorf("error listing elements")
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("asdfsadf")
	if err.Error() != UserDoesNotExistError {
		t.Errorf("Not failing when providing bad id")
	}
	user, createError := CreateUser("asdfs", "asdf", "asdf", "", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}

	err = DeleteUser(user.UUID)

	if err != nil {
		t.Errorf(err.Error())
	}

}
