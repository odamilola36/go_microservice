package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"microservices/authentication/models"
	"microservices/authentication/repositories"
	"microservices/authentication/validators"
	"microservices/pb"
	"microservices/security"
	"strings"
	"time"
)

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) pb.AuthServiceServer {
	return &authService{
		repo,
	}
}

func (a *authService) Signup(ctx context.Context, u *pb.User) (*pb.User, error) {
	err := validators.ValidateSignUp(u)
	if err != nil {
		return nil, err
	}
	//a.repo
	u.Name = strings.TrimSpace(u.Name)
	u.Password, err = security.EncryptPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Email = validators.NormalizeEmail(u.Email)
	var result models.User
	result, err = a.repo.GetByEmail(u.Email)
	if err == mongo.ErrNoDocuments {
		user := new(models.User)
		user.FromProtocol(u)
		err := a.repo.Save(user)
		if err != nil {
			return nil, err
		}
		return user.ToProtoBuffer(), nil
	}

	if result == (models.User{}) {
		return nil, err
	}

	return nil, validators.EmailAlreadyExist
}

func (a *authService) GetUser(ctx context.Context, u *pb.GetUserRequest) (*pb.User, error) {
	if !primitive.IsValidObjectID(u.Id) {
		return nil, validators.ErrorInvalidUserId
	}
	found, err := a.repo.GetById(u.Id)
	if err != nil {
		return nil, err
	}
	return found.ToProtoBuffer(), nil
}

func (a *authService) Update(ctx context.Context, u *pb.User) (*pb.User, error) {
	if !primitive.IsValidObjectID(u.Id) {
		return nil, validators.ErrorInvalidUserId
	}
	user, err := a.repo.GetById(u.GetId())

	if err != nil {
		return nil, err
	}
	u.Name = strings.TrimSpace(u.Name)

	if u.Name == "" {
		return nil, validators.ErrorName
	}

	if u.Name == user.Name {
		return user.ToProtoBuffer(), nil
	}

	user.Name = u.Name
	user.Updated = time.Now()
	us := mapper(user)
	err = a.repo.UpdateUser(us)
	return user.ToProtoBuffer(), err
}

func (a *authService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error {
	users, err := a.repo.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := stream.Send(user.ToProtoBuffer()); err != nil {
			return err
		}
	}
	return nil
}

func (a *authService) DeleteUser(ctx context.Context, u *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	if !primitive.IsValidObjectID(u.Id) {
		return nil, validators.ErrorInvalidUserId
	}

	err := a.repo.DeleteUser(u.Id)
	return &pb.DeleteUserResponse{Id: u.Id}, err
}

func (a *authService) SignIn(ctx context.Context, request *pb.SigninRequest) (*pb.SigninResponse, error) {
	user, err := a.repo.GetByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	ok := security.VerifyPassword(user.Password, request.Password)
	if !ok {
		return nil, validators.ErrorInvalidPassword
	}
	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		return nil, err
	}
	return &pb.SigninResponse{
		User:  user.ToProtoBuffer(),
		Token: token,
	}, nil
}

func mapper(user models.User) *models.User {
	us := new(models.User)
	us.Email = user.Email
	us.Name = user.Name
	us.Id = user.Id
	us.Created = user.Created
	us.Updated = user.Updated
	return us
}
