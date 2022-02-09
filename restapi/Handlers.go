package restapi

import (
	"UserCrud/pb"
	"UserCrud/user"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
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

		parsedDob := time.Time{}

		if dateOfBirth != "" {
			parsedDob_, err := time.Parse("2006-01-02", dateOfBirth)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			parsedDob = parsedDob_
		}

		user, err := user.CreateUser(email, firstName, lastName, phoneNumber, parsedDob)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
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

		user, err := user.GetUser(id)
		if err != nil {

			http.Error(rw, err.Error(), http.StatusInternalServerError)

			return
		}

		encoder := json.NewEncoder(rw)
		encoder.Encode(user)
	}
}

func ListUsersHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		users, err := user.ListUsers()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
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

		parsedDob := time.Time{}

		if dateOfBirth != "" {
			parsedDob_, err := time.Parse(time.RFC3339, dateOfBirth)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			parsedDob = parsedDob_
		}

		user, err := user.UpdateUser(id, &pb.User{EmailAddress: email, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, DOB: timestamppb.New(parsedDob)})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
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

		err := user.DeleteUser(id)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(rw, "Deletion Successful")
	}
}
