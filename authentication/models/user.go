package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"microservices/pb"
	"time"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
	Created  time.Time          `json:"created"`
	Updated  time.Time          `json:"updated"`
}

func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Created:  u.Created.Unix(),
		Updated:  u.Updated.Unix(),
	}
}

func (u *User) FromProtocol(user *pb.User) {
	hex, err := primitive.ObjectIDFromHex(user.GetId())
	if err != nil {
		panic(err)
	}
	u.Id = hex
	u.Email = user.GetEmail()
	u.Name = user.GetName()
	u.Password = user.GetPassword()
	u.Created = time.Unix(user.GetCreated(), 0)
	u.Updated = time.Unix(user.GetUpdated(), 0)
}
