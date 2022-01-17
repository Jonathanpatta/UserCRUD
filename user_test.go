package main

import (
	"testing"
)

func TestUserCreate1(t *testing.T) {
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

func TestUpdateUser1(t *testing.T) {
	user, createError := CreateUser("asdfasdf", "asdfasdfsadf", "asdfasdfasdf", "asdfasdf", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}
	userWithoutEmail := User{}
	userWithoutName := User{EmailAddress: "asdasdfasdf"}
	normalUser := User{EmailAddress: "asdfasdf", FirstName: "sadfasdf", PhoneNumber: "asdfoasdfjoasdf"}

	_, err := UpdateUser(user.UUID, userWithoutEmail)

	if err.Error() != EmptyEmailError {
		t.Errorf("")
	}
	_, err = UpdateUser(user.UUID, userWithoutName)
	if err.Error() != EmptyFirstNameError {
		t.Errorf("")
	}

	_, err = UpdateUser(user.UUID, normalUser)
	if err != nil {
		t.Errorf("Error with updating function")
	}
}

<<<<<<< HEAD
func TestGetUser1(t *testing.T) {
=======
func TestGetUser(t *testing.T) {
>>>>>>> 76376a2f6ce54106062dacfc0298baee29a397c2
	user, createError := CreateUser("asdfsadf", "asdfasdf", "asdf", "0", "")
	if createError != nil {
		t.Errorf("User Creation error")
	}
	id := user.UUID
	user, err := GetUser(id)
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

func TestDeleteUser1(t *testing.T) {
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
