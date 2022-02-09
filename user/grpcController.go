package user

import (
	"UserCrud/pb"
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := CreateUser(in.EmailAddress, in.FirstName, in.LastName, in.PhoneNumber, in.DOB.AsTime())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UpdateUserInput) (*pb.User, error) {
	user, err := UpdateUser(in.Id, in.User)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.IdInput) (*pb.User, error) {
	user, err := GetUser(in.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.IdInput) (*emptypb.Empty, error) {
	err := DeleteUser(in.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListUsers(ctx context.Context, e *emptypb.Empty) (*pb.ListUserOutput, error) {
	users, err := ListUsers()
	if err != nil {
		return nil, err
	}
	usersOutput := pb.ListUserOutput{Users: users}
	return &usersOutput, nil
}
