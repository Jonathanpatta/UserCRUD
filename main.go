package main

import (
	"fmt"
	"net/http"
)

// handles both cases where the url ends with "/" and the case where it does'nt
func HandleBoth(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
	http.HandleFunc(url+"/", handler)
}

func main() {

	HandleBoth("/ping", PingHandler())
	HandleBoth("/user", CreateUserHandler())
	HandleBoth("/getuser", GetUserHandler())
	HandleBoth("/users/list", ListUsersHandler())
	HandleBoth("/updateuser", UpdateUserHandler())

	fmt.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
