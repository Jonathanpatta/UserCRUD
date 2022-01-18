package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PingHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "pong")
	}
}

func CreateUserHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		email := r.Form.Get("email")
		firstName := r.Form.Get("firstname")
		lastName := r.Form.Get("lastname")
		phoneNumber := r.Form.Get("phonenumber")
		dateOfBirth := r.Form.Get("dob")

		user, err := CreateUser(email, firstName, lastName, phoneNumber, dateOfBirth)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(user)
	}
}

func GetUserHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		r.ParseForm()
		id := r.Form.Get("id")

		user, err := GetUser(id)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(user)
	}
}

func ListUsersHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		users, err := ListUsers()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(users)
	}
}
