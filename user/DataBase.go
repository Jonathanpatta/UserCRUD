package user

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
	fmt.Println("connected!")
	db = db_

	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

}
