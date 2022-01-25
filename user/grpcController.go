package user

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, in *User) (*User, error) {
	user, err := CreateUser(in.EmailAddress, in.FirstName, in.LastName, in.PhoneNumber, in.DOB.AsTime())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *UpdateUserInput) (*User, error) {
	user, err := UpdateUser(in.Id, in.User)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) GetUser(ctx context.Context, in *IdInput) (*User, error) {
	user, err := GetUser(in.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *IdInput) (*emptypb.Empty, error) {
	err := DeleteUser(in.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListUsers(ctx context.Context, e *emptypb.Empty) (*ListUserOutput, error) {
	users, err := ListUsers()
	if err != nil {
		return nil, err
	}
	usersOutput := ListUserOutput{Users: users}
	return &usersOutput, nil
}
