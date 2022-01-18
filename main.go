package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := NewRouter()
	router.RegisterHandler("/ping", "GET", PingHandler())
	router.RegisterHandler("/user", "POST", CreateUserHandler())
	router.RegisterHandler("/getuser", "POST", GetUserHandler())
	router.RegisterHandler("/users/list", "GET", ListUsersHandler())
	router.RegisterHandler("/updateuser", "POST", UpdateUserHandler())
	router.RegisterHandler("/deleteuser", "POST", DeleteUserHandler())

	router.RegisterHandler("/test", "GET", func(rw http.ResponseWriter, r *http.Request) { fmt.Fprintf(rw, "test endpoint") })

	fmt.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
