package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := NewRouter()
	router.RegisterHandler("/ping", "GET", PingHandler())
	router.RegisterHandler("/user", "POST", CreateUserHandler())
	router.RegisterDynamicHandler("/user/{id}", "PUT", UpdateUserHandler())
	router.RegisterDynamicHandler("/user/{id}", "GET", GetUserHandler())
	router.RegisterDynamicHandler("/user/{id}", "DELETE", DeleteUserHandler())
	router.RegisterHandler("/users", "GET", ListUsersHandler())

	fmt.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
