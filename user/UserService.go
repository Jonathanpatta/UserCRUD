package user

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// type User struct {
// 	EmailAddress string    `json:"emailaddress,omitempty"`
// 	FirstName    string    `json:"firstname,omitempty"`
// 	LastName     string    `json:"lastname,omitempty"`
// 	PhoneNumber  string    `json:"phonenumber,omitempty"`
// 	DOB          time.Time `json:"dob,omitempty"`
// 	UUID         string    `json:"UUID,omitempty"`
// }

var ErrEmptyEmail = errors.New("email address cannot be empty")
var ErrEmptyFirstName = errors.New("first name cannot be empty")
var ErrUserDoesNotExist = errors.New("user does not exist")

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateUser(emailAddr string, firstName string, lastName string, phoneNumber string, dob time.Time) (*User, error) {
	if emailAddr == "" {
		return nil, ErrEmptyEmail
	} else if firstName == "" {
		return nil, ErrEmptyFirstName
	}
	id := randSeq(36)

	myuser := User{}
	myuser.UUID = id
	myuser.EmailAddress = emailAddr
	myuser.FirstName = firstName
	myuser.LastName = lastName
	myuser.PhoneNumber = phoneNumber
	time := timestamppb.New(dob)
	myuser.DOB = time

	fields, values := ``, ``

	var args []interface{}
	nargs := 1
	if myuser.UUID != "" {
		fields += `"UUID"`
		values += `$` + strconv.Itoa(nargs) + `::text`
		args = append(args, myuser.UUID)
		nargs++

	}
	if myuser.EmailAddress != "" {
		fields += `, "Email"`
		values += `, $` + strconv.Itoa(nargs) + `::text`
		args = append(args, myuser.EmailAddress)
		nargs++

	}
	if myuser.FirstName != "" {
		fields += `, "FirstName"`
		values += `, $` + strconv.Itoa(nargs) + `::text`
		args = append(args, myuser.FirstName)
		nargs++
	}

	if myuser.LastName != "" {
		fields += `, "LastName"`
		values += `, $` + strconv.Itoa(nargs) + `::text`
		args = append(args, myuser.LastName)
		nargs++

	}
	if myuser.PhoneNumber != "" {
		fields += `, "PhoneNumber"`
		values += `, $` + strconv.Itoa(nargs) + `::text`
		args = append(args, myuser.PhoneNumber)
		nargs++

	}

	if !myuser.DOB.AsTime().IsZero() {
		fields += `, "DateOfBirth"`
		values += `, $` + strconv.Itoa(nargs) + `::date`
		args = append(args, myuser.DOB.AsTime().Format("2006-01-02"))
		nargs++

	}

	fields, values = ` (`+fields+`) `, ` (`+values+` ) `

	query := `INSERT INTO ` + UserTableName + fields + `VALUES` + values + ` returning *;`

	rows := db.QueryRow(query, args...)

	var user User
	var lastName_ sql.NullString
	var dateOfBirth sql.NullTime
	var phoneNumber_ sql.NullString

	err := rows.Scan(&user.UUID, &user.EmailAddress, &user.FirstName, &lastName_, &dateOfBirth, &phoneNumber_)

	if lastName_.Valid {
		user.LastName = lastName_.String
	}
	if dateOfBirth.Valid {
		user.DOB = timestamppb.New(dateOfBirth.Time)
	}
	if phoneNumber_.Valid {
		user.PhoneNumber = phoneNumber_.String
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExist
		} else {
			return nil, err
		}
	}

	return &user, err
}

func GetUser(id string) (*User, error) {

	query := `SELECT * from ` + UserTableName + ` WHERE "UUID" = ` + `$1`

	result := db.QueryRow(query, id)

	var user User
	var lastName sql.NullString
	var dateOfBirth sql.NullTime
	var phoneNumber sql.NullString

	err := result.Scan(&user.UUID, &user.EmailAddress, &user.FirstName, &lastName, &dateOfBirth, &phoneNumber)

	if lastName.Valid {
		user.LastName = lastName.String
	}
	if dateOfBirth.Valid {
		user.DOB = timestamppb.New(dateOfBirth.Time)
	}
	if phoneNumber.Valid {
		user.PhoneNumber = phoneNumber.String
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExist
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func UpdateUser(id string, updatedUser *User) (*User, error) {

	if updatedUser.EmailAddress == "" {
		return nil, ErrEmptyEmail
	} else if updatedUser.FirstName == "" {
		return nil, ErrEmptyFirstName
	}

	fields := ``
	var args []interface{}

	args = append(args, updatedUser.EmailAddress, updatedUser.FirstName, updatedUser.LastName, updatedUser.PhoneNumber)

	fields += `"Email" = $1::text, `
	fields += `"FirstName" = $2::text, `
	fields += `"LastName" = $3::text, `
	fields += `"PhoneNumber" = $4::text `
	nargs := 5
	if !updatedUser.DOB.AsTime().IsZero() {
		fields += `, "DateOfBirth" = $5::date `
		args = append(args, updatedUser.DOB.AsTime().Format("2006-01-02"))
		nargs++
	}
	args = append(args, id)

	query := `UPDATE ` + UserTableName + ` SET ` + fields + ` WHERE "UUID" = $` + strconv.Itoa(nargs) + ` returning *;`

	rows := db.QueryRow(query, args...)

	var user User
	var lastName sql.NullString
	var dateOfBirth sql.NullTime
	var phoneNumber sql.NullString

	err := rows.Scan(&user.UUID, &user.EmailAddress, &user.FirstName, &lastName, &dateOfBirth, &phoneNumber)

	if lastName.Valid {
		user.LastName = lastName.String
	}
	if dateOfBirth.Valid {
		user.DOB = timestamppb.New(dateOfBirth.Time)
	}
	if phoneNumber.Valid {
		user.PhoneNumber = phoneNumber.String
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExist
		} else {
			return nil, err
		}
	}

	return &user, nil
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
			user.DOB = timestamppb.New(dateOfBirth.Time)
		}
		if phoneNumber.Valid {
			user.PhoneNumber = phoneNumber.String
		}

		users = append(users, &user)

	}

	return users, nil
}

func DeleteUser(id string) error {

	query := `DELETE FROM ` + UserTableName + ` WHERE "UUID" = $1`

	_, err := db.Query(query, id)

	if err != nil {
		return err
	}

	return nil
}
