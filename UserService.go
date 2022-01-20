package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	EmailAddress string `json:"emailaddress,omitempty"`
	FirstName    string `json:"firstname,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	PhoneNumber  string `json:"phonenumber,omitempty"`
	DOB          string `json:"dob,omitempty"`
	UUID         string `json:"UUID,omitempty"`
}

var errEmptyEmailError = errors.New("email address cannot be empty")
var errEmptyFirstNameError = errors.New("first name cannot be empty")
var errUserDoesNotExistError = errors.New("user does not exist")

var UserStore []*User

func CreateUser(emailAddr string, firstName string, lastName string, phoneNumber string, dob string) (*User, error) {
	if emailAddr == "" {
		return nil, errEmptyEmailError
	} else if firstName == "" {
		return nil, errEmptyFirstNameError
	}
	id, idCreationError := uuid.NewRandom()
	if idCreationError != nil {
		return nil, idCreationError
	}
	myuser := User{}
	myuser.UUID = id.String()
	myuser.EmailAddress = emailAddr
	myuser.FirstName = firstName
	myuser.LastName = lastName
	myuser.PhoneNumber = phoneNumber
	myuser.DOB = dob

	UserStore = append(UserStore, &myuser)

	fields, values := GetFieldsAndValues(&myuser)

	query := `INSERT INTO ` + UserTableName + fields + `VALUES` + values + ` returning *;`

	fmt.Println(query)
	rows, err := db.Query(query)

	var user User
	var lastName_ sql.NullString
	var dateOfBirth sql.NullTime
	var phoneNumber_ sql.NullString

	err = rows.Scan(&user.UUID, &user.EmailAddress, &user.FirstName, &lastName, &dateOfBirth, &phoneNumber)

	if lastName_.Valid {
		user.LastName = lastName_.String
	}
	if dateOfBirth.Valid {
		user.DOB = dateOfBirth.Time.String()
	}
	if phoneNumber_.Valid {
		user.PhoneNumber = phoneNumber_.String
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errUserDoesNotExistError
		}
	}

	return &user, err
}

func GetUser(id string) (*User, error) {
	for _, val := range UserStore {
		if val.UUID == id {
			return val, nil
		}
	}

	query := `SELECT * from ` + UserTableName + ` WHERE "UUID" = ` + `'` + id + `'`

	fmt.Println(query)

	result := db.QueryRow(query)

	var user User
	var lastName sql.NullString
	var dateOfBirth sql.NullTime
	var phoneNumber sql.NullString

	err := result.Scan(&user.UUID, &user.EmailAddress, &user.FirstName, &lastName, &dateOfBirth, &phoneNumber)

	if lastName.Valid {
		user.LastName = lastName.String
	}
	if dateOfBirth.Valid {
		user.DOB = dateOfBirth.Time.String()
	}
	if phoneNumber.Valid {
		user.PhoneNumber = phoneNumber.String
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errUserDoesNotExistError
		}
	}

	return &user, nil
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

func ListUsers() ([]*User, error) {

	query := `SELECT * FROM ` + UserTableName

	var users []*User

	result, err := db.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		} else {
			return nil, err
		}
	}

	for result.Next() {
		var user User
		var lastName sql.NullString
		var dateOfBirth sql.NullTime
		var phoneNumber sql.NullString

		scanerr := result.Scan(&user.UUID, &user.EmailAddress, &user.FirstName, &lastName, &dateOfBirth, &phoneNumber)

		if scanerr != nil {
			fmt.Println(scanerr)
		}

		if lastName.Valid {
			user.LastName = lastName.String
		}
		if dateOfBirth.Valid {
			user.DOB = dateOfBirth.Time.String()
		}
		if phoneNumber.Valid {
			user.PhoneNumber = phoneNumber.String
		}

		users = append(users, &user)

	}

	return users, nil
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
