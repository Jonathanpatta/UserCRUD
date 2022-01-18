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
			fmt.Fprint(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(user)
	}
}

func GetUserHandler() func(rw http.ResponseWriter, r *http.Request, m map[string]string) {
	return func(rw http.ResponseWriter, r *http.Request, m map[string]string) {
		rw.Header().Set("Content-Type", "application/json")
		id := m["id"]

		user, err := GetUser(id)
		if err != nil {
			fmt.Fprint(rw, err.Error())
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
			fmt.Fprint(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(users)
	}
}

func UpdateUserHandler() func(rw http.ResponseWriter, r *http.Request, m map[string]string) {
	return func(rw http.ResponseWriter, r *http.Request, m map[string]string) {
		rw.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		email := r.Form.Get("email")
		firstName := r.Form.Get("firstname")
		lastName := r.Form.Get("lastname")
		phoneNumber := r.Form.Get("phonenumber")
		dateOfBirth := r.Form.Get("dob")

		id := m["id"]

		user, err := UpdateUser(id, &User{EmailAddress: email, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, DOB: dateOfBirth})
		if err != nil {
			fmt.Fprint(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(user)
	}
}

func DeleteUserHandler() func(rw http.ResponseWriter, r *http.Request, m map[string]string) {
	return func(rw http.ResponseWriter, r *http.Request, m map[string]string) {
		rw.Header().Set("Content-Type", "application/json")

		id := m["id"]

		err := DeleteUser(id)
		if err != nil {
			fmt.Fprint(rw, err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(rw, "Deletion Successful")
	}
}
