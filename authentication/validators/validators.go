package validators

import (
	"errors"
	"microservices/pb"
)

var (
	ErrorName         = errors.New("name can't be empty")
	ErrorPassword     = errors.New("password can't be empty")
	ErrorEmail        = errors.New("email can't be empty")
	EmailAlreadyExist = errors.New("email already used by another")
)

func ValidateSignUp(u *pb.User) error {
	if u.Email == "" {
		return ErrorEmail
	}
	if u.Password == "" {
		return ErrorPassword
	}
	if u.Name == "" {
		return ErrorName
	}
	return nil
}
