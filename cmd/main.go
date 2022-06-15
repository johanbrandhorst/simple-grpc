package main

import (
	"context"
	"log"

	usersv1 "github.com/johanbrandhorst/simple-grpc/gen/users/v1"
	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to server: ", err)
	}

	cli := usersv1.NewUsersServiceClient(clientConn)

	resp, err := cli.CreateUser(context.Background(), &usersv1.CreateUserRequest{
		User: &usersv1.User{
			Name: "John Doe",
			Age:  42,
		},
	})
	if err != nil {
		log.Fatal("Failed to call server: ", err)
	}
	log.Println("Response: ", resp)

	resp2, err := cli.GetUser(context.Background(), &usersv1.GetUserRequest{Id: resp.User.Id})
	if err != nil {
		log.Fatal("Failed to call server: ", err)
	}
	log.Println("Response: ", resp2)
}
