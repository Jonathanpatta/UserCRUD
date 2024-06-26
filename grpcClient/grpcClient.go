package main

import (
	"UserCrud/pb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	res, err := client.CreateUser(context.Background(), &pb.User{EmailAddress: "jonaasdjf", FirstName: "jonathan"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = client.GetUser(context.Background(), &pb.IdInput{Id: "tmgsK50iVuKziGYj0tCRaQgH3cjzc8GBB22O"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	updateInput := &pb.UpdateUserInput{
		Id: "tmgsK50iVuKziGYj0tCRaQgH3cjzc8GBB22O",
		User: &pb.User{
			EmailAddress: "jonaasdjf",
			FirstName:    "jonathan",
			PhoneNumber:  "1234567890",
		}}

	res, err = client.UpdateUser(context.Background(), updateInput)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	listUserOut, err := client.ListUsers(context.Background(), &emptypb.Empty{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(listUserOut)
	}
}
