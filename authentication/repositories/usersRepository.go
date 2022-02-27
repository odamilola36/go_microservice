package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"microservices/authentication/models"
	"microservices/dbCon"
)

const userCollection = "users"

type UserRepository interface {
	Save(user *models.User) error
	GetById(id string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetAll() ([]models.User, error)
	UpdateUser(u *models.User) error
	DeleteUser(id string) error
	DeleteAll() error
}

type usersRepository struct {
	c         *mongo.Collection
	dbContext context.Context
}

func NewUsersRepository(c dbCon.Connection) UserRepository {
	return &usersRepository{
		c:         c.DB().Collection(userCollection),
		dbContext: c.DBContext(),
	}
}

func (r *usersRepository) Save(user *models.User) error {
	_, err := r.c.InsertOne(r.dbContext, user)
	return err
}

func (r *usersRepository) GetById(id string) (models.User, error) {
	var user models.User
	hex, _ := primitive.ObjectIDFromHex(id)
	err := r.c.FindOne(r.dbContext, bson.M{"_id": hex}).Decode(&user)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.c.FindOne(r.dbContext, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (r *usersRepository) GetAll() ([]models.User, error) {
	var user []models.User
	cursor, err := r.c.Find(r.dbContext, nil)
	err = cursor.All(nil, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *usersRepository) UpdateUser(u *models.User) error {
	var _, err = r.c.UpdateOne(r.dbContext, bson.M{"_id": u.Id}, bson.M{"$set": u})
	return err
}

func (r *usersRepository) DeleteUser(id string) error {
	hex, _ := primitive.ObjectIDFromHex(id)
	var _, err = r.c.DeleteOne(r.dbContext, bson.M{"_id": hex})
	return err
}

func (r *usersRepository) DeleteAll() error {
	return r.c.Drop(r.dbContext)
}
