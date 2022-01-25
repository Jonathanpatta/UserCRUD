package user

import "context"

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
