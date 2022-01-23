package main

import (
	"testing"
	"time"
)

func TestUserCreate(t *testing.T) {
	DBConnect()

	_, err := CreateUser("", "", "", "", time.Time{})
	if err != ErrEmptyEmail {
		t.Errorf("")
	}
	_, err = CreateUser("Test-firstname", "", "", "", time.Time{})
	if err != ErrEmptyFirstName {
		t.Errorf("")
	}

	_, err = CreateUser("Test-firstName-2", "asdfasdf", "", "", time.Time{})
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	// DBConnect()
	user, createError := CreateUser("Test-email1", "fname1", "asdfasdfasdf", "asdfasdf", time.Time{})
	if createError != nil {
		t.Errorf("User Creation error")
	}

	normalUser := User{EmailAddress: "Test-email2", FirstName: "fname2", PhoneNumber: "asdfoasdfjoasdf"}

	_, err := UpdateUser(user.UUID, &normalUser)
	if err != nil {
		t.Errorf("Error with updating function")
	}
}

func TestGetUser(t *testing.T) {
	// DBConnect()
	user, createError := CreateUser("asdfsadf", "asdfasdf", "asdf", "0", time.Time{})
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
	user, createError := CreateUser("asdfs", "asdf", "asdf", "", time.Time{})
	if createError != nil {
		t.Errorf("User Creation error")
	}

	err := DeleteUser(user.UUID)

	if err != nil {
		t.Errorf(err.Error())
	}

}
