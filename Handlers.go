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
