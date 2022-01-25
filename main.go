package main

import (
	// "UserCrud/restapi"
	"UserCrud/user"
	"net"

	"google.golang.org/grpc"
)

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
	// restapi.RestUserService()
	GRPCUserService()
}
