package validators

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservices/pb"
	"strings"
)

var (
	ErrorInvalidUserId   = errors.New("invalid id")
	ErrorName            = errors.New("name can't be empty")
	ErrorPassword        = errors.New("password can't be empty")
	ErrorEmail           = errors.New("email can't be empty")
	EmailAlreadyExist    = errors.New("email already used by another")
	ErrorInvalidPassword = errors.New("incorrect password")
)

func ValidateSignUp(u *pb.User) error {
	if !primitive.IsValidObjectID(u.Id) {
		return ErrorInvalidUserId
	}
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

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
