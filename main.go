package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	usersv1 "github.com/johanbrandhorst/simple-grpc/gen/users/v1"
	"google.golang.org/grpc"
)

type server struct {
	usersv1.UnimplementedUsersServiceServer

	users map[string]*usersv1.User
}

func (s *server) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	user := s.users[req.Id]
	return &usersv1.GetUserResponse{User: user}, nil
}

func (s *server) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	req.User.Id = uuid.New().String()
	s.users[req.User.Id] = req.User
	return &usersv1.CreateUserResponse{User: req.User}, nil
}

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to create listener: ", err)
	}

	usersv1.RegisterUsersServiceServer(s, &server{
		users: make(map[string]*usersv1.User),
	})

	log.Println("Starting server on port 8080")

	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
