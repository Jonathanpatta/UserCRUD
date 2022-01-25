package main

import (
	"UserCrud/user"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	res, err := client.CreateUser(context.Background(), &user.User{EmailAddress: "jonaasdjf", FirstName: "jonathan"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
