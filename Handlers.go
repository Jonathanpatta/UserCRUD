package main

import (
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
		switch r.Method {
		case "POST":
			fmt.Fprintf(rw, "create user")
		default:
			http.Error(rw, "404 not found.", http.StatusNotFound)
		}
	}
}
