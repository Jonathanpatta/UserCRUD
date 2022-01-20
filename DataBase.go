package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db            *sql.DB
	UserTableName = `"JonathanUserService"."Users"`
)

func DBConnect() {
	connStr := "postgres://rdjocxwy:mK4lzkBCYM-o3i80kiT4eyzd03C8zgFg@satao.db.elephantsql.com/rdjocxwy?sslmode=verify-full"

	db_, err := sql.Open("postgres", connStr)
	db = db_

	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

}

func GetFieldsAndValues(user *User) (fields string, values string) {
	if user.UUID != "" {
		fields += `"UUID"`
		values += `'` + user.UUID + `'` + `::text`
	}
	if user.EmailAddress != "" {
		fields += `, "Email"`
		values += `, '` + user.EmailAddress + `'` + `::text`

	}
	if user.FirstName != "" {
		fields += `, "FirstName"`
		values += `, '` + user.FirstName + `'` + `::text`
	}

	if user.LastName != "" {
		fields += `, "LastName"`
		values += `, '` + user.LastName + `'` + `::text`
	}
	if user.PhoneNumber != "" {
		fields += `, "PhoneNumber"`
		values += `, '` + user.PhoneNumber + `'` + `::text`

	}

	if user.DOB != "" {
		fields += `, "DateOfBirth"`
		values += `, '` + user.DOB + `'` + `::timestamp without time zone `
	}

	return ` (` + fields + `) `, ` (` + values + ` ) `
}
