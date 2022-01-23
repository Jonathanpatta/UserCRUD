package main

import (
	"fmt"
	"net/http"
)

func main() {

	DBConnect()
	router := NewRouter()
	router.RegisterHandler("/ping", "GET", PingHandler())
	router.RegisterHandler("/users", "POST", CreateUserHandler())
	router.RegisterDynamicHandler("/users/{id}", "PUT", UpdateUserHandler())
	router.RegisterDynamicHandler("/users/{id}", "GET", GetUserHandler())
	router.RegisterDynamicHandler("/users/{id}", "DELETE", DeleteUserHandler())
	router.RegisterHandler("/users", "GET", ListUsersHandler())

	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
