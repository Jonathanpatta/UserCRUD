package main

import (
	"testing"
)

func TestUserCreate(t *testing.T) {
	_, err := CreateUser("", "", "", "", "")
	if err != errEmptyEmailError {
		t.Errorf("")
	}
	_, err = CreateUser("asdfsdfsdf", "", "", "", "")
	if err != errEmptyFirstNameError {
		t.Errorf("")
	}
	_, err = CreateUser("asdfsdfsdf", "asdfasdf", "", "", "")
	if err != nil {
		t.Errorf("Error occured")
	}
}

func TestUpdateUser(t *testing.T) {
	user, createError := CreateUser("asdfasdf", "asdfasdfsadf", "asdfasdfasdf", "asdfasdf", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}
	userWithoutEmail := new(User)
	userWithoutName := User{EmailAddress: "asdasdfasdf"}
	normalUser := User{EmailAddress: "asdfasdf", FirstName: "sadfasdf", PhoneNumber: "asdfoasdfjoasdf"}

	_, err := UpdateUser(user.UUID, userWithoutEmail)

	if err != errEmptyEmailError {
		t.Errorf("")
	}
	_, err = UpdateUser(user.UUID, &userWithoutName)
	if err != errEmptyFirstNameError {
		t.Errorf("")
	}

	_, err = UpdateUser(user.UUID, &normalUser)
	if err != nil {
		t.Errorf("Error with updating function")
	}
}

func TestGetUser(t *testing.T) {
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
	_, err := ListUsers()
	if err != nil {
		t.Errorf("error listing elements")
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("asdfsadf")
	if err != errUserDoesNotExistError {
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
