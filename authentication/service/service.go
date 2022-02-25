package service

import (
	"context"
	"microservices/authentication/repositories"
	"microservices/authentication/validators"
	"microservices/pb"
)

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(r repositories.UserRepository) *pb.AuthServiceServer {
	return &authService{repo: r}
}

func (a *authService) SignUp(ctx context.Context, u *pb.User) (*pb.User, error) {
	err := validators.ValidateSignUp(u)
	if err != nil {
		return nil, err
	}
	//a.repo
}

func (a *authService) GetUser(ctx context.Context, u *pb.GetUserRequest) (*pb.User, error) {

}

func (a *authService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthServiceServer) (*pb.User, error) {

}

func (a *authService) UpdateUser(ctx context.Context, pUser *pb.User) (*pb.User, error) {

}

func (a *authService) DeleteUser(ctx context.Context, pUser *pb.User) (*pb.DeleteUserResponse, error) {

}
