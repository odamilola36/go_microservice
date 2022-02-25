package models

import (
	"gopkg.in/mgo.v2/bson"
	"microservices/pb"
	"time"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,"`
	Email    string        `json:"email"`
	Name     string        `json:"name"`
	Password string        `json:"password"`
	Created  time.Time     `json:"created"`
	Updated  time.Time     `json:"updated"`
}

func (u *User) toProtoBuffer() *pb.User {
	return &pb.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Created:  u.Created.Unix(),
		Updated:  u.Updated.Unix(),
	}
}

func (u *User) fromProtocol(user *pb.User) {
	u.Id = bson.ObjectIdHex(user.GetId())
	u.Email = user.GetEmail()
	u.Name = user.GetName()
	u.Password = user.GetPassword()
	u.Created = time.Unix(user.GetCreated(), 0)
	u.Updated = time.Unix(user.GetUpdated(), 0)
}
