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

	router.RegisterHandler("/test", "GET", func(rw http.ResponseWriter, r *http.Request) { fmt.Fprintf(rw, "test endpoint") })

	router.RegisterDynamicHandler("/car/{id}/", "GET", func(rw http.ResponseWriter, r *http.Request, m map[string]string) { fmt.Fprintf(rw, m["id"]) })

	fmt.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
