package main

import (
	"UserCrud/user"
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func RestUserService() {

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

func GRPCUserService() {
	server := grpc.NewServer()
	con, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	var userServer user.Server
	user.RegisterUserServiceServer(server, &userServer)
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}

func main() {

	user.DBConnect()
	// RestUserService()
	GRPCUserService()
}
